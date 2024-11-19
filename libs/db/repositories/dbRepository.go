package repositories

type DbRepositoryInterface interface {
}

type dbRepository struct {
}

func NewDbRepository() DbRepositoryInterface {
	return &dbRepository{}
}

func (d *dbRepository) InsertSomthing() {
}

func (d *dbRepository) UpdateSomthing() {
}

func (d *dbRepository) DeleteSomthing() {
}
