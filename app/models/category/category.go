package category

import (
	"goblog/app/models"
	"goblog/pkg/model"
	"goblog/pkg/route"
	"goblog/pkg/types"
)

// Category 文章分类
type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

func (c Category) Link() string {

	return route.Name2URL("categories.show", "id", c.GetStringID())
}

func Get(idstr string) (Category, error) {

	var category Category
	id := types.StringToUint64(idstr)
	if err := model.DB.First(&category, id).Error; err != nil {
		return category, err
	}

	return category, nil

}
