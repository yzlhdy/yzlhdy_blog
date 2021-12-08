package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
	"yzlhdy_blog/helper"

	"github.com/gin-gonic/gin"
)

type UploadController interface {
	Upload(ctx *gin.Context)
}
type uploadController struct{}

func NewUploadController() UploadController {
	return &uploadController{}
}

//目录是否存在
func isDirExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
	// if os.IsNotExist(err) {
	// 	return false
	// }
	// return true
}

func (u *uploadController) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err == nil {
		Path := "./public/image/"
		t := time.Now()
		date := t.Format("20060102")
		pathTmp := Path + "/ " + date + "/"
		if isDirExists(pathTmp) {
			fmt.Println("目录存在")
		} else {
			fmt.Println("目录不存在")
			err := os.Mkdir(pathTmp, 0777)
			if err != nil {
				//log.Fatal(err)
				response := helper.BuildErrorResponse(401, "目录创建失败", nil, err)
				ctx.JSON(http.StatusOK, response)
			}
		}
		//文件名
		file_name := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + path.Ext(file.Filename)
		purer := ctx.SaveUploadedFile(file, pathTmp+file_name)
		if purer == nil {
			response := helper.BuildResponse(200, "success", date+"/"+file_name)
			ctx.JSON(http.StatusOK, response)
		} else {
			response := helper.BuildErrorResponse(401, "上传失败", nil, nil)
			ctx.JSON(http.StatusOK, response)
		}
		//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	} else {
		response := helper.BuildErrorResponse(401, "上传失败", nil, nil)
		ctx.JSON(http.StatusOK, response)
	}
}
