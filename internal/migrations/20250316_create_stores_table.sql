CREATE TABLE
    stores (
        `id` int AUTO_INCREMENT PRIMARY KEY,
        `name` varchar(255) NOT NULL,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );