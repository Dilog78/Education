package post

import (
	"education/pkg"
)

type Post struct {
	UserID int
	ID     int
	Title  string
	Body   string
}

func (p *Post) GetPost(id int) (Post, error) {
	db := pkg.InitDB()

	if err := db.Table("posts").Where("id", id).Take(&p).Error; err != nil {
		return *p, err
	}

	return *p, nil
}

func GetPosts() ([]Post, error) {
	db := pkg.InitDB()

	var p []Post

	if err := db.Table("posts").Find(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (p *Post) Create() (int, error) {
	db := pkg.InitDB()

	if err := db.Table("posts").Create(&p).Error; err != nil {
		return 0, err
	}
	return p.ID, nil
}

func (p *Post) DeletePost(id int) error {

	db := pkg.InitDB()

	if err := db.Table("posts").Delete(&p, id).Error; err != nil {
		return err
	}

	return nil
}

func NewPost(u int, t, b string) *Post {
	return &Post{
		UserID: u,
		Title:  t,
		Body:   b,
	}
}

func NewUpdate(title, body string) *Post {
	return &Post{
		Title: title,
		Body:  body,
	}
}

func (p *Post) UpdatePost(id int) error {
	db := pkg.InitDB()

	if err := db.Table("posts").Where("id", id).Updates(&p).Error; err != nil {
		return err
	}

	return nil
}
