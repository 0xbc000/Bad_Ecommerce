package controllers

import (
	"net/http"
	"html/template"
	"strconv"
	"bad_ecommerce/models"
)

// CartItem 購物車項目結構
type CartItem struct {
	ProductID int
	Quantity  int
	UserID    int
}

// 模擬購物車數據
var cart = []CartItem{}

// AddToCartHandler 處理添加到購物車請求
// API Name: Add to Cart
// Description: 將選定的商品添加到用戶的購物車。
// Parameters: productID (POST 請求時需要)
// Returns: 添加成功後重定向到購物車頁面。
func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	// 檢查用戶是否已登入
	if sessionCookie, err := r.Cookie("session"); err == nil {
		userID, _ := strconv.Atoi(sessionCookie.Value)
		productID, err1 := strconv.Atoi(r.URL.Query().Get("id"))
		quantity, err2 := strconv.Atoi(r.URL.Query().Get("quantity"))

		if err1 != nil || err2 != nil || quantity <= 0 {
			http.Error(w, "無效的產品ID或數量", http.StatusBadRequest)
			return
		}

		// 添加到購物車
		for i, item := range cart {
			if item.ProductID == productID && item.UserID == userID {
				cart[i].Quantity += quantity
				http.Redirect(w, r, "/cart?user_id="+sessionCookie.Value, http.StatusSeeOther)
				return
			}
		}
		cart = append(cart, CartItem{ProductID: productID, Quantity: quantity, UserID: userID})
		http.Redirect(w, r, "/cart?user_id="+sessionCookie.Value, http.StatusSeeOther)
		return
	}
	http.Error(w, "請先註冊或登入", http.StatusUnauthorized)
}

// ViewCartHandler 處理查看購物車請求
func ViewCartHandler(w http.ResponseWriter, r *http.Request) {
	if sessionCookie, err := r.Cookie("session"); err == nil {
		userID, _ := strconv.Atoi(sessionCookie.Value)
		userCart := []CartItem{}
		for _, item := range cart {
			if item.UserID == userID {
				userCart = append(userCart, item)
			}
		}

		cartItems := []struct {
			ProductName string
			Quantity    int
		}{}

		for _, item := range userCart {
			for _, product := range models.Products {
				if product.ID == item.ProductID {
					cartItems = append(cartItems, struct {
						ProductName string
						Quantity    int
					}{
						ProductName: product.Name,
						Quantity:    item.Quantity,
					})
				}
			}
		}

		tmpl, _ := template.New("cart").Parse(`
		<html>
		<head>
		<title>購物車</title>
		<style>
		body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #333333; margin: 0; padding: 0; }
		.container { width: 80%; margin: 20px auto; background-color: #FF9900; padding: 20px; border-radius: 5px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
		.cart-list { list-style-type: none; padding: 0; }
		.cart-item { background-color: #444444; margin: 10px 0; padding: 10px; border-radius: 5px; box-shadow: 0 0 10px rgba(0,0,0,0.1); display: flex; align-items: center; }
		.cart-item div { flex-grow: 1; }
		.cart-item a { color: #FF9900; text-decoration: none; }
		.cart-item a:hover { text-decoration: underline; }
		div { color: #ffffff; }
		a { display: inline-block; margin-top: 20px; padding: 10px 20px; background-color: #cc7a00; color: #ffffff; text-decoration: none; border-radius: 5px; }
		a:hover { background-color: #b36b00; }
		input[type='submit'] { width: 100%; padding: 10px; background-color: #cc7a00; color: #ffffff; border: none; cursor: pointer; border-radius: 5px; }
		input[type='submit']:hover { background-color: #b36b00; }
		</style>
		</head>
		<body>
		<div class='container'>
		<h1>購物車</h1>
		<ul class='cart-list'>
		{{range .}}
		<li class='cart-item'>
		<div>{{.ProductName}} - 數量: {{.Quantity}}</div>
		</li>
		{{end}}
		</ul>
		<form action='/purchase' method='post'>
		<input type='submit' value='購買'>
		</form>
		<a href='/home'>返回首頁</a>
		</div>
		</body>
		</html>`)
		tmpl.Execute(w, cartItems)
		return
	}
	http.Error(w, "請先註冊或登入", http.StatusUnauthorized)
} 