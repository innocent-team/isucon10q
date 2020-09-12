DROP DATABASE IF EXISTS isuumo;
CREATE DATABASE isuumo;

DROP TABLE IF EXISTS isuumo.estate;
DROP TABLE IF EXISTS isuumo.chair;

CREATE TABLE isuumo.estate
(
    id          INTEGER             NOT NULL PRIMARY KEY,
    name        VARCHAR(64) CHARACTER SET utf8mb4 NOT NULL,
    description VARCHAR(4096) CHARACTER SET utf8mb4 NOT NULL,
    thumbnail   VARCHAR(128) CHARACTER SET utf8mb4 NOT NULL,
    address     VARCHAR(128) CHARACTER SET utf8mb4 NOT NULL,
    latitude    DOUBLE PRECISION    NOT NULL,
    longitude   DOUBLE PRECISION    NOT NULL,
    rent        INTEGER             NOT NULL,
    door_height INTEGER             NOT NULL,
    door_width  INTEGER             NOT NULL,
    features    VARCHAR(64) CHARACTER SET utf8mb4 NOT NULL,
    popularity  INTEGER             NOT NULL
);

CREATE TABLE isuumo.chair
(
    id          INTEGER         NOT NULL PRIMARY KEY,
    name        VARCHAR(64) CHARACTER SET utf8mb4 NOT NULL,
    description VARCHAR(4096) CHARACTER SET utf8mb4 NOT NULL,
    thumbnail   VARCHAR(128) CHARACTER SET utf8mb4 NOT NULL,
    price       INTEGER         NOT NULL,
    height      INTEGER         NOT NULL,
    width       INTEGER         NOT NULL,
    depth       INTEGER         NOT NULL,
    color       VARCHAR(64) CHARACTER SET utf8mb4 NOT NULL,
    features    VARCHAR(64) CHARACTER SET utf8mb4 NOT NULL,
    kind        VARCHAR(64) CHARACTER SET utf8mb4 NOT NULL,
    popularity  INTEGER         NOT NULL,
    stock       INTEGER         NOT NULL
);

CREATE INDEX chair_stock_price ON isuumo.chair (stock, price);
CREATE INDEX estate_rent ON isuumo.estate (rent);
CREATE INDEX estate_latitude_longitude ON isuumo.estate(latitude, longitude);
