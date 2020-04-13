package log

// Config config for file logger.
type Config struct {
	Filepath   string // 日志文件路径, 为 empty 使用将采用 os.TempDir() 临时目录
	MaxSize    int    // 每个日志文件保存的最大尺寸，默认为 100。 单位：M
	MaxBackups int    // 日志文件最多保存多少个备份
	MaxAge     int    // 文件最多保存多长时间  单位：天
	Compress   bool   // 是否压缩。默认不压缩文件
}
