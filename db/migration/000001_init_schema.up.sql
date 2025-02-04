-- Create schema
CREATE SCHEMA markdown_notes_schema;

-- Tables creation
CREATE TABLE markdown_notes_schema.notes (
    id SERIAL PRIMARY KEY,
    html text NOT NULL,
    mk text NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted boolean NOT NULL DEFAULT false
);