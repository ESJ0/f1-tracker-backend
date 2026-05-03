package repository

import (
	"database/sql"
	"fmt"

	"f1-tracker-backend/internal/models"
)

type RaceRepo struct {
	db *sql.DB
}

func NewRaceRepo(db *sql.DB) *RaceRepo {
	return &RaceRepo{db: db}
}

type RaceFilter struct {
	Search string
	Sort   string
	Order  string
	Page   int
	Limit  int
}

func (r *RaceRepo) GetAll(f RaceFilter) ([]models.Race, int, error) {
	validSort := map[string]bool{
		"grand_prix": true, "country": true, "race_date": true, "created_at": true,
	}
	sortCol := "race_date"
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

	var total int
	r.db.QueryRow(
		"SELECT COUNT(*) FROM races WHERE grand_prix ILIKE $1",
		"%"+f.Search+"%",
	).Scan(&total)

	query := fmt.Sprintf(`
		SELECT id, grand_prix, circuit, country, race_date, COALESCE(image_url, ''), created_at, updated_at
		FROM races
		WHERE grand_prix ILIKE $1
		ORDER BY %s %s
		LIMIT $2 OFFSET $3
	`, sortCol, order)

	rows, err := r.db.Query(query, "%"+f.Search+"%", f.Limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var races []models.Race
	for rows.Next() {
		var race models.Race
		if err := rows.Scan(
			&race.ID, &race.GrandPrix, &race.Circuit, &race.Country,
			&race.RaceDate, &race.ImageURL, &race.CreatedAt, &race.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		races = append(races, race)
	}

	return races, total, nil
}

func (r *RaceRepo) GetByID(id int) (*models.Race, error) {
	query := `
		SELECT id, grand_prix, circuit, country, race_date, COALESCE(image_url, ''), created_at, updated_at
		FROM races WHERE id = $1
	`
	var race models.Race
	err := r.db.QueryRow(query, id).Scan(
		&race.ID, &race.GrandPrix, &race.Circuit, &race.Country,
		&race.RaceDate, &race.ImageURL, &race.CreatedAt, &race.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &race, err
}

func (r *RaceRepo) Create(race *models.Race) (*models.Race, error) {
	query := `
		INSERT INTO races (grand_prix, circuit, country, race_date, image_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, grand_prix, circuit, country, race_date, COALESCE(image_url, ''), created_at, updated_at
	`
	var created models.Race
	err := r.db.QueryRow(
		query, race.GrandPrix, race.Circuit, race.Country, race.RaceDate, race.ImageURL,
	).Scan(
		&created.ID, &created.GrandPrix, &created.Circuit, &created.Country,
		&created.RaceDate, &created.ImageURL, &created.CreatedAt, &created.UpdatedAt,
	)
	return &created, err
}

func (r *RaceRepo) Update(id int, race *models.Race) (*models.Race, error) {
	query := `
		UPDATE races
		SET grand_prix=$1, circuit=$2, country=$3, race_date=$4, image_url=$5, updated_at=NOW()
		WHERE id=$6
		RETURNING id, grand_prix, circuit, country, race_date, COALESCE(image_url, ''), created_at, updated_at
	`
	var updated models.Race
	err := r.db.QueryRow(
		query, race.GrandPrix, race.Circuit, race.Country, race.RaceDate, race.ImageURL, id,
	).Scan(
		&updated.ID, &updated.GrandPrix, &updated.Circuit, &updated.Country,
		&updated.RaceDate, &updated.ImageURL, &updated.CreatedAt, &updated.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &updated, err
}

func (r *RaceRepo) Delete(id int) (bool, error) {
	result, err := r.db.Exec("DELETE FROM races WHERE id=$1", id)
	if err != nil {
		return false, err
	}
	rows, _ := result.RowsAffected()
	return rows > 0, nil
}
