package models

//table article_tags
type ArticleTags struct {
	//Model
	ArticleId int `json:"article_id"` // Article id
	TagId     int `json:"tag_id"`     // tag id
}

func DeleteTagsByArticleId(id int) (error) {
	err := db.Where("article_id = ?", id).Delete(&ArticleTags{}).Error
	if err != nil {
		return err
	}
	return nil
}

func AddArticleJoinTags(article_id, tag_id int) (error) {
	err := db.Create(ArticleTags{
		ArticleId: article_id,
		TagId:     tag_id,
	}).Error
	return err
}
