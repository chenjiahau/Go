CREATE TABLE "user_tokens" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY,
  "user_id" int NOT NULL,
  "token" varchar NOT NULL,
  "expired_at" timestamp NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("id")
);

ALTER TABLE "user_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");