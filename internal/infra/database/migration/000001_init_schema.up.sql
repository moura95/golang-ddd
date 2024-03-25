SET TIME ZONE 'America/Sao_Paulo';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS drivers (
                         uuid          UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
                         name VARCHAR(255) NOT NULL,
                         email VARCHAR(255) NOT NULL UNIQUE,
                         tax_id varchar(255) NOT NULL UNIQUE,
                         driver_license VARCHAR(20),
                         date_of_birth VARCHAR,
                         deleted_at TIMESTAMP,
                         created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
                         update_at    TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS vehicles (
                        uuid          UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
                        brand VARCHAR(255) NOT NULL ,
                        model VARCHAR(255) NOT NULL ,
                        year_of_manufacture INTEGER,
                        license_plate VARCHAR(10) NOT NULL UNIQUE,
                        color VARCHAR(50),
                        deleted_at TIMESTAMP,
                        created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
                        update_at    TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS drivers_vehicles (
                                driver_uuid UUID,
                                vehicle_uuid UUID,
                                created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
                                UNIQUE (driver_uuid, vehicle_uuid),
                                FOREIGN KEY (driver_uuid) REFERENCES drivers(uuid),
                                FOREIGN KEY (vehicle_uuid) REFERENCES vehicles(uuid)

);

CREATE INDEX idx_drivers_email ON drivers (email);

CREATE INDEX idx_drivers_tax_id ON drivers (tax_id);

CREATE INDEX idx_vehicles_license_plate ON vehicles (license_plate);
