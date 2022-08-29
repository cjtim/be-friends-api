------------------------------
-- Shelter
------------------------------
-- CREATE TABLE "shelter" (
--     "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
--     lat NUMERIC(15,10) NOT NULL,
--     lng NUMERIC(15,10) NOT NULL,

--     "created_at" timestamp NOT NULL DEFAULT NOW(),
--     "updated_at" timestamp NOT NULL DEFAULT NOW(),

--     CONSTRAINT "PK_shelter" PRIMARY KEY ("id")
-- );
-- -- Index for search by id
-- CREATE INDEX IDX01_id_shelter ON "shelter" (id);

------------------------------
-- User
------------------------------
CREATE TABLE "user" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255),
    "password" VARCHAR(255),
    "line_uid" VARCHAR(255),
    "description" text,
    "picture_url" text,
    "phone" text NOT NULL,
    "is_org" BOOL DEFAULT FALSE NOT NULL,
    "is_admin" BOOL DEFAULT FALSE NOT NULL,

    lat NUMERIC(15,10),
    lng NUMERIC(15,10),

    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "PK_user" PRIMARY KEY ("id"),
    CONSTRAINT "CS_user__email" UNIQUE ("email"),
    CONSTRAINT "CS_user__line_uid" UNIQUE ("line_uid")
);
-- Index for search by id
CREATE INDEX IDX01_id_user ON "user" (id);
CREATE INDEX IDX02_line_uid_user ON "user" (line_uid);
CREATE INDEX IDX03_email_user ON "user" (email);