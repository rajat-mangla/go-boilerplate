package config

import (
	"time"
)

type Config struct {
	AppName string        `yaml:"APP_NAME" env:"APP_NAME"`
	API     HTTPAPIConfig `yaml:"API" env:",prefix=API_"`

	// File paths for translation and configuration files
	//TranslationFileBasePath string `yaml:"TRANSLATION_FILE_BASE_PATH" env:"TRANSLATION_FILE_BASE_PATH"`
	ConfigFilePath string `yaml:"CONFIG_FILE_PATH" env:"CONFIG_FILE_PATH"`

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

	//cfg.TranslationFileBasePath = "i18n"
	cfg.ConfigFilePath = ""

	cfg.LogLevel = "error"
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

type ServiceTimeout struct {
	SampleService time.Duration `yaml:"SAMPLE_SERVICE" env:"SAMPLE_SERVICE"`
}

type Swagger struct {
	Enabled bool   `yaml:"ENABLED" env:"ENABLED"`
	Path    string `yaml:"PATH" env:"PATH"`
}
