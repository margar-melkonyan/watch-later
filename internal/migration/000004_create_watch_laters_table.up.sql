CREATE TABLE watch_laters (
    id SERIAL PRIMARY KEY,
    name VARCHAR(4096),
    text TEXT,
    user_id BIGINT,
    category_id BIGINT DEFAULT NULL,
    platform_id BIGINT DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT fk_watch_laters_users_id
        FOREIGN KEY (user_id)
            REFERENCES users (id),
    CONSTRAINT fk_watch_laters_categories_id
        FOREIGN KEY (category_id)
            REFERENCES categories (id),
    CONSTRAINT fk_watch_laters_platofrms_id
        FOREIGN KEY (platform_id)
            REFERENCES platforms (id)
);