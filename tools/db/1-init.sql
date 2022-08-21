CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

------------------------------
-- Shelter
------------------------------
CREATE TABLE "shelter" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255),
    "password" VARCHAR(255),    
    "picture_url" text,
    "description" text,
    lat NUMERIC(15,10) NOT NULL,
    lng NUMERIC(15,10) NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "PK_shelter" PRIMARY KEY ("id"),
    CONSTRAINT "CS_shelter__email" UNIQUE ("email")
);
-- Index for search by id
CREATE INDEX IDX01_id_shelter ON "shelter" (id);
CREATE INDEX IDX02_name_shelter ON "shelter" (name);

------------------------------
-- User
------------------------------
CREATE TABLE "user" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255),
    "line_uid" VARCHAR(255),
    "shelter_id" uuid,
    "picture_url" text,
    "is_admin" BOOL DEFAULT FALSE NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "PK_user" PRIMARY KEY ("id"),
    CONSTRAINT "CS_user__email" UNIQUE ("email"),
    CONSTRAINT "CS_user__line_uid" UNIQUE ("line_uid"),
    CONSTRAINT "FK_user__shelter" FOREIGN KEY (shelter_id) REFERENCES "shelter" ("id"),
    CONSTRAINT "CS_user__shelter_unique" UNIQUE ("shelter_id")
);
-- Index for search by id
CREATE INDEX IDX01_id_user ON "user" (id);
CREATE INDEX IDX02_line_uid_user ON "user" (line_uid);
CREATE INDEX IDX03_shelter_id_user ON "user" (shelter_id);

------------------------------
-- Campaign
------------------------------
CREATE TABLE "campaign" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "picture_url" text,
    "description" text,
    "min_amount" double precision NOT NULL,
    "shelter_id" uuid NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "PK_campaign" PRIMARY KEY ("id"),
    CONSTRAINT "CS_campaign__name" UNIQUE ("name"),
    CONSTRAINT "FK_campaign__shelter" FOREIGN KEY (shelter_id) REFERENCES "shelter" ("id"),
    CONSTRAINT "CS_campaign__shelter_unique" UNIQUE ("shelter_id")
);
-- Index for search by id
CREATE INDEX IDX01_id_campaign ON "campaign" (id);
CREATE INDEX IDX02_name_campaign ON "campaign" (name);
CREATE INDEX IDX03_shelter_id_campaign ON "campaign" (shelter_id);

------------------------------
-- donate
------------------------------
CREATE TABLE "donate" (
    "id" SERIAL NOT NULL,
    "picture_url" text NOT NULL,
    "description" text,
    "amount" double precision NOT NULL,
    "campaign_id" SERIAL NOT NULL,
    "user_id" uuid NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "PK_donate" PRIMARY KEY ("id"),
    CONSTRAINT "FK_donate__campaign" FOREIGN KEY (campaign_id) REFERENCES "campaign" ("id"),
    CONSTRAINT "FK_donate__user" FOREIGN KEY (user_id) REFERENCES "user" ("id"),
    CONSTRAINT "CS_donate__user_unique" UNIQUE ("user_id")
);
-- Index for search by id
CREATE INDEX IDX01_id_donate ON "donate" (id);
CREATE INDEX IDX02_user_id_donate ON "donate" (user_id);

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

    CONSTRAINT "PK_tag" PRIMARY KEY ("id"),
    CONSTRAINT "CS_tag__name" UNIQUE ("name")
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
    "user_id" uuid NOT NULL,
    lat NUMERIC(15,10) NOT NULL,
    lng NUMERIC(15,10) NOT NULL,

    CONSTRAINT "PK_pet" PRIMARY KEY ("id"),
    CONSTRAINT "FK_pet__user" FOREIGN KEY (user_id) REFERENCES "user" ("id")
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

    CONSTRAINT "PK_tag_pet" PRIMARY KEY (pet_id, tag_id),
    CONSTRAINT "FK_tag_pet__pet" FOREIGN KEY (pet_id) REFERENCES "pet" ("id"),
    CONSTRAINT "FK_tag_pet__tag" FOREIGN KEY (tag_id) REFERENCES "tag" ("id")
);

------------------------------
-- pic_pet Table
-- 1 to 1 tag and url
------------------------------
CREATE TABLE "pic_pet" (
    pet_id SERIAL NOT NULL,
    picture_url TEXT NOT NULL,

    CONSTRAINT "PK_pic_pet" PRIMARY KEY (pet_id, picture_url),
    CONSTRAINT "FK_pic_pet__pet" FOREIGN KEY (pet_id) REFERENCES "pet" ("id")
);
