-- Switch to postgres database first
\c postgres;

-- Terminate existing connections
SELECT pg_terminate_backend(pid) 
FROM pg_stat_activity 
WHERE datname = 'chainindex'
AND pid <> pg_backend_pid();

-- Drop and recreate database and user
DROP DATABASE IF EXISTS chainindex;
DROP USER IF EXISTS chainindex_user;

CREATE USER chainindex_user WITH PASSWORD 'chainindex_password123';
CREATE DATABASE chainindex;

-- Connect to the new database
\c chainindex;

-- Create the blocks table
CREATE TABLE blocks (
  rowid      BIGSERIAL PRIMARY KEY,
  height     BIGINT NOT NULL,
  chain_id   VARCHAR NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  UNIQUE (height, chain_id)
);

CREATE INDEX idx_blocks_height_chain ON blocks(height, chain_id);

CREATE TABLE tx_results (
  rowid BIGSERIAL PRIMARY KEY,
  block_id BIGINT NOT NULL REFERENCES blocks(rowid),
  index INTEGER NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  tx_hash VARCHAR NOT NULL,
  tx_result BYTEA NOT NULL,
  UNIQUE (block_id, index)
);

CREATE TABLE events (
  rowid BIGSERIAL PRIMARY KEY,
  block_id BIGINT NOT NULL REFERENCES blocks(rowid),
  tx_id    BIGINT NULL REFERENCES tx_results(rowid),
  type VARCHAR NOT NULL
);

CREATE TABLE attributes (
   event_id      BIGINT NOT NULL REFERENCES events(rowid),
   key           VARCHAR NOT NULL,
   composite_key VARCHAR NOT NULL,
   value         VARCHAR NULL,
   UNIQUE (event_id, key)
);

CREATE VIEW event_attributes AS
  SELECT block_id, tx_id, type, key, composite_key, value
  FROM events LEFT JOIN attributes ON (events.rowid = attributes.event_id);

CREATE VIEW block_events AS
  SELECT blocks.rowid as block_id, height, chain_id, type, key, composite_key, value
  FROM blocks JOIN event_attributes ON (blocks.rowid = event_attributes.block_id)
  WHERE event_attributes.tx_id IS NULL;

CREATE VIEW tx_events AS
  SELECT height, index, chain_id, type, key, composite_key, value, tx_results.created_at
  FROM blocks JOIN tx_results ON (blocks.rowid = tx_results.block_id)
  JOIN event_attributes ON (tx_results.rowid = event_attributes.tx_id)
  WHERE event_attributes.tx_id IS NOT NULL;

-- Grant all necessary privileges
GRANT ALL PRIVILEGES ON DATABASE chainindex TO chainindex_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO chainindex_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO chainindex_user;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO chainindex_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO chainindex_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON SEQUENCES TO chainindex_user;

