CREATE TABLE `ohlc_data` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `time` timestamp NOT NULL,
    `symbol` varchar(50) NOT NULL,
    `open` DECIMAL(20,10) NOT NULL,
    `high` DECIMAL(20,10) NOT NULL,
    `low` DECIMAL(20,10) NOT NULL,
    `close` DECIMAL(20,10) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;