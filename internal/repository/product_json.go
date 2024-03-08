package repository

import (
	"GoWeb/proyect/internal"
	"encoding/json"
	"os"
)

func NewStorageJsonFile(file string) *StorageJsonFile {
	return &StorageJsonFile{
		file: file,
	}
}

type StorageJsonFile struct {
	file string
}

type ProductJson struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func Read(s *StorageJsonFile) (products []*internal.Product, err error) {
	file, err := os.Open(s.file)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&products); err != nil {
		panic(err)
	}
	return products, nil
}
