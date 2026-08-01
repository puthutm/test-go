package config

type Config struct {
	Database        Database
	Server          Server
	Logrus          Logger
	JWT             JWT
	AppConfig       AppConfig
	Minio           Minio
	Redis           Redis
	Sentry          Sentry
	AppKeyMetrics   string
	InternalService InternalService
	RedisStream     RedisStream
}

type Database struct {
	Host         string
	Port         string
	User         string
	Pass         string
	Name         string
	PoolIdle     int
	PoolMax      int
	PoolLifetime int
}

type Redis struct {
	Host     string
	Password string
	DB       int
}

type Server struct {
	Host    string
	Port    string
	Prefork bool
}

type Logger struct {
	Level int
}

type JWT struct {
	PublicKeyPath string
}

type AppConfig struct {
	AppName string
}

type Minio struct {
	Endpoint        string
	Port            int
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
	MaxSizeFile     int
}

type InternalService struct {
	ApiKey           string
	SsoApp           SsoApp
	SdmApp           SdmApp
	DatareferensiApp DatareferensiApp
}

type SsoApp struct {
	SsoBaseURL       string
	UploadService    UploadService
	PermissionAction PermissionAction
}

type UploadService struct {
	BucketName            string
	MaxSizeFile           int
	UploadServiceEndpoint string
}

type PermissionAction struct {
	PermissionActionEndpoint string
}

type SdmApp struct {
	SdmBaseURL         string
	GeneralInformation GeneralInformation
}

type GeneralInformation struct {
	GeneralInformationEndpoint string
}

type DatareferensiApp struct {
	DatareferensiBaseURL string
	AcademicPeriod       AcademicPeriod
}

type AcademicPeriod struct {
	AcademicPeriodEndpoint string
}

type Sentry struct {
	DSN string
}

type RedisStream struct {
	ApiKey              string
	LogAuditTrailStream LogAuditTrailStream
}

type LogAuditTrailStream struct {
	RedisStreamLogAuditTrail string
}
