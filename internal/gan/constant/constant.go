package constant

const (
	// general
	AppName    = "gan-server"
	AppVersion = "1.0.0"

	// http
	AppNameHeader    = "X-App-Name"
	AppVersionHeader = "X-App-Version"

	// defaults
	DomainDefault = "http://localhost"

	ServerListenDefault = ""
	ServerPortDefault   = "8080"
	SqliteDbNameDefault = "gan.db"

	CorsAllowOriginDefault  = "*"
	CorsAllowMethodsDefault = "GET, POST, PUT, PATCH, DELETE, OPTIONS"
	CorsAllowHeadersDefault = "Authorization, Content-Type"

	MailEnabledDefault    = "true"
	MailAuthTypeDefault   = "PLAIN"
	MailEncryptionDefault = "SSL"

	// environment
	EnvDomain = "DOMAIN"

	EnvDbFile       = "DB_FILE"
	EnvServerPort   = "SERVER_PORT"
	EnvServerListen = "SERVER_LISTEN"

	EnvAdminUser     = "ADMIN_USER"
	EnvAdminPassword = "ADMIN_PASSWORD"

	EnvCorsAllowOrigin  = "CORS_ALLOW_ORIGIN"
	EnvCorsAllowMethods = "CORS_ALLOW_METHODS"
	EnvCorsAllowHeaders = "CORS_ALLOW_HEADERS"

	EnvMailEnabled      = "MAIL_ENABLED"
	EnvMailFrom         = "MAIL_FROM"
	EnvMailHost         = "MAIL_HOST"
	EnvMailPort         = "MAIL_PORT"
	EnvMailEncryption   = "MAIL_ENCRYPTION"
	EnvMailAuthUser     = "MAIL_AUTH_USER"
	EnvMailAuthPassword = "MAIL_AUTH_PASSWORD"
	EnvMailAuthType     = "MAIL_AUTH_TYPE"
)
