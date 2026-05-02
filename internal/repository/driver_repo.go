package repository

import (
	"database/sql"
	"fmt"

	"f1-tracker-backend/internal/models"
)

type DriverRepo struct {
	db *sql.DB
}

func NewDriverRepo(db *sql.DB) *DriverRepo {
	return &DriverRepo{db: db}
}

type DriverFilter struct {
	Search string
	Sort   string
	Order  string
	Page   int
	Limit  int
}

func (r *DriverRepo) GetAll(f DriverFilter) ([]models.Driver, int, error) {
	// Columnas válidas para ordenar
	validSort := map[string]bool{
		"name": true, "team": true, "number": true, "created_at": true,
	}
	sortCol := "created_at"
	if validSort[f.Sort] {
		sortCol = f.Sort
	}

	order := "ASC"
	if f.Order == "desc" {
		order = "DESC"
	}

	if f.Page < 1 {
		f.Page = 1
	}
	if f.Limit < 1 {
		f.Limit = 10
	}
	offset := (f.Page - 1) * f.Limit

	// Contar total
	countQuery := "SELECT COUNT(*) FROM drivers WHERE name ILIKE $1"
	var total int
	r.db.QueryRow(countQuery, "%"+f.Search+"%").Scan(&total)

	query := fmt.Sprintf(`
		SELECT id, name, team, nationality, number, COALESCE(image_url, ''), created_at, updated_at
		FROM drivers
		WHERE name ILIKE $1
		ORDER BY %s %s
		LIMIT $2 OFFSET $3
	`, sortCol, order)

	rows, err := r.db.Query(query, "%"+f.Search+"%", f.Limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var drivers []models.Driver
	for rows.Next() {
		var d models.Driver
		if err := rows.Scan(
			&d.ID, &d.Name, &d.Team, &d.Nationality,
			&d.Number, &d.ImageURL, &d.CreatedAt, &d.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		drivers = append(drivers, d)
	}

	return drivers, total, nil
}

func (r *DriverRepo) GetByID(id int) (*models.Driver, error) {
	query := `
		SELECT id, name, team, nationality, number, COALESCE(image_url, ''), created_at, updated_at
		FROM drivers WHERE id = $1
	`
	var d models.Driver
	err := r.db.QueryRow(query, id).Scan(
		&d.ID, &d.Name, &d.Team, &d.Nationality,
		&d.Number, &d.ImageURL, &d.CreatedAt, &d.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &d, err
}

func (r *DriverRepo) Create(d *models.Driver) (*models.Driver, error) {
	query := `
		INSERT INTO drivers (name, team, nationality, number, image_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, team, nationality, number, COALESCE(image_url, ''), created_at, updated_at
	`
	var created models.Driver
	err := r.db.QueryRow(query, d.Name, d.Team, d.Nationality, d.Number, d.ImageURL).Scan(
		&created.ID, &created.Name, &created.Team, &created.Nationality,
		&created.Number, &created.ImageURL, &created.CreatedAt, &created.UpdatedAt,
	)
	return &created, err
}

func (r *DriverRepo) Update(id int, d *models.Driver) (*models.Driver, error) {
	query := `
		UPDATE drivers
		SET name=$1, team=$2, nationality=$3, number=$4, image_url=$5, updated_at=NOW()
		WHERE id=$6
		RETURNING id, name, team, nationality, number, COALESCE(image_url, ''), created_at, updated_at
	`
	var updated models.Driver
	err := r.db.QueryRow(query, d.Name, d.Team, d.Nationality, d.Number, d.ImageURL, id).Scan(
		&updated.ID, &updated.Name, &updated.Team, &updated.Nationality,
		&updated.Number, &updated.ImageURL, &updated.CreatedAt, &updated.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &updated, err
}

func (r *DriverRepo) Delete(id int) (bool, error) {
	result, err := r.db.Exec("DELETE FROM drivers WHERE id=$1", id)
	if err != nil {
		return false, err
	}
	rows, _ := result.RowsAffected()
	return rows > 0, nil
}
