CREATE TABLE
    stocks (
        `id` varchar(255) PRIMARY KEY,
        `product_id` varchar(255) NOT NULL,
        `quantity` decimal NOT NULL,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `created_by` varchar(255),
        `updated_by` varchar(255),

        INDEX idx_product (`product_id`)
    );