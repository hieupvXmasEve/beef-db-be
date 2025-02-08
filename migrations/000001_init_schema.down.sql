-- Drop tables in correct order (respecting foreign key constraints)
SET
    FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS contact_messages;

DROP TABLE IF EXISTS blog_posts;

DROP TABLE IF EXISTS website_settings;

DROP TABLE IF EXISTS products;

DROP TABLE IF EXISTS categories;

DROP TABLE IF EXISTS users;

SET
    FOREIGN_KEY_CHECKS = 1;