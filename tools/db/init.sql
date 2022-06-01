CREATE TABLE "Users" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "name" varchar(255) NOT NULL,
    "email" varchar(255) CONSTRAINT Users_email UNIQUE,
    "password" varchar(255),
    "line_uid" varchar(255) CONSTRAINT Users_line_uid UNIQUE,
    "picture_url" text,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT "Users_pk" PRIMARY KEY ("id")
)
