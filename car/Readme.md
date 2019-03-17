# README

## System Requirements
( Golang:1.11, Docker )
# Run
### `$ docker-compose up`
##### End points:
- `http://localhost:4422/` Welcome
* `http://localhost:4422/upload_csv/:dealer_id/` Store the cars list from a csv file given the dealer_id.
  - Example:
` curl -i -X POST -F 'csv=@c.csv' http://localhost:4422/upload_csv/1`

- `http://localhost:4422/vehicle_listings/` Store the cars list fror other providers, dealer_id=0 implicitly
  - Example:
`curl -H 'Content-Type: application/json' -X POST -d '[{"code":"a", "make":"renault", "model":"megane", "Kw":132, "year":2015, "color":"red", "price":2584}]' http://localhost:4422/vehicle_listings/`

- `http://localhost:4422/cars/search/` Get all the car listings

- `http://localhost:4422/cars/color/:color/`
Search cars by color
- `http://localhost:4422/cars/make/:make/`
Search cars by make
- `http://localhost:4422/cars/model/:model/`
Search cars by model
- `http://localhost:4422/cars/year/:year/`
Search cars by year 
# Tests
Run the tests
- `go test ./...`