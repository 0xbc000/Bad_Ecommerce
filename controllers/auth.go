package controllers

import (
	"net/http"
	"html/template"
	"bad_ecommerce/models"
)

// LoginHandler 處理用戶登錄請求
// API Name: Login
// Description: 處理用戶登錄，驗證用戶憑證。
// Parameters: username / password (POST 請求時需要)
// Returns: 登錄成功後重定向到首頁，失敗則返回錯誤信息。
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		
		_, err := models.Authenticate(username, password)
		if err != nil {
			http.Error(w, "登錄失敗", http.StatusUnauthorized)
			return
		}

		// 登錄成功，重定向到首頁或管理面板
		http.SetCookie(w, &http.Cookie{Name: "session", Value: username})
		if username == "admin" {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	tmpl, _ := template.New("login").Parse(`
	<html>
	<head>
		<style>
			body { background-color: #333333; font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; }
			.container { width: 300px; margin: 100px auto; padding: 20px; background-color: #444444; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
			h2 { text-align: center; color: #ffffff; }
			input[type='text'], input[type='password'] { width: 100%; padding: 10px; margin: 10px 0; }
			input[type='submit'] { width: 100%; padding: 10px; background-color: #FF9900; color: #ffffff; border: none; cursor: pointer; }
			input[type='submit']:hover { background-color: #e68a00; }
		</style>
	</head>
	<body>
		<div class="container">
			<h2>登錄</h2>
			<form method='post'>
				<input name='username' type='text' placeholder='用戶名'>
				<input name='password' type='password' placeholder='密碼'>
				<input type='submit' value='登錄'>
			</form>
		</div>
	</body>
	</html>
	`)
	tmpl.Execute(w, nil)
}

// LogoutHandler 處理用戶登出請求
// API Name: Logout
// Description: 處理用戶登出，清除 session。
// Parameters: 無
// Returns: 登出成功後重定向到首頁。
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// 清除Session Cookie
	http.SetCookie(w, &http.Cookie{Name: "session", Value: "", MaxAge: -1})
	// 重定向到首頁並顯示登出成功信息
	http.Redirect(w, r, "/home?message=logout_success", http.StatusSeeOther)
}