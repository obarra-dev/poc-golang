package service

import (
	"gin-service/entity"
	"strings"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindWithFilter(title string) []entity.Video
	FindOne(GUID string) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{
		videos: []entity.Video{{
			GUID:        "abc234",
			Title:       "Cool title",
			Description: "Some movie",
			URL:         "https://www.youtube.com/watch?v=vjd3XmzXZ5s",
			Actors:      40,
			Author: entity.Person{
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

func (service *videoService) FindWithFilter(title string) []entity.Video {
	if title == "" {
		return service.FindAll()
	}

	var filtered []entity.Video
	for _, video := range service.videos {
		if strings.ContainsAny(video.Title, title) {
			filtered = append(filtered, video)
		}
	}
	return filtered
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}

func (service *videoService) FindOne(GUID string) entity.Video {
	for _, video := range service.videos {
		if video.GUID == GUID {
			return video
		}
	}
	return entity.Video{}
}
