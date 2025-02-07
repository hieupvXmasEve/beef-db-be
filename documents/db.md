# Database Schema for Beef Products Display Website

## 1. Users Table (For Authentication)
```sql
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
- **Indexes**: `email` (for faster authentication)

## 2. Product Categories Table
```sql
CREATE TABLE product_categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
- **Indexes**: `name` (for quick category lookups)

## 3. Products Table
```sql
CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    category_id INT NOT NULL,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES product_categories(id) ON DELETE CASCADE
);
```
- **Indexes**: `category_id` (to speed up filtering by category)

## 4. Website Settings Table (For Site-Wide Configurations)
```sql
CREATE TABLE website_settings (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    value TEXT NOT NULL
);
```
- This table can store key-value pairs like site name, email, phone, etc.

## 5. Blog Posts Table
```sql
CREATE TABLE blog_posts (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
- **Indexes**: `title` (to optimize search performance)

## 6. Contact Messages Table (For Contact Form Submissions)
```sql
CREATE TABLE contact_messages (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
- **Indexes**: `email` (to group user messages efficiently)

## Indexes for Optimization
- `users.email` (for faster authentication)
- `product_categories.name` (for quick lookups)
- `products.category_id` (to speed up category filtering)
- `blog_posts.title` (to optimize search performance)

## Additional Features (Optional)
Would you like to include:
- **Product tags** for better filtering?
- **User roles** (admin, editor, etc.)?
- **SEO fields** for blogs and products (meta title, meta description)?

Let me know if you need any modifications! 🚀
