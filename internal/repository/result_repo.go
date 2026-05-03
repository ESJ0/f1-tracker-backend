package repository

import (
	"database/sql"

	"f1-tracker-backend/internal/models"
)

type ResultRepo struct {
	db *sql.DB
}

func NewResultRepo(db *sql.DB) *ResultRepo {
	return &ResultRepo{db: db}
}

func (r *ResultRepo) GetByRace(raceID int) ([]models.ResultDetail, error) {
	query := `
		SELECT 
			res.id, res.driver_id, res.race_id,
			res.position, res.points, res.fastest_lap, res.created_at,
			d.name AS driver_name, d.team AS driver_team, d.number AS driver_number,
			ra.grand_prix, ra.circuit, ra.country, ra.race_date
		FROM results res
		JOIN drivers d  ON d.id  = res.driver_id
		JOIN races   ra ON ra.id = res.race_id
		WHERE res.race_id = $1
		ORDER BY res.position ASC
	`
	return r.scanResults(r.db.Query(query, raceID))
}

func (r *ResultRepo) GetByDriver(driverID int) ([]models.ResultDetail, error) {
	query := `
		SELECT 
			res.id, res.driver_id, res.race_id,
			res.position, res.points, res.fastest_lap, res.created_at,
			d.name AS driver_name, d.team AS driver_team, d.number AS driver_number,
			ra.grand_prix, ra.circuit, ra.country, ra.race_date
		FROM results res
		JOIN drivers d  ON d.id  = res.driver_id
		JOIN races   ra ON ra.id = res.race_id
		WHERE res.driver_id = $1
		ORDER BY ra.race_date DESC
	`
	return r.scanResults(r.db.Query(query, driverID))
}

func (r *ResultRepo) Create(res *models.Result) (*models.Result, error) {
	query := `
		INSERT INTO results (driver_id, race_id, position, points, fastest_lap)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, driver_id, race_id, position, points, fastest_lap, created_at
	`
	var created models.Result
	err := r.db.QueryRow(
		query, res.DriverID, res.RaceID, res.Position, res.Points, res.FastestLap,
	).Scan(
		&created.ID, &created.DriverID, &created.RaceID,
		&created.Position, &created.Points, &created.FastestLap, &created.CreatedAt,
	)
	return &created, err
}

func (r *ResultRepo) Delete(id int) (bool, error) {
	result, err := r.db.Exec("DELETE FROM results WHERE id=$1", id)
	if err != nil {
		return false, err
	}
	rows, _ := result.RowsAffected()
	return rows > 0, nil
}

// helper para no repetir el scan
func (r *ResultRepo) scanResults(rows *sql.Rows, err error) ([]models.ResultDetail, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.ResultDetail
	for rows.Next() {
		var rd models.ResultDetail
		if err := rows.Scan(
			&rd.ID, &rd.DriverID, &rd.RaceID,
			&rd.Position, &rd.Points, &rd.FastestLap, &rd.CreatedAt,
			&rd.DriverName, &rd.DriverTeam, &rd.DriverNumber,
			&rd.GrandPrix, &rd.Circuit, &rd.Country, &rd.RaceDate,
		); err != nil {
			return nil, err
		}
		results = append(results, rd)
	}
	return results, nil
}
