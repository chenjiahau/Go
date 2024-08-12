CREATE TABLE "member_roles" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY,
  "title" varchar NOT NULL,
  "abbr" varchar NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "members" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY,
  "member_role_id" INT NOT NULL,
  "name" varchar NOT NULL,
  "is_alive" boolean NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("id"),
  FOREIGN KEY ("member_role_id") REFERENCES "member_roles" ("id")
);

CREATE TABLE "user_members" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY,
  "user_id" INT NOT NULL,
  "member_id" INT NOT NULL,
  PRIMARY KEY ("id")
);

ALTER TABLE "user_members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "user_members" ADD FOREIGN KEY ("member_id") REFERENCES "members" ("id") ON DELETE CASCADE ON UPDATE CASCADE;