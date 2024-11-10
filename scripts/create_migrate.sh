#!/bin/bash

clear
cd backend/pkg/db/migrations/sqlite
migrate create -ext sql test