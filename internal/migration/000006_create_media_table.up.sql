CREATE TABLE media (
    id SERIAL PRIMARY KEY,
    model_type VARCHAR(1024),
    model_id BIGINT,
    name VARCHAR(1024),
    file_name VARCHAR(1024),
    mime_type VARCHAR(1024), 
    default_properties JSONB,
    custom_properties JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);