package appctx

import (
	"awesomeProject/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetUploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, provider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, uploadProvider: provider}
}

func (ctx appCtx) GetMainDBConnection() *gorm.DB {
	return appCtx{}.db
}
func (ctx appCtx) GetUploadProvider() uploadprovider.UploadProvider {
	return appCtx{}.uploadProvider
}
