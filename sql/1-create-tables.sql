-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,                   -- Auto-incrementing primary key
    username VARCHAR(255) NOT NULL UNIQUE,   -- Username must be unique and not null
    password TEXT NOT NULL,                  -- Password field, storing as TEXT for simplicity
    created_at TIMESTAMPTZ DEFAULT NOW()     -- Timestamp for when the user was created
);

-- Create an index on the username column for faster lookups
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);

-- Add a unique constraint on the username column
ALTER TABLE users ADD CONSTRAINT unique_username UNIQUE(username);

-- Ensure there's at least one user by defining a rule to deny user deletion if only one user exists
CREATE OR REPLACE FUNCTION ensure_user_count() RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT COUNT(*) FROM users) = 1 THEN
        RAISE EXCEPTION 'Cannot delete the only remaining user';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

-- Trigger to prevent deleting the last user in the system
CREATE TRIGGER prevent_last_user_deletion
    BEFORE DELETE ON users
    FOR EACH ROW EXECUTE FUNCTION ensure_user_count();

-- Add a dummy admin user for initial setup (optional)
INSERT INTO users (username, password)
VALUES ('admin', 'adminpassword')
ON CONFLICT DO NOTHING;  -- Ensures this insertion runs only if no user exists
