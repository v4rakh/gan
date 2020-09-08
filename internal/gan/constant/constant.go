package constant

const (
	// general
	AppName    = "gan"
	AppVersion = "1.0.0-SNAPSHOT"

	// http
	AppNameHeader    = "X-App-Name"
	AppVersionHeader = "X-App-AppVersion"

	// defaults
	ServerListenDefault = ""
	ServerPortDefault   = "8080"
	SqliteDbNameDefault = "gan.db"

	CorsAllowOriginDefault  = "*"
	CorsAllowMethodsDefault = "GET, POST, PUT, PATCH, DELETE, OPTIONS"
	CorsAllowHeadersDefault = "Authorization, Content-Type"

	// environment
	EnvDbFile       = "DB_FILE"
	EnvServerPort   = "SERVER_PORT"
	EnvServerListen = "SERVER_LISTEN"

	EnvAdminUser     = "ADMIN_USER"
	EnvAdminPassword = "ADMIN_PASSWORD"

	EnvCorsAllowOrigin  = "CORS_ALLOW_ORIGIN"
	EnvCorsAllowMethods = "CORS_ALLOW_METHODS"
	EnvCorsAllowHeaders = "CORS_ALLOW_HEADERS"
)
