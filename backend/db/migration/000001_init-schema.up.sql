CREATE TABLE "users" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY,
  "email" varchar NOT NULL UNIQUE,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("id")
);