#!/bin/bash

# Load environment variables from .env file
set -a
source ../.env.local
set +a

# Function to URL encode strings
urlencode() {
    local string="${1}"
    local strlen=${#string}
    local encoded=""
    local pos c o

    for (( pos=0 ; pos<strlen ; pos++ )); do
        c=${string:$pos:1}
        case "$c" in
            [-_.~a-zA-Z0-9] ) o="${c}" ;;
            * )               printf -v o '%%%02x' "'$c"
        esac
        encoded+="${o}"
    done
    echo "${encoded}"
}

# Construct database URL with encoded credentials
DB_USER_ENCODED=$(urlencode "${DB_USER}")
DB_PASS_ENCODED=$(urlencode "${DB_PASSWORD}")
DB_URL="postgres://${DB_USER_ENCODED}:${DB_PASS_ENCODED}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=require"

echo "Database URL: ${DB_URL}"

# Function to show usage
usage() {
    echo "Usage: $0 <command>"
    echo "Commands:"
    echo "  up        Run all migrations"
    echo "  down      Revert all migrations"
    echo "  force     Force to specific version (requires version number)"
    echo "  create    Create new migration (requires name)"
    echo "  version   Show current migration version"
    exit 1
}

# Check if command is provided
if [ "$#" -lt 1 ]; then
    usage
fi

# Execute command
case "$1" in
    "up")
        echo "Running migrations up..."
        migrate -path ../migrations -database "${DB_URL}" up
        ;;
    "down")
        echo "Running migrations down..."
        migrate -path ../migrations -database "${DB_URL}" down
        ;;
    "force")
        if [ -z "$2" ]; then
            echo "Error: Version number required for force command"
            exit 1
        fi
        echo "Forcing migration version..."
        migrate -path ../migrations -database "${DB_URL}" force "$2"
        ;;
    "create")
        if [ -z "$2" ]; then
            echo "Error: Migration name required for create command"
            exit 1
        fi
        echo "Creating new migration..."
        migrate create -ext sql -dir ../migrations -seq "$2"
        ;;
    "version")
        echo "Checking migration version..."
        migrate -path ../migrations -database "${DB_URL}" version
        ;;
    *)
        usage
        ;;
esac 