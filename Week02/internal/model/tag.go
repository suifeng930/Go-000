package model

import (
	"database/sql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State int64  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"

}

func (t Tag) CreateTag(db *gorm.DB) error {

	err := db.Create(&t).Error
	if err != nil {
		return errors.Wrap(err," insert into tag is fail ")
	}
	return nil
}

func (t Tag) DeleteTagByID(db *gorm.DB) error  {

	err := db.Where("id=? and is_del=?", t.ID, 0).Delete(&t).Error
	if err!=nil {

		return errors.Wrap(err,"delete tag line is fail .")
	}
	return nil
}

func (t Tag) UpdateTagByID(db *gorm.DB,values interface{}) error  {

	err :=db.Model(&Tag{}).Where("id=? and is_del=?",t.ID,0).Updates(values).Error
	if err!=nil {
		return errors.Wrap(err,"update tag by id  is fail .")

	}
	return nil


}

func (t Tag) GetTag(db *gorm.DB) (*Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? ", t.ID, 0).First(&tag).Error
	if err != nil && errors.Is(err,sql.ErrNoRows) {
		return &tag, errors.Wrap(gorm.ErrRecordNotFound,"get tag by id  is fail .")
	}

	return &tag, nil
}