package repository

import "github.com/elkmos/my_go/models"

type IdfRepository interface {
	Save(*models.Idf) error
	Update(string, *models.Idf) error
	Delete(string) error
	FindByID(string) (*models.Idf, error)
	FindAll() (models.Idfs, error)
}
