package service

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type BuildingRepository interface {
	CreateBuilding(building Building) (Building, error)
	GetAllBuilding() ([]Building, error)
}

type buildingRepository struct {
	db *sqlx.DB
}

func NewBuildingRepository(db *sqlx.DB) *buildingRepository {
	return &buildingRepository{db: db}
}

func (r *buildingRepository) GetAllBuilding() ([]Building, error) {
	var buildings []Building
	query := `SELECT id, name, city, year, floors_count FROM buildings WHERE 1=1`
	err := r.db.Select(&buildings, query)

	if err != nil {
		log.Printf("Error querying buildings: %v", err)
		return nil, err
	}

	return buildings, nil
}

func (r *buildingRepository) CreateBuilding(building Building) (Building, error) {
	query := `INSERT INTO buildings (name, city, year, floors_count) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(query, building.Name, building.City, building.Year, building.FloorsCount).Scan(&building.ID)
	if err != nil {
		return Building{}, err
	}

	return building, nil
}
