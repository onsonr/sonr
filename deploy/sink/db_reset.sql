-- Connect to a different database first (postgres) since we can't drop a database while connected to it
\c postgres;

-- Terminate all connections to the databases
SELECT pg_terminate_backend(pid) 
FROM pg_stat_activity 
WHERE datname IN ('chainindex', 'highway', 'matrixhs')
AND pid <> pg_backend_pid();

-- Drop the databases if they exist
DROP DATABASE IF EXISTS chainindex;
DROP DATABASE IF EXISTS highway;
DROP DATABASE IF EXISTS matrixhs;

-- Drop the users if they exist
DROP USER IF EXISTS chainindex_user;
DROP USER IF EXISTS highway_user;
DROP USER IF EXISTS matrixhs_user;

