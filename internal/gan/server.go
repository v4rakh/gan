package gan

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/v4rakh/gan/internal/gan/api/handler"
	"github.com/v4rakh/gan/internal/gan/api/middleware"
	"github.com/v4rakh/gan/internal/gan/constant"
	"github.com/v4rakh/gan/internal/gan/domain/announcement"
	"os"
)

func StartServer() {
	boostrapConfiguration()
	env := bootstrapDatabase()

	if gin.Mode() == gin.ReleaseMode {
		gin.DisableConsoleColor()
	}

	announcementRepo := announcement.NewRepo(env.db)
	announcementService := announcement.NewService(announcementRepo)

	announcementHandler := handler.NewAnnouncementHandler(announcementService)
	infoHandler := handler.NewInfoHandler()
	userHandler := handler.NewUserHandler()

	router := gin.Default()
	router.Use(middleware.AppName())
	router.Use(middleware.AppVersion())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv(constant.EnvCorsAllowOrigin)},
		AllowMethods:     []string{os.Getenv(constant.EnvCorsAllowMethods)},
		AllowHeaders:     []string{os.Getenv(constant.EnvCorsAllowHeaders)},
		AllowCredentials: true,
	}))

	apiPublicGroup := router.Group("/api")
	apiPublicGroup.GET("/info", infoHandler.ShowInfo)
	apiPublicGroup.GET("/announcements", announcementHandler.ListAnnouncements)
	apiPublicGroup.GET("/announcements/:id", announcementHandler.GetAnnouncement)

	apiAdminGroup := router.Group("/api/admin", gin.BasicAuth(gin.Accounts{
		os.Getenv(constant.EnvAdminUser): os.Getenv(constant.EnvAdminPassword),
	}))
	apiAdminGroup.GET("/login", userHandler.Login)

	apiAdminGroup.POST("/announcements", announcementHandler.CreateAnnouncement)
	apiAdminGroup.PUT("/announcements", announcementHandler.UpdateAnnouncement)
	apiAdminGroup.DELETE("/announcements/:id", announcementHandler.DeleteAnnouncement)

	_ = router.Run(fmt.Sprintf("%s:%s", os.Getenv(constant.EnvServerListen), os.Getenv(constant.EnvServerPort)))
}
