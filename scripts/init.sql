CREATE DATABASE IF NOT EXISTS seckill;
USE seckill;

CREATE TABLE seckill_order (
                               id BIGINT PRIMARY KEY AUTO_INCREMENT,
                               order_id VARCHAR(64) UNIQUE,
                               user_id BIGINT NOT NULL,
                               activity_id BIGINT NOT NULL,
                               status TINYINT DEFAULT 1,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               UNIQUE KEY uk_user_activity (user_id, activity_id)
) ENGINE=InnoDB;