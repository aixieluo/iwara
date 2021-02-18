package repositories

import (
	"iwara/models"
)

type videoRepository struct{}

func (r *videoRepository) NewIndexRepository() *videoRepository {
	return &videoRepository{}
}

func (r *videoRepository) Find(id int) *models.Video {

}


