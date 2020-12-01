package service

import (
	"github.com/Go-000/Week02/internal/model"
	"github.com/pkg/errors"
	"log"
)

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type GetTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}


func (svc *Service) GetTag(param *GetTagRequest) (*model.Tag, error) {

	tag, err := svc.dao.GetTag(param.ID)
	if err != nil {
		return nil, errors.WithMessage(err," can not found  tag ")
	}
	return tag,nil

}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	err := svc.dao.CreateTag(param.Name, int64(param.State), param.CreatedBy)
	log.Println("CreateTag(param *CreateTagRequest) :", err)
	return err

}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, int64(param.State), param.ModifiedBy)

}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)

}
