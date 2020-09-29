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
    popularity  INTEGER             NOT NULL,
    popularity_desc INTEGER AS (-popularity) NOT NULL
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
    popularity_desc INTEGER AS (-popularity) NOT NULL,
    stock       INTEGER         NOT NULL
);


CREATE INDEX chair_pricet_stock ON isuumo.chair (price, stock);
CREATE INDEX chair_height_stock ON isuumo.chair (height, stock);
CREATE INDEX chair_width_stock ON isuumo.chair (width, stock);
CREATE INDEX chair_depth_stock ON isuumo.chair (depth, stock);
CREATE INDEX chair_kind_stock ON isuumo.chair (kind, stock);
CREATE INDEX chair_color_stock ON isuumo.chair (color, stock);
CREATE INDEX chair_features_stock ON isuumo.chair (features, stock);
CREATE INDEX chair_price_id_stock ON isuumo.chair (price, id, stock);

CREATE INDEX estate_door_height ON isuumo.estate (door_height);
CREATE INDEX estate_door_width ON isuumo.estate (door_width);
CREATE INDEX estate_rent ON isuumo.estate (rent);
CREATE INDEX estate_features ON isuumo.estate (features);
CREATE INDEX estate_rent_id ON isuumo.estate (rent, id);

CREATE INDEX estate_latitude_longitude ON isuumo.estate(latitude, longitude);
CREATE INDEX estate_door_height_door_width ON isuumo.estate (door_height, door_width);
CREATE INDEX estate_door_height_rent ON isuumo.estate (door_height, rent);
CREATE INDEX estate_door_width_rent ON isuumo.estate (door_width, rent);
CREATE INDEX estate_door_width_door_widht_rent ON isuumo.estate (door_height, door_width, rent);
