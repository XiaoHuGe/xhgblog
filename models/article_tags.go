package models

//table article_tags
type ArticleTags struct {
	//Model
	ArticleId int `json:"article_id"` // Article id
	TagId     int `json:"tag_id"`     // tag id
}

func DeleteTagsByArticleId(id int) error {
	err := db.Where("article_id = ?", id).Delete(&ArticleTags{}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteTagsByTagId(id int) error {
	err := db.Where("tag_id = ?", id).Delete(&ArticleTags{}).Error
	if err != nil {
		return err
	}
	return nil
}

func AddArticleJoinTags(articleId, tagId int) error {
	err := db.Create(ArticleTags{
		ArticleId: articleId,
		TagId:     tagId,
	}).Error
	return err
}

func GetArticleTotalByTagId(tagId int) (int, error) {
	var articleTags []*ArticleTags
	var count int
	err := db.Where("tag_id = ?", tagId).Find(&articleTags).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
