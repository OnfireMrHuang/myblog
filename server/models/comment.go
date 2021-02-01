package models

type Comment struct {
	Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Content  string `json:"content"`
	IP       string `json:"ip"`
	ArticleID int `json:"article_id"`
}

func GetComments(pageNum int, pageSize int, maps interface{}) (comments []Article) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&comments)
	return
}

func AddComment(data map[string]interface{}) error {
	db.Create(&Comment{
		ArticleID:     data["article_id"].(int),
		Username:     data["username"].(string),
		Email:      data["email"].(string),
		Content:   data["content"].(string),
		IP:     data["ip"].(string),
	})
	return nil
}

func DeleteComment(id int) bool {
	db.Where("id = ?", id).Delete(Comment{})
	return true
}
