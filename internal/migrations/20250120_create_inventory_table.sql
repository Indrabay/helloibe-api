CREATE TABLE
    inventories (
        `id` varchar(255) PRIMARY KEY,
        `product_id` varchar(255) NOT NULL,
        `quantity` DECIMAL(10, 2) NOT NULL DEFAULT 0,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `created_by` varchar(255),
        `updated_by` varchar(255),

        INDEX idx_product (`product_id`),
        FOREIGN KEY (`product_id`) REFERENCES products(`id`) ON DELETE RESTRICT
    );

-- Note: inventory.quantity should be maintained as the sum of all inventory_detail.quantity 
-- for the same inventory_id. This can be done via application logic or database triggers.

