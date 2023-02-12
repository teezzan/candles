CREATE TABLE `process_status` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `file_name` varchar(50) NOT NULL,
    `status` varchar(50) NOT NULL,
    `error` varchar(100) NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;