CREATE TYPE "category" AS ENUM (
  'npwp',
  'paspor',
  'asabri'
);

CREATE TABLE "personnel" (
  "id" BIGSERIAL PRIMARY KEY,
  "nama" varchar,
  "nrp" varchar,
  "tmt_masuk" date,
  "tmt_perwira" date,
  "pangkat" varchar,
  "korps" varchar,
  "profesi" varchar,
  "spesialisasi" varchar,
  "tanggal_lahir" date,
  "kesatuan" varchar,
  "jabatan" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "npwp" (
  "id" BIGSERIAL PRIMARY KEY,
  "personnel_id" int,
  "npwp" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "paspor" (
  "id" BIGSERIAL PRIMARY KEY,
  "personnel_id" int,
  "nomor_passport" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "asabri" (
  "id" BIGSERIAL PRIMARY KEY,
  "personnel_id" int,
  "nomor_asabri" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "lampiran" (
  "id" BIGSERIAL PRIMARY KEY,
  "kategori" category,
  "personnel_id" int,
  "link" varchar,
  "nama" varchar,
  "keterangan" varchar,
  "tipe" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "npwp" ADD FOREIGN KEY ("personnel_id") REFERENCES "personnel" ("id");

ALTER TABLE "paspor" ADD FOREIGN KEY ("personnel_id") REFERENCES "personnel" ("id");

ALTER TABLE "asabri" ADD FOREIGN KEY ("personnel_id") REFERENCES "personnel" ("id");

ALTER TABLE "lampiran" ADD FOREIGN KEY ("personnel_id") REFERENCES "personnel" ("id");

CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Apply to "personnel"
CREATE TRIGGER set_updated_at_personnel
BEFORE UPDATE ON personnel
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- Apply to "npwp"
CREATE TRIGGER set_updated_at_npwp
BEFORE UPDATE ON npwp
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- Apply to "paspor"
CREATE TRIGGER set_updated_at_paspor
BEFORE UPDATE ON paspor
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- Apply to "asabri"
CREATE TRIGGER set_updated_at_asabri
BEFORE UPDATE ON asabri
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- Apply to "lampiran"
CREATE TRIGGER set_updated_at_lampiran
BEFORE UPDATE ON lampiran
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();
