/**
* Created by GoLand.
* User: 平凡
* Date: 2024/1/7
* Time: 16:03
* Remark: 上传单个图片
 */
package images_api

import (
	"admin/global"
	"admin/models/res"
	"admin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"path"
	"strings"
)

var (
	// WhiteImageList 图片上传的白名单
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// ImageUploadView 上传多个图片，返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	// 上传多个图片
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}

	// 判断路径是否存在
	//pathList,ok := path.Split(global.Config.Upload.Path,"/")
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 递归创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	// 不存在就创建
	var resList []FileUploadResponse

	for _, file := range fileList {
		fileName := file.Filename
		nameList := strings.Split(fileName, ".")
		suffix := strings.ToLower(nameList[len(nameList)-1])

		if !utils.InList(suffix, WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}

		filePath := path.Join(basePath, file.Filename)
		// 判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过设定大小，当前大小为:%.2fMB，设定大小为: %dMb", size, global.Config.Upload.Size),
			})
			continue
		}

		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})
	}
	res.OkWithData(resList, c)
}
