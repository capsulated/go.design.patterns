package storage

import "log"

type Storager interface {
	Store() (d *data)
	Delete() (d *data)
}

type Storage struct{}

func NewStorage() *Storage {
	return new(Storage)
}

func (*Storage) Store() {
	log.Println("⏳ Saving...")
	// ...
}

func (*Storage) Delete() {
	log.Println("⏳ Delete...")
	// ...
}
