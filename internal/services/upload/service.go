package upload

import (
	"github.com/BarniBl/SolAr_2020/internal/models"
)

type Service interface {
	File(request models.WriteFile) (response models.File, err error)
	Photo(request models.WritePhoto) (response models.Photo, err error)
}

type service struct {
	uploadStorage uploadStorage
}

func NewService(uploadStorage uploadStorage) Service {
	return &service{
		uploadStorage: uploadStorage,
	}
}

func (s *service) File(request models.WriteFile) (response models.File, err error) {
	response, err = s.uploadStorage.SaveFile(request)
	if err != nil {
		return
	}

	response.ID, err = s.uploadStorage.InsertFile(response)

	return
}

func (s *service) Photo(request models.WritePhoto) (response models.Photo, err error) {
	response, err = s.uploadStorage.SavePhoto(request)
	if err != nil {
		return
	}

	response.ID, err = s.uploadStorage.InsertPhoto(response)

	return
}