CREATE TABLE "stations" (
    "station_id"    BIGSERIAL PRIMARY KEY,
    "name"          VARCHAR(40) NOT NULL,
    "lat"           DECIMAL NOT NULL,
    "lng"           DECIMAL NOT NULL,
    "provider"      VARCHAR(40) NOT NULL
);