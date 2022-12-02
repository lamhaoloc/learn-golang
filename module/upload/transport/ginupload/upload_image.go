package ginupload

import (
	"awesomeProject/common"
	"awesomeProject/component/appctx"
	uploadService "awesomeProject/module/upload/service"
	uploadStorage "awesomeProject/module/upload/storage"
	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := ctx.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadStorage.NewSQLStore(appCtx.GetMainDBConnection())
		service := uploadService.NewUploadService(appCtx.GetUploadProvider(), imgStore)
		img, err := service.Upload(ctx.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		ctx.JSON(200, common.SimpleSuccessResponse(img))
	}
}
