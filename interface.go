package main

import (
	"encoding/json"
)

type CrudMethods interface {
	Create(Post) error
	DeleteById([]int) error
	GetById(int) ([]byte, error)
	UpdateById(int, string, string) error
	getAll() ([]byte, error)
}

func (p *Post) Create(target Post) error {

	db := Connect()

	if err := db.Table("posts").Create(&target).Error; err != nil {
		return err
	}

	return nil
}

func (p *Post) DeleteById(id []int) error {

	db := Connect()

	if err := db.Table("posts").Delete(&p, id).Error; err != nil {
		return err
	}

	return nil
}

func (p *Post) GetById(id int) ([]byte, error) {
	db := Connect()

	err := db.Table("posts").Where("id", id).Take(&p).Error
	if err != nil {
		return nil, err
	}

	res, _ := json.Marshal(p)

	return res, nil
}

func (p *Post) UpdateById(id int, target string, value string) error {

	db := Connect()
	if err := db.Model(&p).Table("posts").Where("id", id).Update(target, value).Error; err != nil {
		return err
	}

	return nil
}

func (p *Post) getAll() ([]byte, error) {
	db := Connect()

	if err := db.Table("posts").Find(&p).Error; err != nil {
		return nil, err
	}
	x, err := json.MarshalIndent(p, " ", " ")
	if err != nil {
		return nil, err
	}

	return x, nil
}
