ALTER TABLE refresh_tokens
ADD COLUMN expired_token TIMESTAMP NOT NULL;

