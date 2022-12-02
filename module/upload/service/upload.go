package uploadService

import (
	"awesomeProject/common"
	"awesomeProject/component/uploadprovider"
	uploadModel "awesomeProject/module/upload/model"
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	//CreateImage(context context.Context, data *common.Image) error
}

type uploadService struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadService(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *uploadService {
	return &uploadService{provider: provider, imgStore: imgStore}
}
func (service *uploadService) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadModel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg

	img, err := service.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadModel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	//if err := service.imgStore.CreateImage(ctx, img); err != nil {
	//	item := fmt.Sprintf("%s/%s", folder, fileName)
	//	return nil, uploadModel.ErrCannotSaveFile(err)
	//}

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
