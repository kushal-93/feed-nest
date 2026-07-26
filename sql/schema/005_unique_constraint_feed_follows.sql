-- +goose Up
ALTER TABLE feed_follows
ADD CONSTRAINT feed_follows_feed_id_user_id_key
UNIQUE (feed_id, user_id);

-- +goose Down
ALTER TABLE feed_follows
DROP CONSTRAINT feed_follows_feed_id_user_id_key;