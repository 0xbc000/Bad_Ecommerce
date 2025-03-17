package controllers

import (
	"net/http"
	"html/template"
	"strconv"
	"bad_ecommerce/models"
)

// LoginHandler 處理用戶登入請求
// API Name: Login
// Description: 處理用戶登入，驗證用戶憑證。
// Parameters: username 和 password (POST 請求時需要)。
// Returns: 登入成功後重定向到首頁，失敗則返回錯誤信息。

// EnhancedProductSearchHandler 處理產品搜索請求
// API Name: Enhanced Product Search
// Description: 提供增強的產品搜索功能。
// Parameters: search query (URL query)
// Returns: 返回符合搜索條件的產品列表。

// ViewCartHandler 顯示用戶購物車
// API Name: View Cart
// Description: 顯示當前用戶的購物車內容。
// Parameters: 無
// Returns: 返回購物車頁面。

// AddToCartHandler 添加商品到購物車
// API Name: Add to Cart
// Description: 將選定的商品添加到用戶的購物車。
// Parameters: productID (POST 請求時需要)
// Returns: 添加成功後重定向到購物車頁面。

// PurchaseHandler 處理購買請求
// API Name: Purchase
// Description: 處理用戶的購買請求，生成訂單。
// Parameters: 無
// Returns: 購買成功後重定向到訂單確認頁面。

// ViewOrdersHandler 顯示用戶訂單
// API Name: View Orders
// Description: 顯示當前用戶的所有訂單。
// Parameters: 無
// Returns: 返回訂單列表頁面。

// ViewOrderDetailHandler 顯示訂單詳情
// API Name: View Order Detail
// Description: 顯示指定訂單的詳細信息。
// Parameters: orderID (URL query)
// Returns: 返回訂單詳情頁面。

// AdminHandler 處理管理員請求
// API Name: Admin
// Description: 提供管理員功能的入口。
// Parameters: 無
// Returns: 返回管理員頁面。

// PaymentHandler 處理支付請求
// API Name: Payment
// Description: 處理用戶的支付請求。
// Parameters: payment details (POST 請求時需要)
// Returns: 支付成功後重定向到確認頁面。

// LogoutHandler 處理用戶登出請求
// API Name: Logout
// Description: 處理用戶登出，清除 session。
// Parameters: 無
// Returns: 登出成功後重定向到登入頁面。

// OrderHistoryHandler 顯示訂單歷史
// API Name: Order History
// Description: 顯示用戶的訂單歷史。
// Parameters: 無
// Returns: 返回訂單歷史頁面。

// UserHandler 處理用戶管理請求
// API Name: User Management
// Description: 顯示用戶列表，並允許創建新用戶。
// Parameters: POST 請求時需要 username 和 password。
// Returns: GET 請求返回用戶列表，POST 請求成功後重定向到用戶列表。
func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// 創建新用戶
		newUser := models.User{
			ID:       len(models.Users) + 1,
			Username: username,
			Password: password,
		}
		models.Users = append(models.Users, newUser)

		// 重定向到用戶列表
		http.Redirect(w, r, "/users", http.StatusSeeOther)
		return
	}

	tmpl, _ := template.New("user").Parse(`
	<html>
	<head>
		<title>用戶管理</title>
		<style>
			body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #333333; margin: 0; padding: 0; }
			.container { width: 80%; margin: 20px auto; padding: 20px; background-color: #444444; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
			h1, p, div { color: #cccccc; }
			table { width: 100%; border-collapse: collapse; margin-top: 20px; }
			th, td { padding: 10px; text-align: left; border-bottom: 1px solid #ddd; }
			th { background-color: #FF9900; color: #ffffff; }
			tr:hover { background-color: #555555; }
			button { padding: 10px 20px; background-color: #FF9900; color: #ffffff; border: none; cursor: pointer; margin: 5px; }
			button:hover { background-color: #e68a00; }
			input[type='submit'] { width: 100%; padding: 10px; background-color: #FF9900; color: #ffffff; border: none; cursor: pointer; }
			input[type='submit']:hover { background-color: #e68a00; }
			a { display: inline-block; margin-top: 20px; padding: 10px 20px; background-color: #FF9900; color: #ffffff; text-decoration: none; }
			a:hover { background-color: #e68a00; }
			.username { color: #ffffff; }
		</style>
	</head>
	<body>
		<div class='container'>
			<h1>用戶管理</h1>
			<table>
				<tr>
					<th>用戶名</th>
					<th>操作</th>
				</tr>
				{{range .Users}}
				<tr>
					<td class='username'>{{.Username}}</td>
					<td><button onclick="location.href='/users/edit?id={{.ID}}'">編輯</button></td>
				</tr>
				{{end}}
			</table>
			<a href='/admin' style='display: inline-block; margin-top: 20px; padding: 10px 20px; background-color: #FF9900; color: #fff; text-decoration: none;'>返回管理面板</a>
		</div>
	</body>
	</html>`)
	tmpl.Execute(w, struct{ Users []models.User }{Users: models.Users})
}

// EditUserHandler 處理用戶編輯請求
// API Name: Edit User
// Description: 編輯指定用戶的信息。
// Parameters: id (URL query), username 和 password (POST 請求時需要)。
// Returns: 編輯成功後重定向到用戶列表。
func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	for i, user := range models.Users {
		if user.ID == id {
			if r.Method == http.MethodPost {
				models.Users[i].Username = r.FormValue("username")
				models.Users[i].Password = r.FormValue("password")
				http.Redirect(w, r, "/users", http.StatusSeeOther)
				return
			}

			tmpl, _ := template.New("editUser").Parse(`
			<html>
			<head>
				<title>編輯用戶</title>
				<style>
					body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #333333; margin: 0; padding: 0; }
					.container { width: 300px; margin: 100px auto; padding: 20px; background-color: #444444; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
					h2 { text-align: center; }
					input[type='text'], input[type='password'] { width: 100%; padding: 10px; margin: 10px 0; }
					input[type='submit'] { width: 100%; padding: 10px; background-color: #FF9900; color: #ffffff; border: none; cursor: pointer; }
					input[type='submit']:hover { background-color: #e68a00; }
				</style>
			</head>
			<body>
				<div class='container'>
					<h2>編輯用戶</h2>
					<form method='post'>
						<input name='username' value='{{.Username}}' type='text' placeholder='用戶名'>
						<input name='password' value='{{.Password}}' type='password' placeholder='密碼'>
						<input type='submit' value='保存'>
					</form>
				</div>
			</body>
			</html>`)
			tmpl.Execute(w, user)
			return
		}
	}
	http.NotFound(w, r)
}

// DeleteUserHandler 處理用戶刪除請求
// API Name: Delete User
// Description: 刪除指定用戶。
// Parameters: id (URL query)
// Returns: 刪除成功後重定向到用戶列表。
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	for i, user := range models.Users {
		if user.ID == id {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			http.Redirect(w, r, "/users", http.StatusSeeOther)
			return
		}
	}
	http.NotFound(w, r)
}

// SignupHandler 處理用戶註冊請求
// API Name: Signup
// Description: 註冊新用戶。
// Parameters: username 和 password (POST 請求時需要)。
// Returns: 註冊成功後重定向到首頁。
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		idnumber := r.FormValue("idnumber")

		// 檢查用戶名是否已存在
		for _, user := range models.Users {
			if user.Username == username {
				http.Error(w, "username exist", http.StatusConflict)
				return
			}
		}

		// 潛在漏洞：不安全的密碼存儲
		// 建議修補方式：使用 bcrypt 或類似的庫對密碼進行哈希處理後再存儲
		// 示例：hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		// if err != nil { // 處理錯誤 }

		// 創建新用戶
		newUser := models.User{
			ID:       len(models.Users) + 1,
			Username: username,
			Password: password, // 這裡應該存儲 hashedPassword
			IDNumber: idnumber,
		}
		models.Users = append(models.Users, newUser)

		// 註冊成功，重定向到首頁
		http.SetCookie(w, &http.Cookie{Name: "session", Value: username})
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	tmpl, _ := template.New("signup").Parse(`
	<html>
	<head>
		<style>
			body { background-color: #333333; font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; }
			.container { width: 300px; margin: 100px auto; padding: 20px; background-color: #444444; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
			h2 { text-align: center; }
			input[type='text'], input[type='password'] { width: 100%; padding: 10px; margin: 10px 0; }
			input[type='submit'] { width: 100%; padding: 10px; background-color: #FF9900; color: #ffffff; border: none; cursor: pointer; }
			input[type='submit']:hover { background-color: #e68a00; }
		</style>
	</head>
	<body>
		<div class="container">
			<h2>註冊</h2>
			<form method='post'>
				<input name='username' type='text' placeholder='用戶名'>
				<input name='password' type='password' placeholder='密碼'>
				<input name='idnumber' type='text' placeholder='身分證號碼'>
				<input type='submit' value='註冊'>
			</form>
		</div>
	</body>
	</html>
	`)
	tmpl.Execute(w, nil)
}

// UserDashboardHandler 處理用戶儀表板請求
// API Name: User Dashboard
// Description: 顯示用戶的帳號信息和訂單歷史鏈接。
// Parameters: 無
// Returns: 如果用戶已登入，返回用戶儀表板；否則返回未授權錯誤。
func UserDashboardHandler(w http.ResponseWriter, r *http.Request) {
	if sessionCookie, err := r.Cookie("session"); err == nil {
		username := sessionCookie.Value
		var currentUser *models.User
		for _, user := range models.Users {
			if user.Username == username {
				currentUser = &user
				break
			}
		}
		if currentUser != nil {
			tmpl, _ := template.New("userDashboard").Parse(`
			<html>
			<head>
			<title>用戶儀表板</title>
			<style>
			body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #333333; margin: 0; padding: 0; }
			.container { width: 80%; margin: 20px auto; padding: 20px; background-color: #444444; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
			h1, p, div { color: #cccccc; }
			p { font-size: 18px; }
			a { display: inline-block; margin-top: 20px; padding: 10px 20px; background-color: #FF9900; color: #ffffff; text-decoration: none; }
			a:hover { background-color: #e68a00; }
			</style>
			</head>
			<body>
			<div class='container'>
			<h1>用戶儀表板</h1>
			<p>帳號: {{.Username}}</p>
			<p>身分證號碼: {{.IDNumber}}</p>
			<a href='/order-history'>查看訂單歷史</a>
			<a href='/home'>返回首頁</a>
			</div>
			</body>
			</html>`)
			tmpl.Execute(w, currentUser)
			return
		}
	}
	http.Error(w, "請先註冊或登入", http.StatusUnauthorized)
} 