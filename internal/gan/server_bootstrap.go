package gan

import (
	"github.com/adrg/xdg"
	"github.com/v4rakh/gan/internal/gan/constant"
	"github.com/v4rakh/gan/internal/gan/domain/announcement"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

type env struct {
	db *gorm.DB
}

func boostrapConfiguration() {
	if os.Getenv(constant.EnvAdminUser) == "" {
		log.Fatalf("Not all required ENV variables given. Please set %s\n", constant.EnvAdminUser)
		return
	}

	if os.Getenv(constant.EnvAdminPassword) == "" {
		log.Fatalf("Not all required ENV variables given. Please set %s\n", constant.EnvAdminPassword)
		return
	}

	setConfigurationDefaults(constant.EnvServerPort, constant.ServerPortDefault)
	setConfigurationDefaults(constant.EnvServerListen, constant.ServerListenDefault)
	setConfigurationDefaults(constant.EnvCorsAllowOrigin, constant.CorsAllowOriginDefault)
	setConfigurationDefaults(constant.EnvCorsAllowMethods, constant.CorsAllowMethodsDefault)
	setConfigurationDefaults(constant.EnvCorsAllowHeaders, constant.CorsAllowHeadersDefault)
}

func setConfigurationDefaults(key string, defaultValue string) {
	var err error
	if os.Getenv(key) == "" {
		err = os.Setenv(key, defaultValue)

		if err != nil {
			log.Fatalf("Could not set default value for ENV variable '%s'", key)
		}
	}
}

func bootstrapDatabase() *env {
	if os.Getenv(constant.EnvDbFile) == "" {
		defaultDbFile, err := xdg.DataFile(constant.AppName + "/" + constant.SqliteDbNameDefault)

		if err != nil {
			log.Fatalf("Database file '%s' could not be created. Reason: %v", defaultDbFile, err)
		}

		setConfigurationDefaults(constant.EnvDbFile, defaultDbFile)
	}

	dbFile := os.Getenv(constant.EnvDbFile)
	log.Printf("Using database file '%s'\n", dbFile)

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not setup database: %v\n", err)
	}

	env := &env{db: db}
	err = env.db.AutoMigrate(&announcement.Announcement{})
	if err != nil {
		log.Fatalf("Could not migrate database schema: %v\n", err)
	}

	return env
}
