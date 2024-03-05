#!/bin/bash

set -e

# Check if the database exists

if [ ! -d /var/lib/postgresql/data ]; then

  # Initialize the database

  echo "Initializing the database..."

  /usr/local/bin/postgres --single -c "CREATE DATABASE postgres"

  # Create the user and grant permissions

  echo "Creating the user and granting permissions..."

  /usr/local/bin/psql -U postgres -d postgres -c "CREATE USER postgres WITH PASSWORD 'postgres'"

  /usr/local/bin/psql -U postgres -d postgres -c "GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres"

  # Apply the migrations

  echo "Applying the migrations..."

  for file in /migrations/*.sql; do
    /usr/local/bin/psql -U postgres -d postgres -f "$file"
  done

fi

# Start the server

echo "Starting the server..."

/usr/local/bin/postgres