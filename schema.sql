--- Draft schema

CREATE DATABASE automobile_api;

CREATE TABLE makes (
    make_id uuid,
    name text,
    founded_at date,
    founded_in text
);

CREATE TABLE vehicle_types (
    vehicle_type_id uuid,
    name text
);

CREATE TABLE vehicles (
    vehicle_id uuid,
    vehicle_type_id uuid
    make_id uuid,
    model_name text
);
