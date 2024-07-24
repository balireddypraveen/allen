package configs

const (
	VKEYS_HOST_IP   = "host.ip"
	VKEYS_HOST_PORT = "host.port"
	VKEYS_HOST_TYPE = "HOST_TYPE"

	VKEYS_DATABASE_POSTGRES_SOURCE_HOST              = "database.postgres.source.host"
	VKEYS_DATABASE_POSTGRES_SOURCE_PORT              = "database.postgres.source.port"
	VKEYS_DATABASE_POSTGRES_SOURCE_DB_NAME           = "database.postgres.source.db_name"
	VKEYS_DATABASE_POSTGRES_SOURCE_PASSWORD          = "database.postgres.source.password"
	VKEYS_DATABASE_POSTGRES_SOURCE_USER              = "database.postgres.source.user"
	VKEYS_DATABASE_POSTGRES_SOURCE_MAX_IDLE_CONN     = "database.postgres.max_idle_connections"
	VKEYS_DATABASE_POSTGRES_SOURCE_MAX_OPEN_CONN     = "database.postgres.max_open_connections"
	VKEYS_DATABASE_POSTGRES_SOURCE_MAX_CONN_LIFETIME = "database.postgres.max_connection_lifetime"

	VKEYS_CORS_ORIGINS              = "cors.origins"
	VKEYS_SERVER_ALLOW_HEADERS      = "server.header_allows"
	VKEYS_ALLOW_METHODS             = "server.method_allows"
	VKEYS_EXPOSED_HEADERS           = "server.exposed_headers"
	VKEYS_IDLE_TIMEOUT_SERVER       = "server.idle_timeout"
	VKEYS_READ_WRITE_TIMEOUT_SERVER = "server.read_write_timeout"

	VKEYS_NEWRELIC_LICENSE  = "new_relic.license"
	VKEYS_NEWRELIC_ENABLED  = "new_relic.enabled"
	VKEYS_NEWRELIC_APP_NAME = "new_relic.app_name"

	VKEYS_REDIS_CLUSTERS_HOST_URL             = "redis_cluster.host_cluster_url"

	VKEYS_LOGGING_FORMAT = "logging.format"
	VKEYS_LOGGING_LEVEL  = "logging.level"


)

var (
	PreProdHost = "preprod"
	ProdHost    = "prod"
	Local       = "local"

	ProdHosts  = []interface{}{ProdHost, PreProdHost}
	LocalHosts = []interface{}{Local}
)
