package dto

type User struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement;index" query:"id"`
	Username   string `json:"username" gorm:"unique" query:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname,omitempty"`
	Permission int    `json:"permission" gorm:"default:0"`
	Followed   int    `json:"followed" gorm:"default:0"`  // Number of people followed
	Followers  int    `json:"followers" gorm:"default:0"` // Number of followers
	Oauth      bool   `json:"oauth" gorm:"default:false"` // Whether it is a third-party login
}
