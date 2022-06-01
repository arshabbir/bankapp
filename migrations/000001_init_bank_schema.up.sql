CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "owner_name" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "created_at" timestamptz DEFAULT 'now()',
  "currency" varchar
);

CREATE TABLE "entries" (
  "id" bigint PRIMARY KEY,
  "account_id" bigserial NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "transactions" (
  "id" bigint PRIMARY KEY,
  "from_account_id" bigserial NOT NULL,
  "to_account_id" bigserial NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE INDEX ON "account" ("owner_name");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transactions" ("from_account_id");

CREATE INDEX ON "transactions" ("to_account_id");

CREATE INDEX ON "transactions" ("from_account_id", "to_account_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");
