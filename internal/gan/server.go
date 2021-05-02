package gan

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/v4rakh/gan/internal/gan/api/handler"
	"github.com/v4rakh/gan/internal/gan/api/middleware"
	"github.com/v4rakh/gan/internal/gan/constant"
	"github.com/v4rakh/gan/internal/gan/domain/announcement"
	"github.com/v4rakh/gan/internal/gan/domain/subscription"
	"github.com/v4rakh/gan/internal/gan/service/i18n"
	"github.com/v4rakh/gan/internal/gan/service/mail"
	"os"
)

func StartServer() {
	env := configEnv()

	if gin.Mode() == gin.ReleaseMode {
		gin.DisableConsoleColor()
	}

	announcementRepo := announcement.NewRepo(env.db)
	subscriptionRepo := subscription.NewRepo(env.db)

	mailService := mail.NewService()
	i18nService := i18n.NewService()
	subscriptionService := subscription.NewService(subscriptionRepo, mailService, i18nService)
	announcementService := announcement.NewService(announcementRepo, subscriptionService)

	announcementHandler := handler.NewAnnouncementHandler(announcementService)
	subscriptionHandler := handler.NewSubscriptionHandler(subscriptionService)
	infoHandler := handler.NewInfoHandler()
	userHandler := handler.NewUserHandler()

	router := gin.Default()
	router.Use(middleware.AppName())
	router.Use(middleware.AppVersion())
	router.Use(middleware.AppErrorRecoveryHandler())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv(constant.EnvCorsAllowOrigin)},
		AllowMethods:     []string{os.Getenv(constant.EnvCorsAllowMethods)},
		AllowHeaders:     []string{os.Getenv(constant.EnvCorsAllowHeaders)},
		AllowCredentials: true,
	}))

	apiPublicGroup := router.Group("/api/v1")
	apiPublicGroup.GET("/info", infoHandler.ShowInfo)

	apiPublicGroup.GET("/announcements", announcementHandler.PaginateAnnouncements)
	apiPublicGroup.GET("/announcements/:id", announcementHandler.GetAnnouncement)

	apiPublicGroup.POST("/subscriptions", subscriptionHandler.CreateSubscription)
	apiPublicGroup.PATCH("/subscriptions", subscriptionHandler.VerifySubscription)
	apiPublicGroup.POST("/subscriptions/rescue", subscriptionHandler.RescueSubscription)
	apiPublicGroup.DELETE("/subscriptions", subscriptionHandler.DeleteSubscription)

	apiAdminGroup := router.Group("/api/v1/admin", gin.BasicAuth(gin.Accounts{
		os.Getenv(constant.EnvAdminUser): os.Getenv(constant.EnvAdminPassword),
	}))

	apiAdminGroup.GET("/login", userHandler.Login)

	apiAdminGroup.POST("/announcements", announcementHandler.CreateAnnouncement)
	apiAdminGroup.PUT("/announcements", announcementHandler.UpdateAnnouncement)
	apiAdminGroup.DELETE("/announcements/:id", announcementHandler.DeleteAnnouncement)

	apiAdminGroup.GET("/subscriptions", subscriptionHandler.PaginateSubscriptions)
	apiAdminGroup.DELETE("/subscriptions", subscriptionHandler.DeleteSubscriptionByAddress)

	_ = router.Run(fmt.Sprintf("%s:%s", os.Getenv(constant.EnvServerListen), os.Getenv(constant.EnvServerPort)))
}
