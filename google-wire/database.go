package google_wire

type Database struct {
	Name string
}

type DBPostgreSQL Database
type DBMongoDB Database

func NewDatabasePostgreSQL() *DBPostgreSQL {
	return (*DBPostgreSQL)(&Database{Name: "PostgreSQL"})
}

func NewDatabaseMongoDB() *DBMongoDB {
	return (*DBMongoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DBPostgreSQL
	DatabaseMongoDB    *DBMongoDB
}

func NewDatabaseRepository(databasePostgreSQL *DBPostgreSQL, databaseMongoDB *DBMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: databasePostgreSQL,
		DatabaseMongoDB:    databaseMongoDB,
	}
}
