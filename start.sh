#!/bin/sh

set -e
echo "running db migration....************************************************************************************..."
echo `"postgresql://$DB_USER:$DB_PASSWORD@$DB_HOST:5432/$DB_NAME?sslmode=disable"`
/bankapp/migrate -path /bankapp/migrations -database "postgresql://$DB_USER:$DB_PASSWORD@$DB_HOST:5432/$DB_NAME?sslmode=disable" -verbose up

#echo `"postgresql://$DB_USER:$DB_PASSWORD@$DB_HOST:5432/$DB_NAME?sslmode=disable"`
echo "Starting the app..."
exec "$@"



