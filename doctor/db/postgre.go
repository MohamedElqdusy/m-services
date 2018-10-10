package db

import (
	"context"
	"database/sql"

	"doctor/models"
	"doctor/utils"

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

func (p *PostgreRepository) CreateTables() {
	p.db.Exec("DROP TABLE doctors")
	_, err := p.db.Exec(" CREATE TABLE IF NOT EXISTS doctors(id serial NOT NULL,first_name character varying(100) ,last_name character varying(100),password character varying(100),email character varying(100) NOT NULL,created_at date,CONSTRAINT doctors_pkey PRIMARY KEY (id))")
	utils.HandleError(err)
}

func (p *PostgreRepository) RegisterDoctor(ctx context.Context, doctor models.Doctor) error {
	_, err := p.db.Exec("INSERT INTO doctors(first_name, last_name, password, email, created_at) VALUES($1, $2, $3, $4, $5)", doctor.FirstName, doctor.LastName, doctor.Password, doctor.Email, doctor.CreatedAt)
	return err
}

func (p *PostgreRepository) FindAllDoctors(stx context.Context) ([]models.Doctor, error) {
	rows, err := p.db.Query("SELECT * FROM doctors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// extract doctors attributes
	doctors := []models.Doctor{}
	for rows.Next() {
		doctor := models.Doctor{}
		err = rows.Scan(&doctor.Id, &doctor.FirstName, &doctor.LastName, &doctor.Password, &doctor.Email, &doctor.CreatedAt)
		doctors = append(doctors, doctor)
		utils.HandleError(err)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return doctors, nil
}

func (p *PostgreRepository) FindDoctor(stx context.Context, id uint64) (models.Doctor, error) {
	doctor := models.Doctor{}

	rows, err := p.db.Query("SELECT * FROM doctors WHERE id = $1", id)
	if err != nil {
		return doctor, err
	}
	defer rows.Close()

	// extract doctor attributes
	for rows.Next() {
		err = rows.Scan(&doctor.Id, &doctor.FirstName, &doctor.LastName, &doctor.Password, &doctor.Email, &doctor.CreatedAt)
		utils.HandleError(err)
	}
	if err != nil {
		return doctor, err
	}
	return doctor, nil
}
