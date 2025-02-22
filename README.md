# Mock AWS SES API (Go + Gin)

This is a **mock implementation of AWS SES** built using **Go** and **Gin**, designed for testing email sending **without actually sending emails**. It mimics AWS SES behavior and tracks email statistics.

## 📌n Features
👉 **Mock email sending** (follows AWS SES API structure)  
👉 **Tracks statistics** (total emails sent, per user, last sent times)  
👉 **Implements AWS SES "warming-up" rule** (daily limits)  
👉 **Thread-safe** with `sync.Mutex`  

---

## Special Rules Implemented
### Email Warming Up Rule
AWS SES imposes restrictions on new accounts, limiting the number of emails that can be sent per day.

- **Currently Hardcoded:** 10 emails per day.
- **Behavior:** If the limit is reached, an error is returned.
- **Possible Enhancement:** Make the limit configurable.

### Error Codes and Meanings
| HTTP Status | Error Code       | Description                               |
|-------------|-----------------|-------------------------------------------|
| 200         | Success          | Email request accepted (simulated).      |
| 400         | LimitExceeded    | Exceeded daily email limit.              |
| 400         | InvalidEmail     | Invalid email format or missing fields.  |
| 500         | InternalError    | Unexpected server error.                 |

---

## 🚀 Current Implementation Approach
### 1. Gin Web Framework
Used to define HTTP routes and handle requests efficiently.

### 2. Hardcoded Rate Limits
- The API currently enforces a **daily limit of 10 emails per sender**.
- If exceeded, it returns `400 Bad Request - LimitExceeded`.

### 3. Data Storage (In-Memory)
- Email logs and statistics are stored in **Go maps** (not persistent).
- This makes it **fast** but resets when the server restarts.

---

## 🚀 Setup & Run

### **1️⃣ Clone the repository**  
```sh
git clone https://github.com/yourusername/mock-aws-ses.git
cd mock-aws-ses
```

### **2️⃣ Install dependencies**  
```sh
go mod tidy
```

### **3️⃣ Run the server**  
```sh
go run main.go
```

### **4️⃣ Test with curl or Postman**  
#### **Send an email (mock)**  
```sh
curl -X POST "http://localhost:8080/send-email" -H "Content-Type: application/json" -d '{
  "sender": "user@example.com",
  "recipient": "receiver@example.com",
  "subject": "Test",
  "body": "Hello, this is a test email!"
}'
```

#### **Check email stats**  
```sh
curl -X GET "http://localhost:8080/stats"
```

---

## 📊 API Endpoints

### **1️⃣ Send Email**
- **POST** `/send-email`
- **Request Body**
```json
{
  "sender": "user@example.com",
  "recipient": "receiver@example.com",
  "subject": "Hello",
  "body": "Test email"
}
```
- **Response**
```json
{
  "message": "Email accepted for delivery (mock)",
  "sender": "user@example.com"
}
```

### **2️⃣ Get Email Stats**
- **GET** `/stats`
- **Response**
```json
{
  "total_sent": 15,
  "emails_per_user": {
    "user1@example.com": 5,
    "user2@example.com": 10
  },
  "daily_limit": 10,
  "last_sent_times": {
    "user1@example.com": "2025-02-20T14:35:00Z",
    "user2@example.com": "2025-02-20T14:40:00Z"
  }
}
```
