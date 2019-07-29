package service

import (
	"golang-crud/model"
	"golang-crud/serializer"
)

// CreateVideoService 视频上传服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Info string `form:"info" json:"info" binding:"required,min=0,max=300"`
}

// Create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info: service.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil{
		return serializer.Response{
			Status: 500,
			Msg: "视频上传失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
