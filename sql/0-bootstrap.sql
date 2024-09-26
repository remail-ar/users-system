-- 1. Create the users-service database
CREATE DATABASE "users-service";

-- Connect to the newly created database
\c "users-service";

-- 2. Create the users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,          -- Auto-incrementing ID
    name VARCHAR(100) NOT NULL,     -- User's name (non-null)
    email VARCHAR(100) UNIQUE NOT NULL,  -- Unique email address (non-null)
    created_at TIMESTAMPTZ DEFAULT NOW()  -- Timestamp when the user is created
);