package main

import (
	"encoding/xml"
	"log"
)

func (*XmlPost) getXml() ([]byte, error) {
	db := Connect()
	var x []XmlPost
	if err := db.Table("posts").Find(&x).Error; err != nil {
		log.Fatal(err)
	}

	s, err := xml.MarshalIndent(x, " ", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return s, nil
}

func (x *XmlPost) getXmlById(id int) ([]byte, error) {
	db := Connect()

	if err := db.Table("posts").Where("id", id).Take(&x).Error; err != nil {
		return nil, err
	}

	b, _ := xml.MarshalIndent(x, " ", " ")
	return b, nil
}
