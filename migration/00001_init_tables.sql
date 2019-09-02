-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "provinces" (
    "id" bigserial,
    "name" text,
    "created_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "cities" (
    "id" bigserial,
    "province_id" bigserial,
    "name" text,
    "type" text DEFAULT ('kota'),
    "created_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("province_id") REFERENCES provinces(id),
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "districts" (
    "id" bigserial,
    "city_id" bigserial,
    "name" text,
    "created_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("city_id") REFERENCES cities(id),
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "villages" (
    "id" bigserial,
    "district_id" bigserial,
    "name" text,
    "zip_code" int,
    "area_code" varchar(255),
    "created_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("district_id") REFERENCES districts(id),
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS "villages" CASCADE;
DROP TABLE IF EXISTS "districts" CASCADE;
DROP TABLE IF EXISTS "cities" CASCADE;
DROP TABLE IF EXISTS "provinces" CASCADE;