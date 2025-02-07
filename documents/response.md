# REST API Response Format

## 1. Introduction
The response of a REST API typically follows a standard structure to help clients easily process data. Below are some general guidelines that an API should follow.

---

## 2. General Response Format
Typically, REST API responses use JSON format with the following important fields:

```json
{
  "status": "success",
  "message": "Data retrieved successfully",
  "data": null,
  "errors": null
}
```

### **Important Fields**
- **`status`**: The response status, can be either `"success"` or `"error"`.
- **`message`**: A message describing the API result.
- **`data`**: Contains the returned data, which can be an object or an array.
- **`errors`**: A list of errors if the request fails.

---

## 3. Error Response Format
When a request fails, the API should return the following format:

```json
{
  "status": "error",
  "message": "An error occurred",
  "errors": [
    {
      "field": "id",
      "message": "Invalid ID"
    }
  ]
}
```

---

## 4. Common Status Codes

| Status Code | Meaning |
|------------|---------|
| `200 OK` | Request successful |
| `201 Created` | Resource successfully created |
| `400 Bad Request` | Invalid request |
| `401 Unauthorized` | Authentication required |
| `403 Forbidden` | Access denied |
| `404 Not Found` | Resource not found |
| `500 Internal Server Error` | Server error |

---

## 5. Common Response and Error Structure in Go
Below is a shared struct for the entire project:

```go
package common

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, errors interface{}) Response {
	return Response{
		Status:  "error",
		Message: message,
		Errors:  errors,
	}
}
```

Example usage in an API handler:

```go
package main

import (
	"encoding/json"
	"net/http"
	"your_project/common"
)

func handler(w http.ResponseWriter, r *http.Request) {
	response := common.SuccessResponse("Data retrieved successfully", []map[string]interface{}{
		{"id": 1, "name": "Kobe Beef", "price": 5000000, "image": "https://example.com/images/kobe.jpg"},
		{"id": 2, "name": "Wagyu Beef", "price": 3000000, "image": "https://example.com/images/wagyu.jpg"},
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/products", handler)
	http.ListenAndServe(":8080", nil)
}
```

---

## 6. Conclusion
- **Use JSON** as the standard response format.
- **Return appropriate status codes** to help clients understand errors.
- **Create a common response struct** to standardize API data.
- **Avoid unnecessary data** to optimize performance.

**Hope this document helps you build a professional REST API! 🚀**

