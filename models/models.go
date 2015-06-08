package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Blog struct {
	Id      int `PK`
	Title   string
	Content string
	Created time.Time
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/blog?charset=utf8", 30)

}

func GetAll() (blogs []Blog) {
	db := orm.NewOrm()
	db.Raw("SELECT * FROM blog ").QueryRows(&blogs)

	return blogs
}

func GetBlog(id int) (blog Blog) {
	db := orm.NewOrm()
	db.Raw("SELECT * FROM blog WHERE id=?", id).QueryRow(&blog)
	return blog
}

func SaveBlog(blog Blog) (bg Blog) {
	db := orm.NewOrm()
	db.Raw("insert into blog values(?,?,?,?)", blog.Id, blog.Title, blog.Content, blog.Created).Exec()
	return blog
}

func DelBlog(blog Blog) {
	db := orm.NewOrm()
	db.Raw("delete FROM blog WHERE id=?", blog.Id).Exec()
	return
}
