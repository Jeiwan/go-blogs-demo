#!/usr/bin/env bash

migrate -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path /app/migrations/goblogs up && \
    migrate -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/tenant?sslmode=disable" -path /app/migrations/tenant up && \
    ./goblogs web