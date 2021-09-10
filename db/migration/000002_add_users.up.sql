CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hash_password" varchar NOT NULL,
  "password_changed_at" timestamptz DEFAULT '0001-01-01 00:00:00',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");


-- one owner only have one kind of currency
-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");

ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");