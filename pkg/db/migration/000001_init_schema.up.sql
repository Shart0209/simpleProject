CREATE TABLE "documents" (
  "id" int PRIMARY KEY,
  "name" varchar NOT NULL,
  "create_gk" timestamp NOT NULL,
  "description" text,
  "filesname" int,
  "start_date" timestamp,
  "end_date" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "distributor" varchar,
  "method" varchar,
  "price" float NOT NULL
);

CREATE TABLE "files" (
  "id" int PRIMARY KEY,
  "file_name" varchar,
  "file_size" int
);


ALTER TABLE "documents" ADD FOREIGN KEY ("filesname") REFERENCES "files" ("id");
