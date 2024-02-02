package service

import "github.com/lordofthemind/goGinHTTPFramework/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	Videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

// Save implements VideoService.
func (service *videoService) Save(video entity.Video) entity.Video {
	service.Videos = append(service.Videos, video)
	return video
}

// FindAll implements VideoService.
func (service *videoService) FindAll() []entity.Video {
	return service.Videos
}
