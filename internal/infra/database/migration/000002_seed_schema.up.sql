INSERT INTO drivers (name, email, tax_id, driver_license, date_of_birth)
VALUES
    ('Motorista 1', 'motorista1@example.com', '12345678901', 'ABC12345', '1990-01-01'),
    ('Motorista 2', 'motorista2@example.com', '23456789012', 'XYZ54321', '1985-03-15'),
    ('Motorista 3', 'motorista3@example.com', '34567890123', 'DEF67890', '1995-07-20'),
    ('Motorista 4', 'motorista4@example.com', '45678901234', 'GHI56789', '1988-12-10'),
    ('Motorista 5', 'motorista5@example.com', '56789012345', 'JKL98765', '1992-06-05');

-- Insert 5 vehicles with truck names
INSERT INTO vehicles (brand, model, year_of_manufacture, license_plate, color)
VALUES
    ('Scania', 'R500', 2020, 'ABC123', 'Blue'),
    ('Volvo', 'FH16', 2019, 'XYZ987', 'Red'),
    ('Mercedes-Benz', 'Actros', 2021, 'DEF456', 'Silver'),
    ('MAN', 'TGX', 2018, 'GHI789', 'Black'),
    ('Kenworth', 'W900', 2022, 'JKL321', 'White');

-- Insert Relation drivers and vehicles
INSERT INTO drivers_vehicles (driver_uuid, vehicle_uuid)
VALUES
    ((SELECT uuid FROM drivers WHERE email = 'motorista1@example.com'), (SELECT uuid FROM vehicles WHERE license_plate = 'ABC123')),
    ((SELECT uuid FROM drivers WHERE email = 'motorista2@example.com'), (SELECT uuid FROM vehicles WHERE license_plate = 'XYZ987')),
    ((SELECT uuid FROM drivers WHERE email = 'motorista3@example.com'), (SELECT uuid FROM vehicles WHERE license_plate = 'DEF456')),
    ((SELECT uuid FROM drivers WHERE email = 'motorista4@example.com'), (SELECT uuid FROM vehicles WHERE license_plate = 'GHI789')),
    ((SELECT uuid FROM drivers WHERE email = 'motorista5@example.com'), (SELECT uuid FROM vehicles WHERE license_plate = 'JKL321'));
