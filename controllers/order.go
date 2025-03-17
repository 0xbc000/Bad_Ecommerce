package controllers

import (
	"net/http"
	"html/template"
	"strconv"
	"time"
	"bad_ecommerce/models"
)

// OrderDetail 訂單詳細結構
// 包含所有必要的訂單信息
// 用於 /order, /order-history, /orders
type OrderDetail struct {
	OrderID     int
	UserID      int
	Username    string
	ProductName string
	Quantity    int
	Timestamp   string
}

// 模擬訂單數據
var orderDetails = []OrderDetail{
	{OrderID: 1, UserID: 1, Username: "user1", ProductName: "西瓜", Quantity: 2, Timestamp: "2023-10-01 10:00:00"},
	{OrderID: 2, UserID: 1, Username: "user1", ProductName: "蘋果", Quantity: 1, Timestamp: "2023-10-02 11:00:00"},
	{OrderID: 3, UserID: 2, Username: "user2", ProductName: "葡萄", Quantity: 3, Timestamp: "2023-10-03 12:00:00"},
	{OrderID: 4, UserID: 2, Username: "user2", ProductName: "芒果", Quantity: 1, Timestamp: "2023-10-04 13:00:00"},
	{OrderID: 5, UserID: 3, Username: "user3", ProductName: "檸檬", Quantity: 2, Timestamp: "2023-10-05 14:00:00"},
	{OrderID: 6, UserID: 3, Username: "user3", ProductName: "可樂", Quantity: 1, Timestamp: "2023-10-06 15:00:00"},
}

// PurchaseHandler 處理購買請求
// API Name: Purchase
// Description: 處理用戶的購買請求，生成訂單。
// Parameters: 無
// Returns: 購買成功後重定向到訂單確認頁面。
func PurchaseHandler(w http.ResponseWriter, r *http.Request) {
	if sessionCookie, err := r.Cookie("session"); err == nil {
		for _, item := range cart {
			newOrderDetail := OrderDetail{
				OrderID:     len(orderDetails) + 1,
				UserID:      0, // 假設用戶ID為0，需根據實際情況更新
				Username:    sessionCookie.Value,
				ProductName: models.Products[item.ProductID-1].Name,
				Quantity:    item.Quantity,
				Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
			}
			orderDetails = append(orderDetails, newOrderDetail)
		}
		cart = []CartItem{}
		http.Redirect(w, r, "/order?order_id="+strconv.Itoa(len(orderDetails)), http.StatusSeeOther)
		return
	}
	http.Error(w, "請先註冊或登入", http.StatusUnauthorized)
}

// ViewOrdersHandler 處理查看訂單請求
// API Name: View Orders
// Description: 顯示所有用戶的所有訂單。
// Parameters: 無
// Returns: 返回訂單列表頁面。
func ViewOrdersHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("orders").Parse(`
	<html>
	<head>
		<title>訂單列表</title>
		<style>
			body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #333333; margin: 0; padding: 0; }
			.container { width: 80%; margin: 20px auto; padding: 20px; background-color: #444444; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
			h1, p, div { color: #ffffff; }
			table { width: 100%; border-collapse: collapse; margin-top: 20px; }
			th, td { padding: 10px; text-align: left; border-bottom: 1px solid #ddd; color: #ffffff; }
			th { background-color: #FF9900; color: #ffffff; }
			tr:hover { background-color: #555555; }
			button { padding: 10px 20px; background-color: #FF9900; color: #ffffff; border: none; cursor: pointer; margin: 5px; }
			button:hover { background-color: #e68a00; }
			.order-item { background-color: #444444; margin: 10px 0; padding: 10px; border-radius: 5px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
		</style>
	</head>
	<body>
		<div class='container'>
			<h1>訂單列表</h1>
			<table>
				<tr>
					<th>用戶名</th>
					<th>產品名稱</th>
					<th>數量</th>
					<th>時間</th>
				</tr>
				{{range .}}
				<tr>
					<td>{{.Username}}</td>
					<td>{{.ProductName}}</td>
					<td>{{.Quantity}}</td>
					<td>{{.Timestamp}}</td>
				</tr>
				{{end}}
			</table>
			<a href='/admin' style='display: inline-block; margin-top: 20px; padding: 10px 20px; background-color: #FF9900; color: #ffffff; text-decoration: none;'>返回管理面板</a>
		</div>
	</body>
	</html>`)
	tmpl.Execute(w, orderDetails)
}

// ViewOrderDetailHandler 處理查看訂單詳情請求
// API Name: View Order Detail
// Description: 顯示指定訂單的詳細信息。
// Parameters: orderID (URL query)
// Returns: 返回訂單詳情頁面。
func ViewOrderDetailHandler(w http.ResponseWriter, r *http.Request) {
	orderID, _ := strconv.Atoi(r.URL.Query().Get("order_id"))
	for _, orderDetail := range orderDetails {
		if orderDetail.OrderID == orderID {
			tmpl, _ := template.New("orderDetail").Parse(`
			<html>
			<head>
			<title>訂單詳情</title>
			<style>
			body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #333333; margin: 0; padding: 0; }
			.container { width: 80%; margin: 20px auto; }
			.order-detail { list-style-type: none; padding: 0; }
			.order-item { background-color: #444444; margin: 10px 0; padding: 10px; border-radius: 5px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
			h1, p, div { color: #ffffff; }
			</style>
			</head>
			<body>
			<div class='container'>
			<h1>訂單詳情</h1>
			<ul class='order-detail'>
			<li class='order-item'>產品名稱: {{.ProductName}}, 數量: {{.Quantity}}</li>
			</ul>
			<p>訂單時間: {{.Timestamp}}</p>
			<a href='/home'>返回首頁</a>
			</div>
			</body>
			</html>`)
			tmpl.Execute(w, orderDetail)
			return
		}
	}
	http.Error(w, "請先註冊或登入", http.StatusUnauthorized)
}

// OrderHistoryHandler 處理查看訂單歷史請求
// API Name: Order History
// Description: 顯示用戶的訂單歷史。
// Parameters: 無
// Returns: 返回訂單歷史頁面。
func OrderHistoryHandler(w http.ResponseWriter, r *http.Request) {
	username := ""
	if cookie, err := r.Cookie("session"); err == nil {
		username = cookie.Value
	}

	userOrderDetails := []OrderDetail{}
	for _, orderDetail := range orderDetails {
		if orderDetail.Username == username {
			userOrderDetails = append(userOrderDetails, orderDetail)
		}
	}

	tmpl, _ := template.New("orderHistory").Parse(`
	<html>
	<head>
	<title>訂單歷史</title>
	<style>
	body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #333333; margin: 0; padding: 0; }
	.container { width: 80%; margin: 20px auto; padding: 20px; background-color: #444444; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
	h1, p, div { color: #ffffff; }
	.order-item { background-color: #444444; margin: 10px 0; padding: 10px; border-radius: 5px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
	</style>
	</head>
	<body>
	<div class='container'>
	<h1>訂單歷史</h1>
	<ul>
	{{range .}}
	<li class='order-item'>訂單ID: {{.OrderID}}, 產品名稱: {{.ProductName}}, 數量: {{.Quantity}}, 時間: {{.Timestamp}}</li>
	{{end}}
	</ul>
	<a href='/home' style='display: inline-block; margin-top: 20px; padding: 10px 20px; background-color: #FF9900; color: #ffffff; text-decoration: none;'>返回首頁</a>
	</div>
	</body>
	</html>`)
	tmpl.Execute(w, userOrderDetails)
}
