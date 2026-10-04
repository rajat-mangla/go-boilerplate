package config

import (
	"time"
)

type Config struct {
	AppName string        `yaml:"APP_NAME" env:"APP_NAME"`
	API     HTTPAPIConfig `yaml:"API" env:",prefix=API_"`

	// Database configuration for the application
	DatabaseEnabled bool           `yaml:"DB_ENABLED" env:"DB_ENABLED"`
	Database        DatabaseConfig `yaml:"DB" env:",prefix=DB_"`

	// File paths for translation and configuration files
	//TranslationFileBasePath string `yaml:"TRANSLATION_FILE_BASE_PATH" env:"TRANSLATION_FILE_BASE_PATH"`

	// Logging configuration
	LogLevel  string `yaml:"LOG_LEVEL" env:"LOG_LEVEL"`
	LogOutput string `yaml:"LOG_OUTPUT" env:"LOG_OUTPUT"` // Should be one of "console" | "stdout"
	LogFormat string `yaml:"LOG_FORMAT" env:"LOG_FORMAT"`

	// Service timeout configuration for services
	ServiceTimeout ServiceTimeout `yaml:"SERVICE_TIMEOUT" env:",prefix=SERVICE_TIMEOUT_"`

	// Swagger configuration for API documentation
	Swagger Swagger `yaml:"SWAGGER" env:",prefix=SWAGGER_"`

	// HTTP client configuration for client services
	SampleClientConfig HTTPConfig `yaml:"SAMPLE" env:",prefix=SAMPLE_CLIENT_"`
}

func (cfg *Config) SetDefaults() {
	cfg.AppName = "boilerplate-service"

	cfg.API.ListenAddr = ":8080"

	cfg.Database = DatabaseConfig{
		Host:            "localhost",
		Port:            5432,
		User:            "postgres",
		Name:            "boilerplate",
		ConnectTimeout:  5 * time.Second,
		MaxOpenConns:    5,
		MaxIdleConns:    1,
		ConnMaxLifetime: 5 * time.Minute,
		ConnMaxIdleTime: 1 * time.Minute,
	}

	//cfg.TranslationFileBasePath = "i18n"

	cfg.LogLevel = "info"
	cfg.LogOutput = "console"
	cfg.LogFormat = "text"

	cfg.ServiceTimeout = ServiceTimeout{SampleService: time.Second * 5}

	cfg.SampleClientConfig = HTTPConfig{
		HTTPTimeout:            1000,
		HystrixTimeout:         1000,
		MaxConcurrentRequests:  100,
		RequestVolumeThreshold: 100,
		SleepWindow:            100,
		ErrorPercentThreshold:  10,
	}
}

func NewConfig(path string) (*Config, error) {
	cfg := &Config{}
	cfg.SetDefaults()

	err := LoadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// HTTPAPIConfig is used for server configuration.
type HTTPAPIConfig struct {
	ListenAddr string `yaml:"LISTEN_ADDR" env:"HTTP_LISTEN_ADDR"`
	DebugMode  bool   `yaml:"DEBUG_MODE" env:"DEBUG_MODE"`
}

type DatabaseConfig struct {
	Host            string        `yaml:"HOST" env:"HOST"`
	Port            int           `yaml:"PORT" env:"PORT"`
	User            string        `yaml:"USER" env:"USER"`
	Password        string        `yaml:"PASSWORD" env:"PASSWORD"`
	Name            string        `yaml:"NAME" env:"NAME"`
	ConnectTimeout  time.Duration `yaml:"CONNECT_TIMEOUT" env:"CONNECT_TIMEOUT"`
	MaxOpenConns    int           `yaml:"MAX_OPEN_CONNS" env:"MAX_OPEN_CONNS"`
	MaxIdleConns    int           `yaml:"MAX_IDLE_CONNS" env:"MAX_IDLE_CONNS"`
	ConnMaxLifetime time.Duration `yaml:"CONN_MAX_LIFETIME" env:"CONN_MAX_LIFETIME"`
	ConnMaxIdleTime time.Duration `yaml:"CONN_MAX_IDLE_TIME" env:"CONN_MAX_IDLE_TIME"`
}

type ServiceTimeout struct {
	SampleService time.Duration `yaml:"SAMPLE_SERVICE" env:"SAMPLE_SERVICE"`
}

type Swagger struct {
	Enabled bool   `yaml:"ENABLED" env:"ENABLED"`
	Path    string `yaml:"PATH" env:"PATH"`
}
