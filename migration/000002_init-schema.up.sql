CREATE SCHEMA users;

CREATE TABLE users.users(
    id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    status integer NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)