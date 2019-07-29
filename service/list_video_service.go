package service



import (
	"golang-crud/model"
	"golang-crud/serializer"


)



// ShowVideoService 视频列表的服务

type ListVideoService struct {
}



// Show 视频列表

func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error

	if err != nil {
		return serializer.Response{
			Status: 502,
			Msg:    "查询异常",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}

}