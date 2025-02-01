package di

//import (
//	"github.com/gofiber/fiber/v2"
//	"gorm.io/gorm"
//	"pushpost/internal/services/user_service/transport"
//)
//
//type Handler []interface{}
//
//type Container struct {
//	Server  *fiber.App
//	DB      *gorm.DB
//	Handler []transport.Handler
//}
//
//func NewContainer() (*Container, error) {
//	container := &Container{}
//
//	//userRepo.DB.AutoMigrate(entity.User{})       //fixme make goose migrations
//
//	return container, nil
//}
//
//func (c *Container) RegisterHandler(handlerGroup transport.UserHandler) {
//	c.Handler = handlerGroup
//}
