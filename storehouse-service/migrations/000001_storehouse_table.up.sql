CREATE TABLE IF NOT EXISTS "storehouses" (
  "id" uuid PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "price" float NOT NULL,
  "amount" int DEFAULT 0,
  "type_artillery" varchar NOT NULL,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE  IF NOT EXISTS "resource_usage" (
  "id" uuid PRIMARY KEY,
  "soldier_id" uuid NOT NULL,
  "storage_id" uuid NOT NULL,
  "amount" int DEFAULT 0,
  "total_price" float NOT NULL,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP,
  "deleted_at" timestamp
);