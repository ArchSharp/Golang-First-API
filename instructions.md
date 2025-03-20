docker build -t golang-first-api
docker run -p 8080:8080 -tid golang-first-api
GIN_MODE=release

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=Alade1&&&
DB_NAME=postgres
DB_SSLMODE=disable

BASEURL=https://api.flutterwave.com/v3/bill-categories
FLWSECK_TEST=FLWSECK_TEST-cf7e992ac80002f5ab21ada1e87f94c3-X
