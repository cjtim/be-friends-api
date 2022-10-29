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
-- status lookup Table
------------------------------
CREATE TABLE "status" (
    name VARCHAR(255) NOT NULL,

    CONSTRAINT "PK_status" PRIMARY KEY (name)
);

------------------------------
-- pet Table
-- 1 to 1 tag and user
------------------------------
CREATE TABLE "pet" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255),
    lat NUMERIC(15,10) NOT NULL,
    lng NUMERIC(15,10) NOT NULL,

    "user_id" uuid NOT NULL,

    "status" VARCHAR(255) NOT NULL,

    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "PK_pet" PRIMARY KEY ("id"),
    CONSTRAINT "FK_pet__user" FOREIGN KEY (user_id) REFERENCES "user" ("id"),
    CONSTRAINT "FK_pet__status" FOREIGN KEY (status) REFERENCES "status" ("name")
);
-- Search by id and name
CREATE INDEX IDX01_id_pet ON "pet" (id);
CREATE INDEX IDX02_name_pet ON "pet" (name);
CREATE INDEX IDX03_status_pet ON "pet" (status);

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

CREATE TABLE "interested" (
    pet_id SERIAL NOT NULL,
    user_id uuid NOT NULL,
    step VARCHAR(255) DEFAULT 'ได้รับข้อมูลแล้ว' NOT NULL,

    CONSTRAINT "PK_interested" PRIMARY KEY (pet_id, user_id),
    CONSTRAINT "FK_interested__pet" FOREIGN KEY (pet_id) REFERENCES "pet" ("id"),
    CONSTRAINT "FK_interested__user" FOREIGN KEY (user_id) REFERENCES "user" ("id")
);

CREATE TABLE "liked" (
    pet_id SERIAL NOT NULL,
    user_id uuid NOT NULL,

    CONSTRAINT "PK_liked" PRIMARY KEY (pet_id, user_id),
    CONSTRAINT "FK_liked__pet" FOREIGN KEY (pet_id) REFERENCES "pet" ("id"),
    CONSTRAINT "FK_liked__user" FOREIGN KEY (user_id) REFERENCES "user" ("id")
);