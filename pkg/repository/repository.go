import ()

type Authorization interface {
	CreateUser(user kurs.User) (int, error)
}
type DeviceIot interface {
}
type DeviceData interface {
}
type Repository struct {
	Authorization
	DeviceData
	DeviceIot
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db)
	}
}