package models

// Product 產品結構
type Product struct {
	ID       int
	Name     string
	Price    float64
	Category string
}

// 模擬產品數據庫
var Products = []Product{
	{ID: 1, Name: "蘋果", Price: 1.99, Category: "水果"},
	{ID: 2, Name: "香蕉", Price: 0.99, Category: "水果"},
	{ID: 3, Name: "橙子", Price: 1.49, Category: "水果"},
	{ID: 4, Name: "葡萄", Price: 2.99, Category: "水果"},
	{ID: 5, Name: "西瓜", Price: 3.99, Category: "水果"},
	{ID: 6, Name: "鳳梨", Price: 2.49, Category: "水果"},
	{ID: 7, Name: "草莓", Price: 2.99, Category: "水果"},
	{ID: 8, Name: "藍莓", Price: 3.49, Category: "水果"},
	{ID: 9, Name: "芒果", Price: 1.99, Category: "水果"},
	{ID: 10, Name: "檸檬", Price: 0.89, Category: "水果"},
	{ID: 11, Name: "可樂", Price: 1.29, Category: "飲料"},
	
} 