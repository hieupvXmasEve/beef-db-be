# API Workflow Documentation: `/api/products`

## Overview

The `/api/products` endpoint is a RESTful API designed to handle operations related to products in the system. This documentation provides a detailed explanation of the workflow, including the functions involved in processing requests and generating responses.

## Endpoint Details

- **URL:** `/api/products`
- **Methods:**
  - `GET` - Retrieve a list of products.
  - `POST` - Create a new product. *(Requires authentication)*

## Workflow Steps

### 1. Routing the Request

When a client sends a request to `/api/products`, the request is first routed through the main router defined in `cmd/api/main.go`. Depending on the HTTP method, the request is directed to the appropriate handler function.
```go
// Public product routes
r.Get("/products", productHandler.ListProducts)
r.Post("/products", productHandler.CreateProduct)
```

### 2. Handling the Request in the Handler Layer

#### a. `ListProducts` Handler

Handles `GET` requests to retrieve a list of products.

```go
// ListProducts retrieves all products
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	pagination := utils.GetPaginationFromRequest(r)

	products, totalCount, err := h.productService.ListProducts(r.Context(), pagination)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError,
			model.NewErrorResponse("Failed to retrieve products", err.Error()))
		return
	}

	paginatedResp := model.NewPaginatedResponse(products, totalCount, pagination.Page, pagination.PageSize)
	utils.SendResponse(w, http.StatusOK,
		model.NewSuccessResponse("Products retrieved successfully", paginatedResp))
}
```

**Functionality:**

1. **Extract Pagination Parameters:** Utilizes `utils.GetPaginationFromRequest` to parse pagination details (`page`, `page_size`) from the request query parameters.
2. **Service Layer Interaction:** Calls `productService.ListProducts` with the extracted pagination to fetch products and the total count.
3. **Error Handling:** If an error occurs during service interaction, responds with a `500 Internal Server Error` and an appropriate error message.
4. **Response Formation:** Constructs a paginated response using `model.NewPaginatedResponse` and sends a `200 OK` response with the product data.

#### b. `CreateProduct` Handler

Handles `POST` requests to create a new product. *(Requires authentication)*

```go
// CreateProduct handles product creation
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req model.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Invalid request body", []model.ValidationError{
				model.NewValidationError("body", "Invalid JSON format"),
			}))
		return
	}

	product, err := h.productService.CreateProduct(r.Context(), req)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest,
			model.NewErrorResponse("Failed to create product", err.Error()))
		return
	}

	utils.SendResponse(w, http.StatusCreated,
		model.NewSuccessResponse("Product created successfully", product))
}
```

**Functionality:**

1. **Decode Request Body:** Parses the JSON payload into a `CreateProductRequest` struct.
2. **Input Validation:** If decoding fails, responds with a `400 Bad Request` and validation errors.
3. **Service Layer Interaction:** Calls `productService.CreateProduct` with the validated request data to create a new product.
4. **Error Handling:** If product creation fails, responds with a `400 Bad Request` and an error message.
5. **Response Formation:** On successful creation, sends a `201 Created` response with the newly created product data.

### 3. Service Layer Processing

The service layer encapsulates the business logic and interacts with the repository layer to perform database operations.

#### a. `ListProducts` Service

```go
func (s *ProductService) ListProducts(ctx context.Context, pagination model.Pagination) ([]model.Product, int64, error) {
	// Get total count first
	totalCount, err := s.queries.GetTotalProducts(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %v", err)
	}

	params := repository.ListProductsParams{
		Limit:  int32(pagination.GetLimit()),
		Offset: int32(pagination.GetOffset()),
	}

	products, err := s.queries.ListProducts(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	// Convert repository products to model products
	result := make([]model.Product, len(products))
	for i, p := range products {
		result[i] = model.Product{
			ID:                int(p.ID),
			CategoryID:        int(p.CategoryID),
			Name:              p.Name,
			Slug:              p.Slug,
			Description:       p.Description,
			Price:             p.Price,
			PriceSale:         p.PriceSale,
			ImageURL:          p.ImageUrl,
			ThumbURL:          p.ThumbUrl,
			CreatedAt:         p.CreatedAt.Time,
			CategoryName:      p.CategoryName,
			CategorySlug:      p.CategorySlug,
			UnitOfMeasurement: p.UnitOfMeasurement,
		}
	}

	return result, totalCount, nil
}
```

**Functionality:**

1. **Retrieve Total Count:** Calls `queries.GetTotalProducts` to determine the total number of products in the database.
2. **Fetch Products:** Uses `queries.ListProducts` with pagination parameters (`limit`, `offset`) to retrieve the desired subset of products.
3. **Data Transformation:** Converts the repository layer's `Product` models to the service layer's `model.Product` structures.
4. **Return Data:** Provides the list of products and the total count to the handler.

#### b. `CreateProduct` Service

```go
func (s *ProductService) CreateProduct(ctx context.Context, req model.CreateProductRequest) (*model.Product, error) {
	result, err := s.queries.CreateProduct(ctx, repository.CreateProductParams{
		CategoryID:        int32(req.CategoryID),
		Name:              req.Name,
		Slug:              req.Slug,
		Description:       req.Description,
		Price:             req.Price,
		PriceSale:         req.PriceSale,
		ImageUrl:          req.ImageURL,
		UnitOfMeasurement: req.UnitOfMeasurement,
		ThumbUrl:          req.ThumbURL,
	})
	if err != nil {
		return nil, err
	}

	return s.GetProduct(ctx, int(result))
}
```

**Functionality:**

1. **Database Insertion:** Calls `queries.CreateProduct` with the provided product details to insert a new record into the database.
2. **Retrieve Created Product:** After insertion, fetches the newly created product using its `ID` by calling `s.GetProduct`.
3. **Return Data:** Provides the complete product details back to the handler.

### 4. Repository Layer Interaction

The repository layer handles direct interactions with the database using SQL queries generated by `sqlc`.

#### a. `ListProducts` Query

```sql:sqlc/query.sql
-- name: ListProducts :many
SELECT
    p.id,
    p.category_id,
    p.name,
    p.slug,
    p.description,
    p.price,
    p.price_sale,
    p.unit_of_measurement,
    p.image_url,
    p.thumb_url,
    p.created_at,
    c.name as category_name,
    c.slug as category_slug
FROM products p
JOIN categories c ON p.category_id = c.id
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2;
```

**Functionality:**

- **Purpose:** Retrieves a list of products with associated category details.
- **Parameters:**
  - `$1` - `LIMIT` for pagination.
  - `$2` - `OFFSET` for pagination.
- **Joins:** Combines `products` with `categories` to include category names and slugs.

#### b. `CreateProduct` Query

```sql:sqlc/query.sql
-- name: CreateProduct :one
INSERT INTO products (
    category_id,
    name,
    slug,
    description,
    price,
    price_sale,
    unit_of_measurement,
    image_url,
    thumb_url
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;
```

**Functionality:**

- **Purpose:** Inserts a new product into the `products` table.
- **Parameters:**
  - `$1` - `category_id`
  - `$2` - `name`
  - `$3` - `slug`
  - `$4` - `description`
  - `$5` - `price`
  - `$6` - `price_sale`
  - `$7` - `unit_of_measurement`
  - `$8` - `image_url`
  - `$9` - `thumb_url`
- **Returns:** The `id` of the newly created product.

### 5. Response Formation

Responses adhere to the standardized format defined in `documents/response.md`. The `common.Response` struct includes fields like `status`, `message`, `data`, and `errors`.

```go
```json
{
  "status": "success",
  "message": "Data retrieved successfully",
  "data": null,
  "errors": null
}
```
```

**Response Structure:**

- **Success Response:**
  ```json
  {
    "status": "success",
    "message": "Products retrieved successfully",
    "data": { /* Product Data */ },
    "errors": null
  }
  ```

- **Error Response:**
  ```json
  {
    "status": "error",
    "message": "Failed to retrieve products",
    "data": null,
    "errors": [ /* Error Details */ ]
  }
  ```

### 6. Error Handling

Throughout the workflow, errors are handled gracefully with appropriate HTTP status codes and descriptive messages.

- **400 Bad Request:** Invalid input data or request parameters.
- **404 Not Found:** Requested resource does not exist.
- **500 Internal Server Error:** Unexpected server-side issues.

Errors are encapsulated using the `model.ErrorResponse` struct and sent back to the client following the standardized response format.

### 7. Authentication & Authorization *(For Protected Routes)*

Endpoints like `POST /api/products` require authenticated access. Authentication is managed using JWT tokens handled in `internal/utils/jwt.go`.

- **JWT Generation:** Tokens are generated upon successful user authentication and set as HTTP-only cookies.
- **Middleware:** Protected routes utilize middleware to verify JWT tokens and authorize requests.

```go
// GenerateJWT creates a new JWT token for a user
func GenerateJWT(userID int64) (string, error) {
	// Implementation...
}
```

### 8. Database Schema

The `sqlc`-generated queries interact with the PostgreSQL database. Key tables involved in the `/api/products` workflow include:

- **`products`**
- **`categories`**
- **`users`** *(For authentication)*

Ensure that the database schema is up-to-date and that all necessary indexes are in place for optimal performance.

```sql:sqlc/schema.sql
-- Products Table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL,
    name VARCHAR(150) NOT NULL,
    slug VARCHAR(200) NOT NULL UNIQUE,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    price_sale DECIMAL(10, 2) NULL,
    unit_of_measurement VARCHAR(50) NOT NULL DEFAULT 'piece',
    image_url VARCHAR(255),
    thumb_url VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);
```

## Summary

The `/api/products` endpoint follows a layered architecture comprising the handler, service, and repository layers to ensure separation of concerns, maintainability, and scalability. Proper error handling, standardized responses, and adherence to RESTful principles make the API robust and developer-friendly.

For further enhancements, consider implementing features like search, filtering, sorting, and advanced pagination to provide more flexibility to API consumers.

---
