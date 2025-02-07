# Beef Database Backend

## Database Migrations

This project uses `golang-migrate` for database migrations. Below is a comprehensive guide on how to work with migrations.

### Prerequisites

1. Install the migrate tool:
```bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

2. Set up your environment variables in `.env`:
```env
# Database Configuration
DB_HOST=your_host
DB_PORT=3306
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_database
```

### Migration Commands

All migration commands are available through the Makefile:

1. Create a new migration:
```bash
make migrate-create NAME=your_migration_name
```
This will create two files:
- `migrations/XXXXXX_your_migration_name.up.sql` - For applying changes
- `migrations/XXXXXX_your_migration_name.down.sql` - For reverting changes

2. Apply migrations:
```bash
make migrate-up
```
This will apply all pending migrations in sequential order.

3. Revert migrations:
```bash
make migrate-down
```
This will revert all migrations in reverse order.

4. Force a specific version:
```bash
make migrate-force VERSION=X
```
Use this to force the migration version when needed (e.g., to recover from errors).

### Current Database Schema

The database includes the following tables:

1. `users` - User management
   - Primary key: `id` (BIGINT AUTO_INCREMENT)
   - Unique fields: `username`, `email`
   - Role enum: 'admin', 'user'
   - Timestamps: `created_at`, `updated_at`
   - Index: `idx_users_email`

2. `product_categories` - Product categorization
   - Primary key: `id` (INT AUTO_INCREMENT)
   - Unique field: `name`
   - Index: `idx_product_categories_name`

3. `products` - Product information
   - Primary key: `id` (INT AUTO_INCREMENT)
   - Foreign key: `category_id` references `product_categories`
   - Index: `idx_products_category_id`
   - ON DELETE CASCADE for category relationship

4. `website_settings` - Site configuration
   - Primary key: `id` (INT AUTO_INCREMENT)
   - Key-value storage for site settings

5. `blog_posts` - Blog content
   - Primary key: `id` (INT AUTO_INCREMENT)
   - Index: `idx_blog_posts_title`

6. `contact_messages` - Contact form submissions
   - Primary key: `id` (INT AUTO_INCREMENT)
   - Index: `idx_contact_messages_email`

### Migration Files Structure

- `migrations/` directory contains all migration files
- Files are numbered sequentially (e.g., 000001, 000002)
- Each migration has an 'up' and 'down' file
- Migrations are executed in order based on their version number

### Best Practices

1. Always create both up and down migrations
2. Test migrations locally before applying to production
3. Back up your database before running migrations
4. Use transactions for complex migrations
5. Follow the naming convention: `XXXXXX_descriptive_name.sql`
6. Handle foreign key constraints in the correct order
7. Include indexes in the same migration as table creation

### Troubleshooting

1. If migrations fail:
   ```bash
   # Force to a specific version
   make migrate-force VERSION=X
   
   # Then try running migrations again
   make migrate-up
   ```

2. To check current migration status:
   ```bash
   cd scripts && ./migrate.sh version
   ```

3. Common issues:
   - Foreign key constraints: Ensure tables are created/dropped in the correct order
   - Duplicate indexes: Check if indexes already exist
   - Connection issues: Verify database credentials in `.env`

### Security Notes

- Never commit `.env` file with real credentials
- Use `.env.example` as a template
- Ensure proper access controls on production databases
- Regularly audit database access and permissions
