package controllers

import (
	"net/http"
	"html/template"
	"bad_ecommerce/models"
)

// HomeHandler 處理首頁請求
// API Name: Home
// Description: 顯示首頁，根據用戶的 session 狀態顯示不同內容。
// Parameters: search (URL query)
// Returns: 返回首頁頁面，顯示產品列表和用戶狀態。
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	username := ""
	if cookie, err := r.Cookie("session"); err == nil {
		username = cookie.Value
	}

	searchQuery := r.URL.Query().Get("search")

	tmpl, _ := template.New("home").Parse(`
	<html>
	<head>
		<title>首頁</title>
		<style>
			body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #333333; margin: 0; padding: 0; }
			.header { background-color: #FF9900; color: #ffffff; padding: 10px 0; text-align: center; }
			.container { width: 80%; margin: 20px auto; }
			.product-list { list-style-type: none; padding: 0; }
			.product-item { background-color: #444444; margin: 10px 0; padding: 10px; border-radius: 5px; box-shadow: 0 0 10px rgba(0,0,0,0.1); display: flex; align-items: center; }
			.product-item img { width: 50px; height: 50px; margin-right: 20px; }
			.product-item a { margin-left: auto; color: #FF9900; text-decoration: none; }
			.product-item a:hover { text-decoration: underline; }
			.button { background-color: #FF9900; color: #ffffff; padding: 5px 10px; border-radius: 5px; text-decoration: none; display: inline-block; }
			.button img { vertical-align: middle; margin-right: 5px; }
			h1, h2, p, div { color: #ffffff; }
		</style>
	</head>
	<body>
		<div class='header'>
			<div style='float: right; margin-right: 20px;'>
				{{if .Username}}
					<a class='button' href='/user'><img src='../images/user.png' alt='User'>User Dashboard</a> | <a class='button' href='/logout'><img src='../images/logout.png' alt='Logout'>登出</a> | <a class='button' href='/cart'><img src='../images/cart.png' alt='Cart'>購物車</a>
				{{else}}
					<a class='button' href='/signup'><img src='../images/signup.png' alt='Signup'>註冊</a> | <a class='button' href='/login'><img src='../images/login.png' alt='Login'>登入</a>
				{{end}}
			</div>
			<h1>Insecure E-commerce Website</h1>
		</div>
		<div class='container'>
			<h2>Product List</h2>
			<form method='get' action='/home'>
				<input type='text' name='search' placeholder='搜索產品...'>
				<input type='submit' value='搜索'>
			</form>
			{{if .SearchQuery}}
				<p>搜索結果: {{.SearchQuery}}</p> <!-- 警告：這裡存在 XSS 漏洞，應使用適當的輸入驗證和轉義來防止攻擊 -->
			{{end}}
			<ul class='product-list'>
				{{range .Products}}
					<li class='product-item'>
						<img src='../images/{{.ID}}.jpg' alt='{{.Name}}'>
						<div>{{.Name}} - {{.Category}} - ${{.Price}}</div>
						<a href='/cart/add?id={{.ID}}&quantity=1'>加入購物車</a>
					</li>
				{{end}}
			</ul>
		</div>
	</body>
	</html>
	`)

	data := struct {
		Username    string
		Products    []models.Product
		SearchQuery template.HTML
	}{
		Username:    username,
		Products:    models.Products,
		SearchQuery: template.HTML(searchQuery),
	}

	tmpl.Execute(w, data)
} 