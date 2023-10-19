package entity

import uuid "github.com/satori/go.uuid"

// Picture object for REST(CRUD)
type Picture struct {
	ID       uuid.UUID `json:"id" gorm:"column:id;type:uuid"`
	Type     string    `json:"type" gorm:"column:type"`
	Album    string    `json:"album" gorm:"column:album"`
	Filename string    `json:"filename" gorm:"column:filename"`
	Tags     []*Tags   `json:"tags" gorm:"many2many:picture_tags;column:tags"`
}

type Tags struct {
	ID  uuid.UUID `json:"id" gorm:"column:id;type:uuid"`
	Tag string    `json:"tag" gorm:"column:tag"`
}

func (Picture) TableName() string {
	return "react-pictures"
}
