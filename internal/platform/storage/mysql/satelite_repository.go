package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	core "github.com/franciscoruizar/quasar-fire/internal"
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

func (r *SateliteRepository) Save(satelite core.Satelite) error {
	sqlStruct := sqlbuilder.NewStruct(new(sqlSatelite))
	query, args := sqlStruct.InsertInto(sqlSateliteTable, sqlSatelite{
		ID:   satelite.ID.Value,
		Name: satelite.Name,
		X:    satelite.Position.X,
		Y:    satelite.Position.Y,
	}).Build()

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

func (r *SateliteRepository) FindAll() ([]core.Satelite, error) {
	sqlStruct := sqlbuilder.NewStruct(new(sqlSatelite))
	query, args := sqlStruct.SelectFrom("satellites").Build()

	rows, err := r.db.Query(query, args...)
	defer rows.Close()
	if err != nil {
		return nil, errors.New("error trying to find satellites on database")
	}

	var values []sqlSatelite
	rows.Scan(sqlStruct.Addr(&values)...)

	var satellites []core.Satelite

	for i := 0; i < len(values); i++ {
		value, err := core.NewSatelite(values[i].ID, values[i].Name, values[i].X, values[i].Y)
		if err == nil {
			satellites = append(satellites, value)
		}
	}

	return satellites, nil
}

func (r *SateliteRepository) FindByName(name string) (core.Satelite, error) {
	sqlStruct := sqlbuilder.NewStruct(new(sqlSatelite))
	sb := sqlStruct.SelectFrom("satellites")
	sb.Where(sb.Equal("name", name))

	sql, args := sb.Build()
	rows, err := r.db.Query(sql, args...)
	defer rows.Close()

	if err != nil {
		return core.Satelite{}, errors.New("error trying to find satellites on database")
	}

	var value sqlSatelite
	rows.Scan(sqlStruct.Addr(&value)...)

	response, err := core.NewSatelite(value.ID, value.Name, value.X, value.Y)

	return response, err
}
