#!/bin/sh
# wait-for-postgres.sh

set -e

shift
cmd="$@"

until PGPASSWORD=wb_pass psql -h "db" -U "wb_user" -d "wb_db" -c '\q'; do
    >&2 echo "Postgres is unavailable - sleeping"
    sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd