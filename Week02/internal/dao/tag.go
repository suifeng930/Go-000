package dao

import (
	"github.com/Go-000/Week02/internal/model"
)

func (d *Dao)CreateTag(name string,state int64, createBy string)error  {
	tag := model.Tag{

		Model: &model.Model{CreatedBy: createBy},
		Name:  name,
		State: state,
	}
	return tag.CreateTag(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state int64, modifiedBy string) error {
	tag :=model.Tag{
		Model:&model.Model{ID: id},
	}
	values :=map[string]interface{}{
		"state":state,
		"modified_by":modifiedBy,
	}
	if name!="" {
		values["name"]=name

	}
	return tag.UpdateTagByID(d.engine,values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag :=model.Tag{
		Model:&model.Model{ID: id},
	}
	return tag.DeleteTagByID(d.engine)
}

func (d *Dao) GetTag(id uint32) (*model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.GetTag(d.engine)
}