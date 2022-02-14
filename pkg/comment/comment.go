package comment

import (
	"education/pkg"
)

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func GetCommByPost(postId int) ([]Comment, error) {
	var c []Comment
	db := pkg.InitDB()

	if err := db.Table("comments").Where("post_id", postId).Find(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Comment) CreateComm() (int, error) {
	db := pkg.InitDB()

	if err := db.Table("comments").Create(&c).Error; err != nil {
		return 0, err
	}
	return c.ID, nil
}

func NewComm(postId int, name, email, body string) Comment {
	return Comment{
		PostID: postId,
		Name:   name,
		Email:  email,
		Body:   body,
	}
}

func (c *Comment) UpdateComm(id int) error {
	db := pkg.InitDB()

	if err := db.Table("comments").Where("id", id).Updates(&c).Error; err != nil {
		return err
	}

	return nil
}

func NewUpdateComm(name, email, body string) Comment {
	return Comment{
		Name:  name,
		Email: email,
		Body:  body,
	}
}

func (c *Comment) DeleteComm(id int) error {
	db := pkg.InitDB()
	if err := db.Table("comments").Delete(&c, id).Error; err != nil {
		return err
	}

	return nil

}
