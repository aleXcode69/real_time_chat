-- name: CanUserSubscribeSqlc :one
SELECT
  c.can_subscribe_anyone OR EXISTS (
    SELECT 1 FROM user_channel uc
    WHERE uc.user_id = $1 AND uc.channel_id = c.id
  ) AS can_subscribe
FROM channel c
WHERE c.name = $2;

-- name: CanUserPublishSqlc :one
SELECT
  c.can_publish_anyone OR EXISTS (
    SELECT 1 FROM user_channel uc
    WHERE uc.user_id = $1 AND uc.channel_id = c.id AND uc.can_publish = true
  ) AS can_publish
FROM channel c
WHERE c.name = $2;

-- name: EnsureUserSqlc :exec
INSERT INTO "user" (id)
VALUES ($1)
ON CONFLICT (id) DO NOTHING;

-- name: EnsureChannelSqlc :exec
INSERT INTO channel (name)
VALUES ($1)
ON CONFLICT (name) DO NOTHING;