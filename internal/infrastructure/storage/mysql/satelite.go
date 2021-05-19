package mysql

const (
	sqlSateliteTable = "satellites"
)

type sqlSatelite struct {
	ID   string  `db:"id"`
	Name string  `db:"name"`
	X    float64 `db:"x"`
	Y    float64 `db:"y"`
}
