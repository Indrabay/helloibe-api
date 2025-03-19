CREATE TABLE
    stock_expiry_dates (
        `id` varchar(255) PRIMARY KEY,
        `stock_id` varchar(255) NOT NULL,
        `quantity` decimal NOT NULL,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `created_by` varchar(255),
        `updated_by` varchar(255),

        INDEX idx_stock (`stock_id`)
    );