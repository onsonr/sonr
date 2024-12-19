-- Connect to postgres default database
\c postgres;

-- Create databases
CREATE DATABASE chainindex;
CREATE DATABASE highway;
CREATE DATABASE matrixhs;

-- Create users with passwords
CREATE USER chainindex_user WITH PASSWORD 'chainindex_password123';
CREATE USER highway_user WITH PASSWORD 'highway_password123';
CREATE USER matrixhs_user WITH PASSWORD 'matrixhs_password123';

-- Grant privileges for each database to their respective users
\c chainindex;
GRANT ALL PRIVILEGES ON DATABASE chainindex TO chainindex_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO chainindex_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO chainindex_user;

\c highway;
GRANT ALL PRIVILEGES ON DATABASE highway TO highway_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO highway_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO highway_user;

\c matrixhs;
GRANT ALL PRIVILEGES ON DATABASE matrixhs TO matrixhs_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO matrixhs_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO matrixhs_user;

