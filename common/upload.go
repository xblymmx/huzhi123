package common

import (
	"strings"
	"errors"
	"mime"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/xblymmx/huzhi123/config"
	"unicode/utf8"
	"github.com/xblymmx/huzhi123/utils"
	"github.com/xblymmx/huzhi123/model"
	"net/http"
	"github.com/xblymmx/huzhi123/constant"
	"github.com/satori/go.uuid"
)

type ImageUploadInfo struct {
	UploadDir string
	UploadFilePath string
	Filename string
	UUIDName string
	ImgURL string
}

func GenerateImgUploadInfo(ext string) ImageUploadInfo {
	sep := string(os.PathSeparator)
	uploadImgDir := config.ServerConfig.UploadImgDir
	lastChar, _ := utf8.DecodeLastRuneInString(uploadImgDir)
	ymStr := utils.GetTodayYM(sep)

	var uploadDir string
	if string(lastChar) != sep {
		uploadDir = uploadImgDir + sep + ymStr
	} else {
		uploadDir = uploadImgDir + ymStr
	}

	uuidName, err := uuid.NewV4()
	uuidNameStr := uuidName.String()
	if err != nil {
		fmt.Println("uuid generate error")
	}
	filename := uuidNameStr + ext
	uploadFilePath := uploadDir + sep + filename
	imgURL := strings.Join([]string{
		"https://" + config.ServerConfig.ImgHost + config.ServerConfig.ImgPath,
		ymStr,
		filename,
	}, "/")

	return ImageUploadInfo{
		UploadDir: uploadDir,
		UploadFilePath: uploadFilePath,
		Filename: filename,
		UUIDName: uuidNameStr,
		ImgURL: imgURL,
	}
}

func upload(ctx *gin.Context) (map[string]interface{}, error) {
	file, err := ctx.FormFile("upload")

	if err != nil {
		return nil, err
	}

	filename := file.Filename
	index := strings.Index(filename, ".")

	if index < 0 {
		return nil, errors.New(constant.Msg.InvalidFileName)
	}

	ext := filename[index:]
	if len(ext) < 1 {
		return nil, errors.New(constant.Msg.InvalidFileExtensionName)
	}

	mimeType := mime.TypeByExtension(ext)
	fmt.Printf("filename %s extension name %s mime type %s\n", filename, ext, mimeType)

	if mimeType == "" && ext == "jpeg" {
		mimeType = "image/jpeg"
	}

	if mimeType == "" {
		return nil, errors.New(constant.Msg.InvalidMimeType)
	}

	imageUploadInfo := GenerateImgUploadInfo(ext)
	fmt.Println("upload dir:", imageUploadInfo.UploadDir)

	if err := os.MkdirAll(imageUploadInfo.UploadDir, 0777); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := ctx.SaveUploadedFile(file, imageUploadInfo.UploadFilePath); err != nil {
		fmt.Println(err)
		return nil, err
	}

	image := model.Image{
		Title: imageUploadInfo.Filename,
		OriginalTitle: filename,
		URL: imageUploadInfo.ImgURL,
		Width: 0,
		Height: 0,
		Mime: mimeType,
	}

	if err := model.DB.Create(&image).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return map[string]interface{}{
		"id": image.ID,
		"url": imageUploadInfo.ImgURL,
		"title": imageUploadInfo.Filename,
		"originalTitle": filename,
		"mimeType": mimeType,
	}, nil
}

func UploadHandler(ctx *gin.Context) {
	data, err := upload(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.Code.ERROR,
			"msg": err.Error(),
			data: gin.H{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.Code.SUCCESS,
		"msg": constant.Msg.SUCCESS,
		"data": data,
	})
}
