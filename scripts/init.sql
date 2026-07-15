CREATE TABLE IF NOT EXISTS seckill_order (
                                             id BIGINT PRIMARY KEY AUTO_INCREMENT,
                                             user_id BIGINT NOT NULL,
                                             activity_id BIGINT NOT NULL,
                                             status TINYINT DEFAULT 1,
                                             created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                             UNIQUE KEY uk_user_activity (user_id, activity_id)
    ) ENGINE=InnoDB;