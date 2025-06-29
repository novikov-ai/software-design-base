package storage

type Storage interface {
	Save(string)
	Retrieve(int) string
}
