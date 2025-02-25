# Request AI AGENT to Use `golang-migrate` to Create Migration Files Based on Schema  

## Description
You are an AI specialized in Golang and database migrations. I am using `golang-migrate` to manage database migrations in my project.  

## Requirements
Please perform the following steps:  

### 1. Create a Migration File
Use `golang-migrate` to create a new migration file based on the following schema:

```sql
-- Users Table
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    role ENUM('admin', 'user') NOT NULL DEFAULT 'user',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create index on email for faster authentication
CREATE INDEX idx_users_email ON users(email);

-- Product Categories Table
CREATE TABLE product_categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index on category name for quick lookups
CREATE INDEX idx_product_categories_name ON product_categories(name);

-- Products Table
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

-- Create index on category_id for faster filtering
CREATE INDEX idx_products_category_id ON products(category_id);

-- Website Settings Table
CREATE TABLE website_settings (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    value TEXT NOT NULL
);

-- Blog Posts Table
CREATE TABLE blog_posts (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index on blog post title for search optimization
CREATE INDEX idx_blog_posts_title ON blog_posts(title);

-- Contact Messages Table
CREATE TABLE contact_messages (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index on contact messages email for grouping messages
CREATE INDEX idx_contact_messages_email ON contact_messages(email); 
```

### 2. Generate CLI Command

Write the golang-migrate command to generate a migration file named create_users_table with the correct syntax.

### 3. Write Migration Content

Provide the content of the migration file (up and down).

### 4. Guide to Running Migration

Explain how to run the migration using golang-migrate with MySQL.
Ensure that the commands and code are executable in a real-world environment.