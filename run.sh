#!/bin/bash

go build -o bookings cmd/web/*.go && ./bookings -dbhost=localhost -dbport=5432 -dbname=bookings -dbuser=postgres -dbpass=postgres