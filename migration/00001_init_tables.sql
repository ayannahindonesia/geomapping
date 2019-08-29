-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "provinsi" (
    "id" bigserial,
    "name" text,
    "created_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "kota" (
    "id" bigserial,
    "provinsi_id" bigserial,
    "name" text,
    "type" text DEFAULT ('kota'),
    "created_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("provinsi_id") REFERENCES provinsi(id),
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "kecamatan" (
    "id" bigserial,
    "kota_id" bigserial,
    "name" text,
    "created_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("kota_id") REFERENCES kota(id),
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "kelurahan" (
    "id" bigserial,
    "kecamatan_id" bigserial,
    "name" text,
    "zip_code" int,
    "area_code" varchar(255),
    "created_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_time" timestamptz DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("kecamatan_id") REFERENCES kecamatan(id),
    PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS "kelurahan" CASCADE;
DROP TABLE IF EXISTS "kecamatan" CASCADE;
DROP TABLE IF EXISTS "kota" CASCADE;
DROP TABLE IF EXISTS "provinsi" CASCADE;