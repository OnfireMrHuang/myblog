package models

type Comment struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Content  string `json:"content"`
	Time     string `json:"time"`
	IP       string `json:"ip"`
}

func GetComments(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Comment").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}
