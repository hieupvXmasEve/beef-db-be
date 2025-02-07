-- Add username field back to users table
ALTER TABLE users ADD COLUMN username VARCHAR(255) UNIQUE AFTER id; 