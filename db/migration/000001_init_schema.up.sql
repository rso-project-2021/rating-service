CREATE TABLE "ratings" (
    "rating_id"     BIGSERIAL PRIMARY KEY,
    "station_id"    INT NOT NULL,
    "user_id"       INT NOT NULL,
    "rating"        INT NOT NULL,
    "comment"       VARCHAR(256),
    "created_at"    TIMESTAMP NOT NULL DEFAULT(now())
);