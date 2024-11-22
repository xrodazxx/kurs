import ()

type Authorization interface {
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
	return &Repository{}
}