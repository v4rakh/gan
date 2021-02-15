package gan

import (
	"github.com/adrg/xdg"
	"github.com/v4rakh/gan/internal/gan/constant"
	"github.com/v4rakh/gan/internal/gan/domain/announcement"
	"github.com/v4rakh/gan/internal/gan/domain/subscription"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

type env struct {
	db *gorm.DB
}

func configEnv() *env {
	prepareEnv()

	if os.Getenv(constant.EnvDbFile) == "" {
		defaultDbFile, err := xdg.DataFile(constant.AppName + "/" + constant.SqliteDbNameDefault)

		if err != nil {
			log.Fatalf("Database file '%s' could not be created. Reason: %v", defaultDbFile, err)
		}

		setEnvKeyDefault(constant.EnvDbFile, defaultDbFile)
	}

	dbFile := os.Getenv(constant.EnvDbFile)
	log.Printf("Using database file '%s'\n", dbFile)

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not setup database: %v\n", err)
	}

	env := &env{db: db}
	err = env.db.AutoMigrate(&announcement.Announcement{}, &subscription.Subscription{})
	if err != nil {
		log.Fatalf("Could not migrate database schema: %v\n", err)
	}

	return env
}

func prepareEnv() {
	setEnvKeyDefault(constant.EnvDomain, constant.DomainDefault)
	setEnvKeyDefault(constant.EnvServerPort, constant.ServerPortDefault)
	setEnvKeyDefault(constant.EnvServerListen, constant.ServerListenDefault)
	setEnvKeyDefault(constant.EnvCorsAllowOrigin, constant.CorsAllowOriginDefault)
	setEnvKeyDefault(constant.EnvCorsAllowMethods, constant.CorsAllowMethodsDefault)
	setEnvKeyDefault(constant.EnvCorsAllowHeaders, constant.CorsAllowHeadersDefault)
	setEnvKeyDefault(constant.EnvMailEnabled, constant.MailEnabledDefault)

	vars := []string{constant.EnvAdminUser, constant.EnvAdminPassword}

	if os.Getenv(constant.EnvMailEnabled) == "true" {
		setEnvKeyDefault(constant.EnvMailAuthType, constant.MailAuthTypeDefault)
		setEnvKeyDefault(constant.EnvMailEncryption, constant.MailEncryptionDefault)

		vars = append(vars,
			constant.EnvMailFrom,
			constant.EnvMailHost,
			constant.EnvMailPort,
			constant.EnvMailEncryption,
			constant.EnvMailAuthUser,
			constant.EnvMailAuthPassword,
			constant.EnvMailAuthType)
	}

	for _, s := range vars {
		failIfEnvKeyNotPresent(s)
	}
}

func failIfEnvKeyNotPresent(key string) {
	if os.Getenv(key) == "" {
		log.Fatalf("Not all required ENV variables given. Please set '%s'\n", key)
	}
}

func setEnvKeyDefault(key string, defaultValue string) {
	var err error
	if os.Getenv(key) == "" {
		err = os.Setenv(key, defaultValue)

		if err != nil {
			log.Fatalf("Could not set default value for ENV variable '%s'", key)
		}
	}
}
