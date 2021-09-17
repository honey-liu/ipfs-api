package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ipfs_api/driver"
	"ipfs_api/utils"
	"net/http"
	"path/filepath"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	basePath := "./upload/"
	filename := basePath + filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	data,err := utils.ReadAll(filename)
	if err != nil {
		return
	}
	hash ,err := driver.UploadIPFS(data)
	if err != nil {
		c.String(http.StatusInternalServerError,fmt.Sprintf("'%s' uploaded filed!,err:%s", file.Filename,err))

		return
	}
	c.JSON(http.StatusOK, UploadFileResponse{Code: 200, Status: fmt.Sprintf("'%s' uploaded success!", file.Filename),FileHash: hash})
	return
}