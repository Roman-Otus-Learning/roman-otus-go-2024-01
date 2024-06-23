-- +goose Up
CREATE TABLE IF NOT EXISTS event (
    id BIGSERIAL CONSTRAINT event_pk PRIMARY KEY,
    owner_id BIGINT NOT NULL,
    title VARCHAR (255) NOT NULL,
    description TEXT NOT NULL,
    time_start TIMESTAMP NOT NULL,
    time_end TIMESTAMP NOT NULL,
    notification_time TIMESTAMP NULL DEFAULT NULL
);

-- +goose Down
DROP TABLE IF EXISTS event