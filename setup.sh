#!/bin/bash

# Update go.mod dependencies
go mod tidy

# Start Docker services
docker compose up -d
