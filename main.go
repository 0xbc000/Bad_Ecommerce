package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"bad_ecommerce/controllers"
	"html/template"
)

func main() {
	r := mux.NewRouter()

	// 設置路由
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", controllers.LoginHandler)
	r.HandleFunc("/users", controllers.UserHandler)
	r.HandleFunc("/users/edit", controllers.EditUserHandler)
	r.HandleFunc("/users/delete", controllers.DeleteUserHandler)
	r.HandleFunc("/search", controllers.EnhancedProductSearchHandler)
	r.HandleFunc("/cart", controllers.ViewCartHandler)
	r.HandleFunc("/cart/add", controllers.AddToCartHandler)
	r.HandleFunc("/purchase", controllers.PurchaseHandler)
	r.HandleFunc("/order", controllers.ViewOrderDetailHandler)
	r.HandleFunc("/orders", controllers.ViewOrdersHandler)
	r.HandleFunc("/admin", controllers.AdminHandler)
	r.HandleFunc("/home", controllers.HomeHandler)
	r.HandleFunc("/signup", controllers.SignupHandler)
	r.HandleFunc("/payment", controllers.PaymentHandler)
	r.HandleFunc("/logout", controllers.LogoutHandler)
	r.HandleFunc("/user", controllers.UserDashboardHandler)
	r.HandleFunc("/order-history", controllers.OrderHistoryHandler)
	r.HandleFunc("/admin/orders", controllers.AdminViewOrdersHandler)

	// 根路徑重定向到/home
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	})

	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./images/"))))

	// 啟動HTTPS服務器
	http.ListenAndServeTLS(":443", "cert.pem", "key.pem", r)
}

// HomeHandler 處理首頁請求
// API Name: Home
// Description: 顯示首頁，根據用戶的 session 狀態顯示不同內容。
// Parameters: 無
// Returns: 如果用戶已登入，返回首頁；否則重定向到登入頁面。
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("session"); err == nil {
		tmpl, _ := template.New("home").Parse(`
		<html>
		<head>
		<title>首頁</title>
		<style>
		body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #333333; margin: 0; padding: 0; }
		.container { width: 80%; margin: 20px auto; }
		button { padding: 10px 20px; background-color: #FF9900; color: #ffffff; border: none; cursor: pointer; }
		button:hover { background-color: #e68a00; }
		</style>
		</head>
		<body>
		<div class='container'>
		<h1>歡迎來到電商網站</h1>
		<button onclick="location.href='/user'">User Dashboard</button>
		<button onclick="location.href='/logout'">Logout</button>
		<button onclick="location.href='/cart'">Cart</button>
		<button onclick="location.href='/login'">Login</button>
		<button onclick="location.href='/signup'">Signup</button>
		</div>
		</body>
		</html>`)
		tmpl.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
