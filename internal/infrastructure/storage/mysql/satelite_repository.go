package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
	"github.com/huandu/go-sqlbuilder"
)

type SateliteRepository struct {
	db *sql.DB
}

func NewSateliteRepository(db *sql.DB) *SateliteRepository {
	return &SateliteRepository{
		db: db,
	}
}

func (r *SateliteRepository) Save(satelite domain.Satelite) error {
	sqlStruct := sqlbuilder.NewStruct(new(sqlSatelite))
	query, args := sqlStruct.InsertInto(sqlSateliteTable, sqlSatelite{
		ID:   satelite.ID().Value(),
		Name: satelite.Name().Value(),
		X:    satelite.Position().X().Value(),
		Y:    satelite.Position().Y().Value(),
	}).Build()

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

func (r *SateliteRepository) FindAll() ([]domain.Satelite, error) {
	sqlStruct := sqlbuilder.NewStruct(new(sqlSatelite))
	query, args := sqlStruct.SelectFrom("satellites").Build()

	rows, err := r.db.Query(query, args...)
	if err != nil {
		fmt.Print(err)
		return nil, errors.New("error trying to find satellites on database")
	}

	return r.parseAggregate(rows)
}

func (r *SateliteRepository) FindByName(name string) (domain.Satelite, error) {
	sqlStruct := sqlbuilder.NewStruct(new(sqlSatelite))
	sb := sqlStruct.SelectFrom("satellites")
	sb.Where(sb.Equal("name", name))
	query, args := sb.Build()

	rows, err := r.db.Query(query, args...)

	if err != nil {
		fmt.Print(err)
		return domain.Satelite{}, errors.New("error trying to find satellites on database")
	}

	values, err := r.parseAggregate(rows)

	if err != nil {
		return domain.Satelite{}, errors.New("error trying to find satellites on database")
	}

	if values == nil {
		return domain.Satelite{}, errors.New("error trying to find satelite on database")
	}

	return values[0], nil
}

func (r *SateliteRepository) parseAggregate(rows *sql.Rows) ([]domain.Satelite, error) {
	var values []domain.Satelite
	for rows.Next() {
		var id string
		var name string
		var x float64
		var y float64

		err := rows.Scan(&id, &name, &x, &y)

		if err != nil {
			return nil, err
		}

		response, err := domain.NewSatelite(id, name, x, y)

		if err != nil {
			return nil, err
		}

		values = append(values, response)
	}

	return values, nil
}
