package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var log = logs.NewLogger(10000)

func init() {
	log.SetLogger("console", "")
}

type Blog struct {
	Id              int `PK`
	Author          int
	Date            time.Time
	DateGmt         time.Time
	Content         string
	Title           string
	Excerpt         string
	Status          string
	CommentSatus    string
	PingStatus      string
	Password        string
	Name            string
	ToPing          string
	Pinged          string
	Modified        time.Time
	ModifiedGmt     time.Time
	ContentFiltered string
	Parent          int
	Guid            string
	MenuOrder       int
	Type            string
	MimeType        string
	CommentCount    int
}
type User struct {
	Id            int `PK`
	Login         string
	Pass          string
	NiceName      string
	Email         string
	Url           string
	Registered    time.Time
	ActivationKey string
	Status        int
	DisplayName   string
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/blog?charset=utf8", 30)

}

func GetAll() (blogs []Blog) {
	db := orm.NewOrm()
	db.Raw("SELECT * FROM bb_posts ").QueryRows(&blogs)
	return blogs
}

func GetBlog(id int) (blog Blog) {
	db := orm.NewOrm()
	db.Raw("SELECT * FROM bb_posts WHERE id=?", id).QueryRow(&blog)
	return blog
}

func SaveBlog(blog Blog) (bg Blog) {
	db := orm.NewOrm()
	now := time.Now()
	db.Raw("insert into bb_posts values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", blog.Id,
		blog.Author, now, now, blog.Content, blog.Title, blog.Excerpt, blog.Status, blog.CommentSatus, blog.PingStatus, blog.Password,
		blog.Name, blog.ToPing, blog.Pinged, now, now, blog.ContentFiltered, blog.Parent, blog.Guid, blog.MenuOrder,
		blog.Title, blog.MimeType, blog.CommentCount).Exec()
	return blog
}

func DelBlog(blog Blog) {
	db := orm.NewOrm()
	db.Raw("delete FROM bb_posts WHERE id=?", blog.Id).Exec()
	return
}
func UpdateBlog(id string, key string, value interface{}) {
	db := orm.NewOrm()
	if val, ok := value.(int); ok {
		sql := "update bb_posts set " + key + "='" + strconv.Itoa(val) + "' where id=" + id
		log.Debug("debug", sql)
		db.Raw(sql).Exec()
	} else if val, ok := value.(string); ok {
		sql := "update bb_posts set " + key + "='" + string(val) + "' where id=" + id
		log.Debug("debug", sql)
		db.Raw(sql).Exec()
	}
	return
}
