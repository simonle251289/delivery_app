package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	UploadProvider()
	GetPubSub()
}

type appContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db}
}

func (ctx *appContext) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appContext) SecretKey() string {
	//TODO implement me
	panic("implement me")
}

func (ctx *appContext) UploadProvider() {
	//TODO implement me
	panic("implement me")
}

func (ctx *appContext) GetPubSub() {
	//TODO implement me
	panic("implement me")
}
