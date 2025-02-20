-- Drop triggers
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

DROP TRIGGER IF EXISTS update_pages_updated_at ON pages;

DROP TRIGGER IF EXISTS update_blog_posts_updated_at ON blog_posts;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at_column ();

-- Drop tables
DROP TABLE IF EXISTS contact_messages;

DROP TABLE IF EXISTS pages;

DROP TABLE IF EXISTS blog_posts;

DROP TABLE IF EXISTS website_settings;

DROP TABLE IF EXISTS products;

DROP TABLE IF EXISTS categories;

DROP TABLE IF EXISTS users;