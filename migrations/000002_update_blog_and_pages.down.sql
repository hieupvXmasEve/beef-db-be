-- Drop triggers
DROP TRIGGER IF EXISTS update_pages_updated_at ON pages;

DROP TRIGGER IF EXISTS update_blog_posts_updated_at ON blog_posts;

-- Drop indexes
DROP INDEX IF EXISTS idx_blog_posts_slug;

-- Revert column types
ALTER TABLE pages
ALTER COLUMN slug TYPE TEXT;

ALTER TABLE pages
ALTER COLUMN title TYPE TEXT;

-- Drop added columns
ALTER TABLE blog_posts
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE blog_posts
DROP COLUMN IF EXISTS description;

ALTER TABLE pages
DROP COLUMN IF EXISTS description;