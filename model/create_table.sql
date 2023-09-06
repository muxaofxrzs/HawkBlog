CREATE TABLE `comment_like`(
                       `id` bigint NOT NULL AUTO_INCREMENT,
                       `user_id` bigint NOT NULL,
                       `article_id` bigint NOT NULL,
                       `comment_id` bigint NOT NULL,
                       `status` int NOT NULL DEFAULT 0,
                       PRIMARY KEY (`id`),
                       UNIQUE KEY `idx_username` (`comment_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;