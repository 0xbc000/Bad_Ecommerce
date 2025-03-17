package controllers

import (
	"net/http"
	"html/template"
	"bad_ecommerce/models"
)

// AdminHandler 處理管理面板請求
// API Name: Admin
// Description: 提供管理員功能的入口。
// Parameters: 無
// Returns: 返回管理面板。
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if sessionCookie, err := r.Cookie("session"); err == nil {
		if sessionCookie.Value != "admin" {
			http.Error(w, "未授權的訪問", http.StatusUnauthorized)
			return
		}
	}

	tmpl, _ := template.New("admin").Parse(`
	<html>
	<head>
		<title>管理面板</title>
		<style>
			body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #333333; margin: 0; padding: 0; }
			.container { width: 80%; margin: 20px auto; padding: 20px; background-color: #444444; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
			button { padding: 10px 20px; background-color: #FF9900; color: #ffffff; border: none; cursor: pointer; margin: 5px; }
			button:hover { background-color: #e68a00; }
			h1 { color: #ffffff; }
		</style>
	</head>
	<body>
		<div class='container'>
			<h1>管理面板</h1>
			<button onclick="location.href='/orders'">查看訂單</button>
			<button onclick="location.href='/users'">管理用戶</button>
		</div>
	</body>
	</html>`)
	tmpl.Execute(w, nil)
}

func AdminViewOrdersHandler(w http.ResponseWriter, r *http.Request) {
	var allOrders []models.Order
	allOrders = models.Orders

	productNames := make(map[int]string)
	for _, product := range models.Products {
		productNames[product.ID] = product.Name
	}

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
				{{range .Orders}}
				{{range .CartItems}}
				<tr>
					<td>{{$.Username}}</td>
					<td>{{index $.ProductNames .ProductID}}</td>
					<td>{{.Quantity}}</td>
					<td>{{$.Timestamp}}</td>
				</tr>
				{{end}}
				{{end}}
			</table>
			<a href='/admin' style='display: inline-block; margin-top: 20px; padding: 10px 20px; background-color: #FF9900; color: #ffffff; text-decoration: none;'>返回管理面板</a>
		</div>
	</body>
	</html>`)
	tmpl.Execute(w, struct {
		Orders      []models.Order
		ProductNames map[int]string
	}{allOrders, productNames})
} 