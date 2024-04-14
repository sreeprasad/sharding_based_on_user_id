#!/bin/bash

set -e

echo "creating user 6 for sp"


psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
    DO \$\$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'user5') THEN
            CREATE ROLE user5 WITH LOGIN PASSWORD 'password5';
        ELSE
            RAISE NOTICE 'Role user5 already exists. Skipping creation.';
        END IF;
    END
    \$\$;
EOSQL




