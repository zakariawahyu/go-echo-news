package dependency

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

type DatabaseRepository struct {
	postgre *DatabasePostgreSQL
	mongo   *DatabaseMongoDB
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	database := &Database{
		Name: "MongoDB",
	}

	return (*DatabaseMongoDB)(database)
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	database := &Database{
		Name: "PostgreSQL",
	}
	return (*DatabasePostgreSQL)(database)
}

func NewDatabaseRepository(mongo *DatabaseMongoDB, postgre *DatabasePostgreSQL) *DatabaseRepository {
	return &DatabaseRepository{
		postgre: postgre,
		mongo:   mongo,
	}
}
