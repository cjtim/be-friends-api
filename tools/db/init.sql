CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "name" varchar(255) NOT NULL,
    "email" varchar(255),
    "password" varchar(255),
    "line_uid" varchar(255),
    "picture_url" text,
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
    "name" varchar(255) NOT NULL,
    "description" varchar(255),
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "PK_tag_id" PRIMARY KEY ("id"),
    CONSTRAINT "CS_tag_name" UNIQUE ("name")
);
-- Search by id and name
CREATE INDEX IDX01_id_tag ON "tag" (id);
CREATE INDEX IDX02_name_tag ON "tag" (name);

------------------------------
-- tag_user Table
-- 1 to 1 tag and user
------------------------------
CREATE TABLE "tag_user" (
    user_id uuid NOT NULL,
    tag_id SERIAL NOT NULL,

    CONSTRAINT "PK_user_tag_id" PRIMARY KEY (user_id, tag_id),
    CONSTRAINT "FK_user" FOREIGN KEY (user_id) REFERENCES "user" ("id"),
    CONSTRAINT "FK_tag" FOREIGN KEY (tag_id) REFERENCES "tag" ("id")
);

------------------------------
-- pet Table
-- 1 to 1 tag and user
------------------------------
CREATE TABLE "pet" (
    "id" SERIAL NOT NULL,
    "name" varchar(255) NOT NULL,
    "description" varchar(255),
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "PK_pet_id" PRIMARY KEY ("id")
);
-- Search by id and name
CREATE INDEX IDX01_id_pet ON "pet" (id);
CREATE INDEX IDX02_name_pet ON "pet" (name);

------------------------------
-- tag_user Table
-- 1 to 1 tag and user
------------------------------
CREATE TABLE "tag_pet" (
    pet_id SERIAL NOT NULL,
    tag_id SERIAL NOT NULL,

    CONSTRAINT "PK_pet_tag_id" PRIMARY KEY (pet_id, tag_id),
    CONSTRAINT "FK_pet" FOREIGN KEY (pet_id) REFERENCES "pet" ("id"),
    CONSTRAINT "FK_tag" FOREIGN KEY (tag_id) REFERENCES "tag" ("id")
);
