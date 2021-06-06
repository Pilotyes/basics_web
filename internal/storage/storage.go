package storage

//Storage ...
type Storage interface {
	Posts() PostRepository
}
