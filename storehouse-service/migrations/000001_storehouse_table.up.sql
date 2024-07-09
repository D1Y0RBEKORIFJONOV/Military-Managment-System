CREATE TABLE "storehouses" (
  "id" uuid PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "price" float NOT NULL,
  "amount" int DEFAULT 0,
  "type_artillery" varchar NOT NULL,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE "storehouse" (
  "id" uuid PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "price" float NOT NULL,
  "amount" int DEFAULT 0,
  "type_artillery" varchar NOT NULL,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW()),
  "deleted_at" timestamp
);