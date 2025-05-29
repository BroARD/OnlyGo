package service

type Quote struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Author string `json:"author"`
	Quote string `json:"quote"`
}
