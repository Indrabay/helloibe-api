CREATE TABLE
    products (
        `id` varchar(255) PRIMARY KEY,
        `name` varchar(255) NOT NULL,
        `sku` varchar(255) NOT NULL,
        `barcode` varchar(255) NOT NULL,
        `store_id` int,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `created_by` varchar(255),
        `updated_by` varchar(255),

        INDEX idx_sku (`sku`),
        INDEX idx_store_sku (`store_id`, `sku`),
        INDEX idx_store_barcode (`store_id`, `barcode`)
    );