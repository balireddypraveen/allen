package constants

const (
	ConfigPath        = "./configs"
	YML               = "yml"
	ConnectionAddress = "0.0.0.0:8080"

	ContentTypeXxxFormUrlencoded   = "application/x-www-form-urlencoded"
	RequestFormatJson              = "JSON"
	RequestFormatXml               = "XML"
	RequestFormatPipeDelimited     = "PIPE_DELIMITED"
	RequestHeaderXAuthServiceToken = "X-Auth-Service-Token"
)

const (
	EmptyString = ""
	COLLON      = ":"
	HYPHEN      = "-"
	QUESTION    = "?"
)

// logging
const (
	Zap             = "zap"
	CommonFieldsKey = "commonFields"
	LogLevelDebug   = "debug"
	LogLevelInfo    = "info"
	LogLevelWarn    = "warn"
)

const (
	RequestPath   = "request_path"
	RequestMethod = "request_method"
)

const (
	ContextLogger = "logger"
	ContextNrTxn  = "nr_txn"
)

type Header string

const (
	HeaderAccessToken   Header = "X-Access-Token"
	HeaderSecureToken   Header = "X-Secure-Token"
	HeaderForwardedFor  Header = "X-Forwarded-For"
	HeaderTraceReqId    Header = "X-Request-Id"
	HeaderTraceSpanID   Header = "X-Span-Request-Id"
	HeaderTraceAmznID   Header = "X-Amzn-Trace-Id"
	HeaderNrTraceId     Header = "nr-Trace-ID"
	HeaderDebugMode     Header = "X-Debug-Mode"
	MessageId           Header = "stream-message-id"
	StreamName          Header = "stream-name"
	UserCustodyWalletID Header = "ctx_user_custody_wallet_id"
)

const (
	HeaderContentType     = "Content-Type"
	HeaderContentTypeJson = "application/json"
	HeaderXAuthToken      = "X-Auth-Token"
	HeaderXRequestId      = "X-Request-Id"
)

var (
	PreProdHost = "preprod"
	ProdHost    = "prod"
	Local       = "local"

	ProdHosts  = []interface{}{ProdHost, PreProdHost}
	LocalHosts = []interface{}{Local}
)
