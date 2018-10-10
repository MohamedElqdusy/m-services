package db

import (
	"context"
	"database/sql"

	"patient/models"
	"patient/utils"

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

func (p *PostgreRepository) RegisterPatient(ctx context.Context, patient models.Patient) error {
	_, err := p.db.Exec("INSERT INTO patients(first_name, last_name, password, email, created_at) VALUES($1, $2, $3, $4, $5)", patient.FirstName, patient.LastName, patient.Password, patient.Email, patient.CreatedAt)
	return err
}

func (p *PostgreRepository) FindAllPatients(stx context.Context) ([]models.Patient, error) {
	rows, err := p.db.Query("SELECT * FROM patients")
	utils.HandleError(err)
	defer rows.Close()

	// extract patients attributes
	patients := []models.Patient{}
	for rows.Next() {
		patient := models.Patient{}
		err = rows.Scan(&patient.Id, &patient.FirstName, &patient.LastName, &patient.Password, &patient.Email, &patient.CreatedAt)
		patients = append(patients, patient)
		utils.HandleError(err)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}

func (p *PostgreRepository) FindPatient(stx context.Context, id uint64) (models.Patient, error) {
	patient := models.Patient{}

	rows, err := p.db.Query("SELECT * FROM patients WHERE id = $1", id)
	if err != nil {
		return patient, err
	}
	defer rows.Close()

	// extract patient attributes
	for rows.Next() {
		err = rows.Scan(&patient.Id, &patient.FirstName, &patient.LastName, &patient.Password, &patient.Email, &patient.CreatedAt)
		utils.HandleError(err)
	}
	if err != nil {
		return patient, err
	}
	return patient, nil
}
