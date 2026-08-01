package config

import (
	"github.com/spf13/viper"
	"unsia.ac.id/akademic_be/internal/dto"
)

func NewModel(viper *viper.Viper) *Config {
	dto.APP_DEBUG = viper.GetBool("APP_DEBUG")
	return &Config{
		Server: Server{
			Host:    viper.GetString("SERVER_HOST"),
			Port:    viper.GetString("SERVER_PORT"),
			Prefork: viper.GetBool("SERVER_PREFORK"),
		},
		Database: Database{
			Host:         viper.GetString("DB_HOST"),
			Port:         viper.GetString("DB_PORT"),
			User:         viper.GetString("DB_USER"),
			Pass:         viper.GetString("DB_PASS"),
			Name:         viper.GetString("DB_NAME"),
			PoolIdle:     viper.GetInt("DB_POOL_IDLE"),
			PoolMax:      viper.GetInt("DB_POOL_MAX"),
			PoolLifetime: viper.GetInt("DB_POOL_LIFETIME"),
		},
		Logrus: Logger{
			Level: viper.GetInt("LOGGER_LEVEL"),
		},
		JWT: JWT{
			PublicKeyPath: viper.GetString("JWT_PUBLIC_KEY_PATH"),
		},
		AppConfig: AppConfig{
			AppName: viper.GetString("APP_NAME"),
		},
		Minio: Minio{
			Endpoint:        viper.GetString("MINIO_ENDPOINT_URL"),
			Port:            viper.GetInt("MINIO_ENDPOINT_PORT"),
			AccessKeyID:     viper.GetString("MINIO_ACCESS_KEY_ID"),
			SecretAccessKey: viper.GetString("MINIO_SECRET_ACCESS_KEY"),
			UseSSL:          viper.GetBool("MINIO_USE_SSL"),
			BucketName:      viper.GetString("BUCKET_NAME"),
			MaxSizeFile:     viper.GetInt("MAX_SIZE_FILE"),
		},
		Redis: Redis{
			Host:     viper.GetString("REDIS_HOST"),
			Password: viper.GetString("REDIS_PASS"),
			DB:       viper.GetInt("REDIS_DB"),
		},
		Sentry: Sentry{
			DSN: viper.GetString("SENTRY_DSN"),
		},
		AppKeyMetrics: viper.GetString("APP_KEY_METRICS"),
		// Stream redis
		RedisStream: RedisStream{
			ApiKey: viper.GetString("INTERNAL_API_KEY"),
			LogAuditTrailStream: LogAuditTrailStream{
				RedisStreamLogAuditTrail: viper.GetString("REDIS_STREAM_LOG_AUDIT_TRAIL"),
			},
		},
		// internal Service
		InternalService: InternalService{
			ApiKey: viper.GetString("INTERNAL_API_KEY"),
			SsoApp: SsoApp{
				SsoBaseURL: viper.GetString("INTERNAL_BASE_URL_SSO"),
				UploadService: UploadService{
					UploadServiceEndpoint: viper.GetString("INTERNAL_BASE_URL_SSO") + viper.GetString("INTERNAL_UPLOAD_SERVICE_ENDPOINT_PATH"),
					BucketName:            viper.GetString("INTERNAL_BUCKET_NAME"),
					MaxSizeFile:           viper.GetInt("INTERNAL_MAX_SIZE_FILE"),
				},
				PermissionAction: PermissionAction{
					PermissionActionEndpoint: viper.GetString("INTERNAL_BASE_URL_SSO") + viper.GetString("INTERNAL_PERMISSION_ACTION_ENDPOINT_PATH"),
				},
			},
			SdmApp: SdmApp{
				SdmBaseURL: viper.GetString("INTERNAL_BASE_URL_SDM"),
				GeneralInformation: GeneralInformation{
					GeneralInformationEndpoint: viper.GetString("INTERNAL_BASE_URL_SDM") + viper.GetString("INTERNAL_GENERAL_INFORMATION_ENDPOINT_PATH"),
				},
			},
			DatareferensiApp: DatareferensiApp{
				DatareferensiBaseURL: viper.GetString("INTERNAL_BASE_URL_DATAREFERENSI"),
				AcademicPeriod: AcademicPeriod{
					AcademicPeriodEndpoint: viper.GetString("INTERNAL_BASE_URL_DATAREFERENSI") + viper.GetString("INTERNAL_ACADEMIC_PERIOD_ENDPOINT_PATH"),
				},
			},
		},
	}
}
