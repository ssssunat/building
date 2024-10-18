package service

type Building struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	City        string `db:"city" json:"city"`
	Year        int    `db:"year" json:"year"`
	FloorsCount int    `db:"floors_count" json:"floors_count"`
}
