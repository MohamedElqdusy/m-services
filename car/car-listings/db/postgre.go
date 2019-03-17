package db

import (
	"context"
	"database/sql"
	"car-listings/models"
	"car-listings/utils"
	_ "github.com/lib/pq"
)

type PostgreRepository struct {
	db *sql.DB
}

func NewPostgre(url string) (*PostgreRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgreRepository{
		db,
	}, nil
}

func (p *PostgreRepository) Close() {
	p.db.Close()
}

func (p *PostgreRepository) StoreCar(ctx context.Context, car models.Car, dealer_id uint64) error {
	_, err := p.db.Exec("INSERT INTO cars(dealer_id, code, make, model, kw, year, color, price) VALUES($1, $2, $3, $4, $5, $6, $7, $8)ON CONFLICT ON CONSTRAINT cars_pkey DO UPDATE SET code=$2, make=$3, model=$4, kw=$5, year=$6, color=$7, price=$8",
	dealer_id, car.Code, car.Make, car.Model, car.Kw, car.Year, car.Color, car.Price)
	return err
}

func (p *PostgreRepository) FindCarsByMake(ctx context.Context, make string) ([]models.Car, error) {
	rows, err := p.db.Query("SELECT code, make, model, kw, year, color, price FROM cars WHERE make = $1",make)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return extractCarsAttribute(rows)
}

func (p *PostgreRepository) FindCarsByModel(ctx context.Context, model string) ([]models.Car, error) {
	rows, err := p.db.Query("SELECT code, make, model, kw, year, color, price FROM cars WHERE model = $1",model)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return extractCarsAttribute(rows)
}

func (p *PostgreRepository) FindCarsByYear(ctx context.Context, year uint64) ([]models.Car, error) {
	rows, err := p.db.Query("SELECT code, make, model, kw, year, color, price FROM cars WHERE year = $1",year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return extractCarsAttribute(rows)
}

func (p *PostgreRepository) FindCarsByColor(ctx context.Context, color string) ([]models.Car, error) {
	rows, err := p.db.Query("SELECT code, make, model, kw, year, color, price FROM cars WHERE color = $1",color)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return extractCarsAttribute(rows)
}

func (p *PostgreRepository) FindAllCars(ctx context.Context) ([]models.Car, error) {
	rows, err := p.db.Query("SELECT code, make, model, kw, year, color, price FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return extractCarsAttribute(rows)
}

func extractCarsAttribute(rows *sql.Rows) ([]models.Car, error){
	cars := []models.Car{}
	for rows.Next() {
		car := models.Car{}
		err := rows.Scan(&car.Code, &car.Make, &car.Model, &car.Kw, &car.Year, &car.Color, &car.Price)
		cars = append(cars, car)
		utils.HandleError(err)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}

