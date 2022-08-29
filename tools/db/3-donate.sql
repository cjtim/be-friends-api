------------------------------
-- Campaign
------------------------------
-- CREATE TABLE "campaign" (
--     "id" SERIAL NOT NULL,
--     "name" VARCHAR(255) NOT NULL,
--     "picture_url" text,
--     "description" text,
--     "min_amount" double precision NOT NULL,

--     "shelter_id" uuid NOT NULL,

--     "created_at" timestamp NOT NULL DEFAULT NOW(),
--     "updated_at" timestamp NOT NULL DEFAULT NOW(),

--     CONSTRAINT "PK_campaign" PRIMARY KEY ("id"),
--     CONSTRAINT "CS_campaign__name" UNIQUE ("name"),
--     CONSTRAINT "FK_campaign__shelter" FOREIGN KEY (shelter_id) REFERENCES "shelter" ("id")
-- );
-- -- Index for search by id
-- CREATE INDEX IDX01_id_campaign ON "campaign" (id);
-- CREATE INDEX IDX02_name_campaign ON "campaign" (name);
-- CREATE INDEX IDX03_shelter_id_campaign ON "campaign" (shelter_id);

------------------------------
-- donate
------------------------------
CREATE TABLE "donate" (
    "id" SERIAL NOT NULL,
    "picture_url" text NOT NULL,
    "description" text,
    "amount" double precision NOT NULL,

    "user_id" uuid NOT NULL,

    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "PK_donate" PRIMARY KEY ("id"),
    CONSTRAINT "FK_donate__user" FOREIGN KEY (user_id) REFERENCES "user" ("id")
);
-- Index for search by id
CREATE INDEX IDX01_id_donate ON "donate" (id);
CREATE INDEX IDX02_user_id_donate ON "donate" (user_id);
