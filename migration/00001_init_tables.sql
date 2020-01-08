-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "provinces" (
    "id" bigserial,
    "name" text,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "cities" (
    "id" bigserial,
    "province_id" bigserial,
    "name" text,
    "type" text DEFAULT ('kota'),
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    FOREIGN KEY ("province_id") REFERENCES provinces(id),
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "districts" (
    "id" bigserial,
    "city_id" bigserial,
    "name" text,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    FOREIGN KEY ("city_id") REFERENCES cities(id),
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "villages" (
    "id" bigserial,
    "district_id" bigserial,
    "name" text,
    "zip_code" int,
    "area_code" varchar(255),
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    FOREIGN KEY ("district_id") REFERENCES districts(id),
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "client_configs" (
    "id" bigserial,
    "name" varchar(255) NOT NULL,
    "key" varchar(255) NOT NULL,
    "secret" varchar(255) NOT NULL,
    "role" varchar(255) NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS "villages" CASCADE;
DROP TABLE IF EXISTS "districts" CASCADE;
DROP TABLE IF EXISTS "cities" CASCADE;
DROP TABLE IF EXISTS "provinces" CASCADE;
DROP TABLE IF EXISTS "client_configs" CASCADE;