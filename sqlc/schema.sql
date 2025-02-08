-- Users Table
CREATE TABLE
    users (
        id BIGINT PRIMARY KEY AUTO_INCREMENT,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        role ENUM ('admin', 'user') NOT NULL DEFAULT 'user',
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );

-- Create index on email for faster authentication
CREATE INDEX idx_users_email ON users (email);

-- Product Categories Table
CREATE TABLE
    categories (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(100) NOT NULL,
        slug VARCHAR(150) NOT NULL UNIQUE,
        description TEXT,
        image_url VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- Create index on category name and slug for quick lookups
CREATE INDEX idx_categories_name ON categories (name);

CREATE INDEX idx_categories_slug ON categories (slug);

-- Products Table
CREATE TABLE
    products (
        id INT PRIMARY KEY AUTO_INCREMENT,
        category_id INT NOT NULL,
        name VARCHAR(150) NOT NULL,
        slug VARCHAR(200) NOT NULL UNIQUE,
        description TEXT,
        price DECIMAL(10, 2) NOT NULL,
        image_url VARCHAR(255),
        thumb_url VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (category_id) REFERENCES product_categories (id) ON DELETE CASCADE
    );

-- Create index on category_id for faster filtering
CREATE INDEX idx_products_category_id ON products (category_id);

CREATE INDEX idx_products_slug ON products (slug);

-- Website Settings Table
CREATE TABLE
    website_settings (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        value TEXT NOT NULL
    );

-- Blog Posts Table
CREATE TABLE
    blog_posts (
        id INT PRIMARY KEY AUTO_INCREMENT,
        title VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        image_url VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- Create index on blog post title for search optimization
CREATE INDEX idx_blog_posts_title ON blog_posts (title);

-- Contact Messages Table
CREATE TABLE
    contact_messages (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(150) NOT NULL,
        message TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- Create index on contact messages email for grouping messages
CREATE INDEX idx_contact_messages_email ON contact_messages (email);