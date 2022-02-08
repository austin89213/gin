package service

import "golang-gin/entity"

type VideoService interface {
	Save(video entity.Video)
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{
		videos: []entity.Video{},
	}
}

func (service *videoService) Save(video entity.Video) {
	service.videos = append(service.videos, video)
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
