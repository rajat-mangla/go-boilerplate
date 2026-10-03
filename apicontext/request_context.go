package apicontext

import (
	"context"
	"net/http"
	"uuid"

	"github.com/rajat-mangla/go-boilerplate/constants"
)

const (
	keyRequestContext = "request-context"
)

type RequestContext struct {
	Language    string
	AppID       string
	AppVersion  string
	AppPlatform string
	AppDeviceOS string
	TraceID     string
}

func NewRequestContextFromHTTP(r *http.Request) RequestContext {
	language := r.Header.Get(constants.HeaderAcceptLanguage)
	appID := r.Header.Get(constants.HeaderAppID)
	appVersion := r.Header.Get(constants.HeaderAppVersion)
	appPlatform := r.Header.Get(constants.HeaderAppPlatform)
	requestTraceID := getRequestTraceID(r.Header)
	appDeviceOS := r.Header.Get(constants.HeaderAppDeviceOS)

	return RequestContext{
		Language:    language,
		AppID:       appID,
		AppVersion:  appVersion,
		AppPlatform: appPlatform,
		AppDeviceOS: appDeviceOS,
		TraceID:     requestTraceID,
	}
}

func NewContextWithRequestContext(ctx context.Context, reqCtx RequestContext) context.Context {
	return WithValue(ctx, keyRequestContext, reqCtx)
}

func RequestContextFromContext(ctx context.Context) RequestContext {
	c, ok := Value(ctx, keyRequestContext).(RequestContext)
	if !ok {
		return RequestContext{}
	}
	return c
}

func getRequestTraceID(header http.Header) string {
	if v := header.Get(constants.HeaderXRequestTraceID); v != "" {
		return v
	}

	return uuid.NewV7().String()
}
