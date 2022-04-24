package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

//Article文章模型
type Article struct {
	models.BaseModel

	Title string `gorm:"type:varchar(255);not nul;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`
}

func (article Article) Link() string {

	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}
