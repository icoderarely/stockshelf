CREATE TABLE
  refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    token TEXT UNIQUE NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    is_revoked BOOLEAN DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT now ()
  );

CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens (user_id);