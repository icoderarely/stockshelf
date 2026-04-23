CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE
  users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role text not null default 'user' CHECK (role IN ('user', 'admin')),
    created_at TIMESTAMPTZ DEFAULT NOW (),
    updated_at TIMESTAMPTZ DEFAULT NOW ()
  );