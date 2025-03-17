package controllers

import (
	"net/http"
	"html/template"
	"strconv"
	"strings"
	"bad_ecommerce/models"
)

// ProductSearchHandler 處理產品搜索請求
func ProductSearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	// 模擬XSS漏洞
	tmpl, _ := template.New("search").Parse(`<html><body><form method='get'><input name='q' value='` + query + `'><input type='submit'></form><div>搜索結果：` + query + `</div></body></html>`)
	tmpl.Execute(w, nil)
}

// ProductDetailHandler 處理產品詳情請求
func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	productID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	for _, product := range models.Products {
		if product.ID == productID {
			tmpl, _ := template.New("productDetail").Parse(`<html><body><h1>{{.Name}}</h1><p>價格: ${{.Price}}</p><p>這裡是產品描述。</p></body></html>`)
			tmpl.Execute(w, product)
			return
		}
	}
	http.NotFound(w, r)
}

// EnhancedProductSearchHandler 處理增強的產品搜索請求
// API Name: Enhanced Product Search
// Description: 提供增強的產品搜索功能。
// Parameters: search query (URL query)
// Returns: 返回符合搜索條件的產品列表。
func EnhancedProductSearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	category := r.URL.Query().Get("category")

	filteredProducts := []models.Product{}
	for _, product := range models.Products {
		if (query == "" || contains(product.Name, query)) && (category == "" || product.Category == category) {
			filteredProducts = append(filteredProducts, product)
		}
	}

	tmpl, _ := template.New("search").Parse(`<html><body><form method='get'><input name='q' value='` + query + `'><select name='category'><option value=''>所有分類</option><option value='水果'>水果</option><option value='飲料'>飲料</option></select><input type='submit'></form><div>搜索結果：<ul>{{range .}}<li>{{.Name}} - {{.Category}} - ${{.Price}} <a href='/cart/add?id={{.ID}}'>加入購物車</a></li>{{end}}</ul></div></body></html>`)
	tmpl.Execute(w, filteredProducts)
}

func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
} 