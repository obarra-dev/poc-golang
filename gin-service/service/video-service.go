package service

import "gin-service/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{
		videos: []entity.Video{{
			Title:       "Cool title",
			Description: "Some movie",
			URL:         "https://www.youtube.com/watch?v=vjd3XmzXZ5s",
			Actors:      40,
			Author:      entity.Person{
				Name:  "Omar Barra",
				Age:   30,
				Email: "der@gmail.com",
			},
		}},
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
