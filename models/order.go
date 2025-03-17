package models

// CartItem 結構表示一個購物車項目
// 包含產品ID和數量

type CartItem struct {
	ProductID int
	Quantity  int
}

// Order 結構表示一個訂單
// 包含訂單ID、用戶名、產品名稱、數量和時間戳

type Order struct {
	ID          int
	Username    string
	ProductName string
	Quantity    int
	Timestamp   string
	CartItems   []CartItem
}

// Orders 是一個模擬的訂單數據庫
var Orders []Order 