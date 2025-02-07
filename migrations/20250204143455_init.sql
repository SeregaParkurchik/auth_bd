-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    user_type VARCHAR(10) CHECK (user_type IN ('client', 'moderator')) NOT NULL,
    token VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE houses (
    id SERIAL PRIMARY KEY,                     -- Уникальный номер дома
    address VARCHAR(255) NOT NULL,            -- Адрес дома
    year INT NOT NULL,                         -- Год постройки
    developer VARCHAR(255),                    -- Застройщик (может быть NULL)
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания дома
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()  -- Дата последнего обновления
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users';
-- +goose StatementEnd
