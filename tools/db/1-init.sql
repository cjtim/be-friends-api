CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255),
    "password" VARCHAR(255),
    "line_uid" VARCHAR(255),
    "picture_url" text,
    "is_org"  BOOL,
    "is_admin" BOOL DEFAULT FALSE NOT NULL, 
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "PK_user_id" PRIMARY KEY ("id"),
    CONSTRAINT "CS_user_email" UNIQUE ("email"),
    CONSTRAINT "CS_user_line_uid" UNIQUE ("line_uid")
);
-- Index for search by id
CREATE INDEX IDX01_id_user ON "user" (id);

------------------------------
-- Tag table
------------------------------
CREATE TABLE "tag" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255),
    "is_internal"  BOOL DEFAULT FALSE NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "PK_tag_id" PRIMARY KEY ("id"),
    CONSTRAINT "CS_tag_name" UNIQUE ("name")
);
-- Search by id and name
CREATE INDEX IDX01_id_tag ON "tag" (id);
CREATE INDEX IDX02_name_tag ON "tag" (name);

------------------------------
-- pet Table
-- 1 to 1 tag and user
------------------------------
CREATE TABLE "pet" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255),
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),
    lat NUMERIC(15,10) NOT NULL,
    lng NUMERIC(15,10) NOT NULL,

    CONSTRAINT "PK_pet_id" PRIMARY KEY ("id")
);
-- Search by id and name
CREATE INDEX IDX01_id_pet ON "pet" (id);
CREATE INDEX IDX02_name_pet ON "pet" (name);

------------------------------
-- tag_pet Table
-- 1 to 1 tag and user
------------------------------
CREATE TABLE "tag_pet" (
    pet_id SERIAL NOT NULL,
    tag_id SERIAL NOT NULL,

    CONSTRAINT "PK_pet_tag_id" PRIMARY KEY (pet_id, tag_id),
    CONSTRAINT "FK_pet" FOREIGN KEY (pet_id) REFERENCES "pet" ("id"),
    CONSTRAINT "FK_tag" FOREIGN KEY (tag_id) REFERENCES "tag" ("id")
);

------------------------------
-- pic_pet Table
-- 1 to 1 tag and url
------------------------------
CREATE TABLE "pic_pet" (
    pet_id SERIAL NOT NULL,
    picture_url TEXT NOT NULL,

    CONSTRAINT "PK_pet_id_picture_url" PRIMARY KEY (pet_id, picture_url),
    CONSTRAINT "FK_pet" FOREIGN KEY (pet_id) REFERENCES "pet" ("id")
);
