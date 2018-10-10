package db

import (
	"context"
	"database/sql"

	"appoinment/models"
	"appoinment/utils"

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

func (p *PostgreRepository) ReserveAppoinment(ctx context.Context, appoinment models.Appoinment) error {
	_, err := p.db.Exec("INSERT INTO appoinments(doctor_Id, time_point, patient_Id, notes) VALUES($1, $2, $3, $4)", appoinment.DoctorId, appoinment.TimePoint, appoinment.PatientId, appoinment.Notes)
	return err
}

func (p *PostgreRepository) FindAllDoctorAppoinments(stx context.Context, doctorId int) ([]models.Appoinment, error) {
	rows, err := p.db.Query("SELECT * FROM appoinments WHERE doctor_Id = $1", doctorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// extract appoinments attributes
	appoinments := []models.Appoinment{}
	for rows.Next() {
		appoinment := models.Appoinment{}
		err = rows.Scan(&appoinment.DoctorId, &appoinment.TimePoint, &appoinment.PatientId, &appoinment.Notes)
		appoinments = append(appoinments, appoinment)
		utils.HandleError(err)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return appoinments, nil
}
