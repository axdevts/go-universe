package fetch_models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-"`
	Posts    []Post
}

type Post struct {
	CategoryID uint   `gorm:"foreignkey:CategoryID" json:"categoryID"`
	Title      string `gorm:"not null" json:"title"`
	Body       string `gorm:"type:text" json:"body"`
	UserID     uint   `gorm:"foreignkey:UserID" json:"userID"`
	// Category   uint   `gorm:"foreignkey:CategoryID"`
	User User `gorm:"foreignkey:UserID"`
	// Comments   []Comment
}

type PostData struct {
	Data        []Post `json:"data"`
	CurrentPage uint   `json:"current_page"`
	From        uint   `json:"from"`
	To          uint   `json:"to"`
	LastPage    uint   `json:"last_page"`
	PerPage     uint   `json:"per_page"`
	Total       uint   `json:"total"`
}

type PostResponse struct {
	Response PostData `json:"response"`
}
