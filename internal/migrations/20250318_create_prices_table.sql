CREATE TABLE
    prices (
        `id` varchar(255) PRIMARY KEY,
        `product_id` varchar(255) NOT NULL,
        `selling_price` decimal NOT NULL,
        `purchase_price` decimal NOT NULL,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `created_by` varchar(255),
        `updated_by` varchar(255),

        UNIQUE INDEX idx_product_id (`product_id`)
    );