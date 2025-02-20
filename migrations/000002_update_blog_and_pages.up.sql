-- Drop triggers if they exist (this is safe even if tables don't exist)
DROP TRIGGER IF EXISTS update_pages_updated_at ON pages;
DROP TRIGGER IF EXISTS update_blog_posts_updated_at ON blog_posts;

DO $$ 
BEGIN 
    -- Create blog_posts table if it doesn't exist
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'blog_posts') THEN
        CREATE TABLE blog_posts (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            slug VARCHAR(255) NOT NULL UNIQUE,
            content TEXT NOT NULL,
            image_url VARCHAR(255),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            description TEXT
        );
        
        -- Create indexes for blog posts
        CREATE INDEX idx_blog_posts_title ON blog_posts (title);
        CREATE INDEX idx_blog_posts_slug ON blog_posts (slug);
    END IF;

    -- Create pages table if it doesn't exist
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'pages') THEN
        CREATE TABLE pages (
            id SERIAL PRIMARY KEY,
            slug VARCHAR(255) NOT NULL UNIQUE,
            title VARCHAR(255) NOT NULL,
            content TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            description TEXT
        );
        
        -- Create indexes for pages
        CREATE INDEX idx_pages_title ON pages (title);
        CREATE INDEX idx_pages_slug ON pages (slug);
    END IF;

    -- Create or replace the update_updated_at_column function
    IF NOT EXISTS (SELECT 1 FROM pg_proc WHERE proname = 'update_updated_at_column') THEN
        CREATE FUNCTION update_updated_at_column()
        RETURNS TRIGGER AS $$
        BEGIN
            NEW.updated_at = CURRENT_TIMESTAMP;
            RETURN NEW;
        END;
        $$ language 'plpgsql';
    END IF;

    -- Create triggers for both tables
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'pages') THEN
        IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'update_pages_updated_at') THEN
            CREATE TRIGGER update_pages_updated_at
                BEFORE UPDATE ON pages
                FOR EACH ROW
                EXECUTE FUNCTION update_updated_at_column();
        END IF;
    END IF;

    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'blog_posts') THEN
        IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'update_blog_posts_updated_at') THEN
            CREATE TRIGGER update_blog_posts_updated_at
                BEFORE UPDATE ON blog_posts
                FOR EACH ROW
                EXECUTE FUNCTION update_updated_at_column();
        END IF;
    END IF;

    -- Update column types if tables exist
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'pages') THEN
        ALTER TABLE pages ALTER COLUMN slug TYPE VARCHAR(255);
        ALTER TABLE pages ALTER COLUMN title TYPE VARCHAR(255);
    END IF;

    -- Add description fields if they don't exist
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'blog_posts') 
       AND NOT EXISTS (SELECT 1 FROM information_schema.columns 
                      WHERE table_name = 'blog_posts' AND column_name = 'description') THEN
        ALTER TABLE blog_posts ADD COLUMN description TEXT;
    END IF;

    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'pages')
       AND NOT EXISTS (SELECT 1 FROM information_schema.columns 
                      WHERE table_name = 'pages' AND column_name = 'description') THEN
        ALTER TABLE pages ADD COLUMN description TEXT;
    END IF;
END $$; 