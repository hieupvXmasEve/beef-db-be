ALTER TABLE products
ADD COLUMN unit_of_measurement VARCHAR(50) NOT NULL DEFAULT 'piece';

-- Update existing products to have the default value
UPDATE products SET unit_of_measurement = 'piece' WHERE unit_of_measurement IS NULL;
