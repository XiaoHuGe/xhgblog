package service

import "xhgblog/models"

func DeleteTagsByArticleId(id int) error {
	return models.DeleteTagsByArticleId(id)
}
