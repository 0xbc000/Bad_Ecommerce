# Insecure E- Website

## 簡介
這是一個使用Go語言寫的MVC電商網站，其中包含了常見的Web Application漏洞。
該網站包含基本的電商功能，如用戶管理、產品搜索、購物車、購買和訂單管理。

## 安裝方式
1. 確保已安裝Go語言環境。
2. clone repo：
   ```bash
   git clone https://github.com/0xbc000/Bad_Ecommerce
   ```
3. 進入項目目錄：
   ```bash
   cd bad_ecommerce
   ```
4. 初始化Go module並安裝dependency：
   ```bash
   go mod tidy
   ```
5. 安裝localhost憑證
   ```bash
   openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj "/CN=localhost"
   ```
6. 啟動server：
   ```bash
   go run main.go
   ```
7. Check `https://localhost:443`。

## 漏洞點與已知漏洞

- **XSS**：產品搜索功能和訂單詳情中存在跨站腳本攻擊漏洞。
- **IDOR**：訂單管理中存在不安全的直接對象引用漏洞。
- **序列化的訂單號和用戶號**：訂單和用戶ID為可預測的序列號。
- **暴露的管理面板**：管理面板路徑暴露
- **弱密碼**: Admin page使用弱密碼。
- **不正確的緩存設計**：所有用戶共用一個global的購物車緩存，導致數據混淆。
- **錯誤的 session 設計**：錯誤implement session token value
- **敏感資料洩漏** 


## 修復建議

- **XSS**：對用戶輸入進行適當的編碼和過濾。
- **IDOR**：在訪問敏感資源時進行權限檢查。
- **序列化的訂單號和用戶號**：使用隨機生成的ID來替代序列號。
- **暴露的管理面板和弱密碼**：限制管理面板的訪問，並使用強密碼策略。
- **不正確的緩存設計**：為每個用戶單獨維護購物車數據。
- **錯誤的 session 設計**：使用安全的 session ID 生成策略，避免使用可預測的用戶信息。 
- **敏感資料洩漏**: 使用Mask遮掉敏感資訊
