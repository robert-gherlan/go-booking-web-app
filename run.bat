go build -o bookings.exe ./cmd/web/.
bookings.exe -dbhost=localhost -dbport=5432 -dbname=bookings -dbuser=postgres -dbpass=postgres