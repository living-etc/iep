package migrations

type Migration struct {
	Id        string
	Statement string
	Args      []any
}
