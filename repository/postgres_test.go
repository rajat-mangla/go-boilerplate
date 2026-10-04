package repository

import (
	"net/url"
	"testing"

	"github.com/rajat-mangla/go-boilerplate/config"
)

func TestConnectionURLEscapesCredentials(t *testing.T) {
	cfg := config.DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "app user",
		Password: "p@ss/word",
		Name:     "app",
	}

	connectionURL, err := ConnectionURL(cfg)
	if err != nil {
		t.Fatalf("ConnectionURL() error = %v", err)
	}
	parsed, err := url.Parse(connectionURL)
	if err != nil {
		t.Fatalf("parse connection URL: %v", err)
	}

	user := parsed.User.Username()
	password, _ := parsed.User.Password()
	if user != cfg.User {
		t.Fatalf("connection URL user = %q, want %q", user, cfg.User)
	}
	if password != cfg.Password {
		t.Fatalf("connection URL password = %q, want configured password", password)
	}
	if parsed.Host != "localhost:5432" {
		t.Fatalf("connection URL host = %q, want localhost:5432", parsed.Host)
	}
	if parsed.Path != "/app" {
		t.Fatalf("connection URL path = %q, want /app", parsed.Path)
	}
	if parsed.Query().Get("sslmode") != "disable" {
		t.Fatalf("connection URL sslmode = %q, want disable", parsed.Query().Get("sslmode"))
	}
}

func TestConnectionURLRejectsInvalidPort(t *testing.T) {
	_, err := ConnectionURL(config.DatabaseConfig{
		Host: "localhost",
		Port: 70000,
		User: "postgres",
		Name: "app",
	})
	if err == nil {
		t.Fatal("ConnectionURL() error = nil, want invalid port error")
	}
}
