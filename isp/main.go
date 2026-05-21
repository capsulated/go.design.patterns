package isp

import "log"

type Storager interface {
	Store()
	Delete()
}

type FileStorer interface {
	FileStore()
}

type DbStorer interface {
	DbStore()
}

type UniversalStorage interface {
	Storager
	FileStorer
	DbStorer
}

type FileStorage struct{}

type DbStorage struct{}


func NewFileStorage() *FileStorage {
	return new(Storage)
}

func (*FileStorage) Store() {
	log.Println("⏳ Saving...")
	// ...
}

func (*FileStorage) Delete() {
	log.Println("⏳ Delete...")
	// ...
}

func (*FileStorage) FileStore() {
	log.Println("⏳ Delete...")
	// ...
}





func NewDbStorage() *DbStorage {
	return new(Storage)
}

func (*DbStorage) Store() {
	log.Println("⏳ Saving...")
	// ...
}

func (*DbStorage) Delete() {
	log.Println("⏳ Delete...")
	// ...
}

func (*DbStorage) DbStore() {
	log.Println("⏳ Delete...")
	// ...
}

func main() {

}
