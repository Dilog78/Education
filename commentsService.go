package main

import (
	"encoding/json"
	"encoding/xml"
)

func (c *Comments) getComm(postId int) ([]byte, error) {
	db := Connect()

	if err := db.Table("comments").Where("post_id", postId).Find(&c).Error; err != nil {
		return nil, err
	}
	b, _ := json.MarshalIndent(c, " ", " ")

	return b, nil
}

func (*XmlComments) getXmlComById(id int) ([]byte, error) {
	var c []*XmlComments

	db := Connect()

	if err := db.Table("comments").Where("post_id", id).Find(&c).Error; err != nil {
		return nil, err
	}

	b, err := xml.MarshalIndent(c, " ", " ")
	if err != nil {
		return nil, err
	}

	return b, nil
}
