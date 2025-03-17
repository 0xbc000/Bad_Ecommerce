package controllers

import (
	"net/http"
	"html/template"
)

// PaymentHandler 處理支付請求
// API Name: Payment
// Description: 處理用戶的支付請求。
// Parameters: 無
// Returns: 支付成功後重定向到訂單歷史頁面。
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// 模擬支付成功
		http.Redirect(w, r, "/order-history", http.StatusSeeOther)
		return
	}

	tmpl, _ := template.New("payment").Parse(`<html><body><h1>支付頁面</h1><form method='post'><input type='submit' value='確認支付'></form></body></html>`)
	tmpl.Execute(w, nil)
} 