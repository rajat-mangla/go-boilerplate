package config

import (
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

type HTTPConfig struct {
	Host                   string `yaml:"HOST" env:"HOST"  validate:"nonzero"`
	CandidateHost          string `yaml:"CANDIDATE_HOST" env:"CANDIDATE_HOST"`
	HTTPTimeout            int    `yaml:"HTTP_TIMEOUT" env:"HTTP_TIMEOUT"`
	RetryCount             int    `yaml:"RETRY_COUNT" env:"RETRY_COUNT"`
	Username               string `yaml:"USERNAME" env:"USERNAME"`
	Password               string `yaml:"PASSWORD" env:"PASSWORD"`
	HystrixTimeout         int    `yaml:"HYSTRIX_TIMEOUT" env:"TIMEOUT"`
	MaxConcurrentRequests  int    `yaml:"MAX_CONCURRENT_REQUESTS" env:"MAX_CONCURRENT_REQUESTS"`
	RequestVolumeThreshold int    `yaml:"REQUEST_VOLUME_THRESHOLD" env:"REQUEST_VOLUME_THRESHOLD"`
	SleepWindow            int    `yaml:"SLEEP_WINDOW" env:"SLEEP_WINDOW"`
	ErrorPercentThreshold  int    `yaml:"ERROR_PERCENT_THRESHOLD" env:"ERROR_PERCENT_THRESHOLD"`
	RetryOnCodes           []int
	Transport              http.RoundTripper
}

func (c *HTTPConfig) HystrixConfig() hystrix.CommandConfig {
	return hystrix.CommandConfig{
		Timeout:                c.HystrixTimeout,
		MaxConcurrentRequests:  c.MaxConcurrentRequests,
		SleepWindow:            c.SleepWindow,
		ErrorPercentThreshold:  c.ErrorPercentThreshold,
		RequestVolumeThreshold: c.RequestVolumeThreshold,
	}
}
