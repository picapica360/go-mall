package orm

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

const (
	_testTableSQL = `
	CREATE TABLE test_book (
		id int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		name varchar(32) NOT NULL,
		author varchar(64) NOT NULL,
		desc varchar(256),
		ISBN varchar(32) NOT NULL,
		pub date NOT NULL,
		amount decimal(12,2),
		is_putaway tinyint NOT NULL,
		is_deleted tinyint NOT NULL,
		created_at datetime NOT NULL,
		updated_at datetime NOT NULL,
		deleted_at datetime
	  ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='框架功能测试表';	  
		`

	_grant = `
	create user 'mall'@'localhost' identified by '123456';
	grant insert,select,update,delete,drop,create,alter on mall.* to 'mall'@'localhost';
	flush privileges;
	`

	_insertSQL = `
	insert into test_book(name,author,desc,ISBN,pub,amount,is_putaway,s_deleted,created_at,updated_at) values(?,?,?,?,?,?,?,?,?,?)
	`
)

// only for test.
type _book struct {
	ID        int        `gorm:"primary_key"` // id
	Name      string     // varchar
	Author    string     // varchar
	Desc      string     // varchar
	ISBN      string     // varchar
	Pub       time.Time  // date
	Amount    float64    // decimal(12,2)
	IsPutaway bool       // tinyint
	IsDeleted bool       // tinyint
	CreatedAt time.Time  // datetime
	UpdatedAt time.Time  // datetime
	DeletedAt *time.Time // datetime null
}

func (*_book) TableName() string {
	return "test_book"
}

func testNewDB(t *testing.T) *gorm.DB {
	cfg := Config{
		Dialect: "mysql",
		DSN:     "mall:123456@(localhost)/mall?charset=utf8&parseTime=True&loc=Local",
	}
	db := newDB(&cfg, func(err error) {
		t.Errorf("db error: %v", err)
		panic(err)
	})
	return db
}

func testCloseDB(db *gorm.DB) {
	if db != nil {
		db.Close()
	}
}

func s2t(value string) (t time.Time) {
	// 2006-01-02 15:04:05
	t, _ = time.ParseInLocation("2006-01-02", value, time.Local)
	return
}

func TestDBConnect(t *testing.T) {
	db := testNewDB(t)
	defer testCloseDB(db)
}

func TestInsert(t *testing.T) {
	db := testNewDB(t)
	defer testCloseDB(db)

	book := _book{
		Name:      "Golang",
		Author:    "Alan Donovan",
		Desc:      "The Go Programming Language",
		ISBN:      "9787111558421",
		Pub:       s2t("2017-05-01"),
		Amount:    79.99,
		IsPutaway: true,
		IsDeleted: false,
	}

	if err := db.Create(&book).Error; err != nil {
		t.Error(err)
	}
}

func TestInsert2(t *testing.T) {
	db := testNewDB(t)
	defer testCloseDB(db)

	book := _book{
		Name:      "C",
		Author:    "Brian W. Kernighan & Dennis M. Ritchie",
		Desc:      "The C Programming Language",
		ISBN:      "9787111128069",
		Pub:       s2t("2004-01-01"),
		Amount:    30.99,
		IsPutaway: true,
		IsDeleted: false,
	}

	if err := db.Create(&book).Error; err != nil {
		t.Error(err)
	}

	if book.ID == 0 {
		t.Error("the primary_key is zero, must AUTO_INCREMENT.")
	}
}

func TestMultiInsert(t *testing.T) {
	db := testNewDB(t)
	defer testCloseDB(db)

	books := []_book{
		{
			Name:      "Golang",
			Author:    "Alan Donovan",
			Desc:      "The Go Programming Language",
			ISBN:      "9787111558421",
			Pub:       s2t("2017-05-01"),
			Amount:    79.99,
			IsPutaway: true,
			IsDeleted: false,
		},
		{
			Name:      "C",
			Author:    "Brian W. Kernighan & Dennis M. Ritchie",
			Desc:      "The C Programming Language",
			ISBN:      "9787111128069",
			Pub:       s2t("2004-01-01"),
			Amount:    30.99,
			IsPutaway: true,
			IsDeleted: false,
		},
		{
			Name:      "C#",
			Author:    "Anders Hejlsberg",
			Desc:      "The C# Programming Language",
			ISBN:      "9787111282617",
			Pub:       s2t("2010-01-01"),
			Amount:    79.99,
			IsPutaway: true,
			IsDeleted: false,
		},
		{
			Name:      "Java",
			Author:    "Ken Arnold",
			Desc:      "The Java Programming Language, 4th Edition",
			ISBN:      "9787115152978",
			Pub:       s2t("2006-11-01"),
			Amount:    69.00,
			IsPutaway: true,
			IsDeleted: false,
		},
		{
			Name:      "Python",
			Author:    "Y. Daniel Liang",
			Desc:      "Introduction to Programming Using Python",
			ISBN:      "9787111412342",
			Pub:       s2t("2013-03-01"),
			Amount:    79.99,
			IsPutaway: true,
			IsDeleted: false,
		},
	}

	if err := db.Create(&books).Error; err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	db := testNewDB(t)
	defer testCloseDB(db)

	var book _book
	if err := db.Where("name = ?", "Golang").First(&book).Error; err != nil {
		t.Errorf("%v", err)
	}

	if err := db.Model(&book).Update("amount", 78.99).Error; err != nil {
		t.Errorf("%v", err)
	}
}

func TestDelete(t *testing.T) {
	db := testNewDB(t)
	defer testCloseDB(db)

	var book _book
	if err := db.Where("name = ?", "Golang").First(&book).Error; err != nil {
		t.Errorf("%v", err)
	}

	if err := db.Delete(&book).Error; err != nil {
		t.Errorf("%v", err)
	}
}

func TestQuery(t *testing.T) {
	db := testNewDB(t)
	defer testCloseDB(db)

	var (
		err   error
		book  _book
		books []_book // []*_book
	)

	// NOTE: First/Find 方法，在查找单条数据时，若返回空值会报 ErrRecordNotFound 错误

	if err = db.Where("name = ?", "Golang").First(&book).Error; err != nil {
		t.Errorf("%v", err)
	}

	if err = db.Where("name in (?)", []string{"Golang", "C"}).Find(&books).Error; err != nil {
		t.Errorf("%v", err)
	}
}

func TestQueryWithoutDeleted(t *testing.T) {
	db := testNewDB(t)
	defer testCloseDB(db)

	var books []_book
	if err := db.Where("name = ?", "Golang").Find(&books).Error; err != nil {
		t.Errorf("%v", err)
	}

	if len(books) != 0 {
		t.Errorf("%v", books)
	}
}
