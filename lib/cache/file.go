package cache

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"
)

var (
	// FileCachePath cache directory
	FileCachePath = "cache"
	// FileCacheFileSuffix cache file suffix
	FileCacheFileSuffix = ".bin"
	// FileCacheDirectoryLevel cache file deep level if auto generated cache files.
	FileCacheDirectoryLevel = 2
	// FileCacheEmbedExpiry cache expire time, default is no expire forever.
	FileCacheEmbedExpiry time.Duration
)

// FileCacheItem store the cache item.
type FileCacheItem struct {
	Data       interface{}
	Lastaccess time.Time
	Expired    time.Time
}

// FileCache cache adapter for file cache.
type FileCache struct {
	CachePath      string
	FileSuffix     string
	DirectoryLevel int // file deep level
	EmbedExpiry    int
}

func init() {
	Register("file", NewFileCache)
}

// NewFileCache create a new file cache.
func NewFileCache() Cache {
	return &FileCache{}
}

// Get get cached value by key.
func (c *FileCache) Get(key string) interface{} {
	data, err := fileGetContents(c.getCacheFileName(key))
	if err != nil {
		return ""
	}

	var to FileCacheItem
	gobDecode(data, &to)
	if to.Expired.Before(time.Now()) {
		return ""
	}

	return to.Data
}

// GetMulti get a batch cached values.
func (c *FileCache) GetMulti(keys []string) []interface{} {
	var vals []interface{}
	for _, key := range keys {
		vals = append(vals, c.Get(key))
	}

	return vals
}

// Put set cached value.
func (c *FileCache) Put(key string, val interface{}, timeout time.Duration) error {
	gob.Register(val)

	item := FileCacheItem{Data: val}
	if timeout == time.Duration(c.EmbedExpiry) {
		item.Expired = time.Now().Add((86400 * 365 * 10) * time.Second)
	} else {
		item.Expired = time.Now().Add(timeout)
	}
	item.Lastaccess = time.Now()
	data, err := gobEncode(item)
	if err != nil {
		return err
	}

	return filePutContents(c.getCacheFileName(key), data)
}

// Delete delete cached value by key.
func (c *FileCache) Delete(key string) error {
	f := c.getCacheFileName(key)
	if ok, _ := exists(f); ok {
		return os.Remove(f)
	}
	return nil
}

// IsExist check whether cached value exists.
func (c *FileCache) IsExist(key string) bool {
	ok, _ := exists(c.getCacheFileName(key))
	return ok
}

// Increment increment counter.
func (c *FileCache) Increment(key string) error {
	data := c.Get(key)
	var val int
	if reflect.TypeOf(data).Name() != "int" {
		val = 0
	} else {
		val = data.(int) + 1
	}
	c.Put(key, val, time.Duration(c.EmbedExpiry))

	return nil
}

// Decrement decrement counter.
func (c *FileCache) Decrement(key string) error {
	data := c.Get(key)
	var val int
	if reflect.TypeOf(data).Name() != "int" && data.(int)-1 <= 0 {
		val = 0
	} else {
		val = data.(int) - 1
	}
	c.Put(key, val, time.Duration(c.EmbedExpiry))

	return nil
}

// ClearAll clear all cache.
func (c *FileCache) ClearAll() error {
	return nil
}

// StartAndGC will start and begin gc for file cache.
// the config need to be like {"CachePath":"/cache","FileSuffix":".bin","DirectoryLevel":"2","EmbedExpiry":"0"}
func (c *FileCache) StartAndGC(config string) error {
	conf := make(map[string]string)
	err := json.Unmarshal([]byte(config), &conf)
	if err != nil {
		return nil
	}

	if _, ok := conf["CachePath"]; !ok {
		conf["CachePath"] = FileCachePath
	}
	if _, ok := conf["FileSuffix"]; !ok {
		conf["FileSuffix"] = FileCacheFileSuffix
	}
	if _, ok := conf["DirectoryLevel"]; !ok {
		conf["DirectoryLevel"] = strconv.Itoa(FileCacheDirectoryLevel)
	}
	if _, ok := conf["EmbedExpiry"]; !ok {
		conf["EmbedExpiry"] = strconv.FormatInt(int64(FileCacheEmbedExpiry.Seconds()), 10)
	}
	c.CachePath = conf["CachePath"]
	c.FileSuffix = conf["FileSuffix"]
	c.DirectoryLevel, _ = strconv.Atoi(conf["DirectoryLevel"])
	c.EmbedExpiry, _ = strconv.Atoi(conf["EmbedExpiry"])

	// check, create the directory if not exist.
	if ok, _ := exists(c.CachePath); !ok {
		return os.MkdirAll(c.CachePath, os.ModePerm)
	}

	return nil
}

// get cached file name, md5 encode.
func (c *FileCache) getCacheFileName(key string) string {
	m := md5.New()
	io.WriteString(m, key)
	keyMd5 := hex.EncodeToString(m.Sum(nil))
	cachePath := c.CachePath

	switch c.DirectoryLevel {
	case 2:
		cachePath = filepath.Join(cachePath, keyMd5[0:2], keyMd5[2:4])
	case 1:
		cachePath = filepath.Join(cachePath, keyMd5[0:2])
	}

	if ok, _ := exists(cachePath); !ok {
		_ = os.MkdirAll(cachePath, os.ModePerm)
	}

	return filepath.Join(cachePath, fmt.Sprintf("%s%s", keyMd5, c.FileSuffix))
}

func fileGetContents(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func filePutContents(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, os.ModePerm)
}

// check file exists.
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, err
	}
	if os.IsNotExist(err) {
		return false, err
	}

	return false, err
}

func gobEncode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func gobDecode(data []byte, to *FileCacheItem) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
