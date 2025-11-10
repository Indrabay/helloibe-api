CREATE TABLE
    inventory_details (
        `id` varchar(255) PRIMARY KEY,
        `inventory_id` varchar(255) NOT NULL,
        `quantity` DECIMAL(10, 2) NOT NULL DEFAULT 0,
        `expiry_date` DATE,
        `location` varchar(255),
        `status` varchar(50) NOT NULL DEFAULT 'active',
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `created_by` varchar(255),
        `updated_by` varchar(255),

        INDEX idx_inventory (`inventory_id`),
        INDEX idx_expiry_date (`expiry_date`),
        INDEX idx_status (`status`),
        FOREIGN KEY (`inventory_id`) REFERENCES inventories(`id`) ON DELETE CASCADE
    );

