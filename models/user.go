package models

import "fmt"

// User 用戶結構
type User struct {
	ID       int
	Username string
	Password string
	IDNumber string // 身分證號碼
}

// 模擬用戶數據庫
var Users = []User{
	{ID: 1, Username: "admin", Password: "admin123", IDNumber: "A1234567890"},
	{ID: 2, Username: "user", Password: "user123", IDNumber: "B1234567890"},
	{ID: 3, Username: "user1", Password: "user1", IDNumber: "C1234567890"},
	{ID: 4, Username: "user2", Password: "user2", IDNumber: "D1234567890"},
	{ID: 5, Username: "user3", Password: "user3", IDNumber: "E1234567890"},
}

// Authenticate 驗證用戶名和密碼
func Authenticate(username, password string) (*User, error) {
	for _, user := range Users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("用戶名或密碼錯誤")
} 