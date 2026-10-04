package config

import (
	"path/filepath"
	"testing"
	"time"
)

func TestDatabaseDefaultsAndEnvironmentOverrides(t *testing.T) {
	t.Setenv("DB_ENABLED", "true")
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5432")
	t.Setenv("DB_CONNECT_TIMEOUT", "7s")

	cfg, err := NewConfig("")
	if err != nil {
		t.Fatalf("NewConfig() error = %v", err)
	}
	if !cfg.DatabaseEnabled {
		t.Fatal("database should be enabled by environment override")
	}
	if cfg.Database.Host != "localhost" || cfg.Database.Port != 5432 {
		t.Fatalf("database defaults = %s:%d, want localhost:5432", cfg.Database.Host, cfg.Database.Port)
	}
	if cfg.Database.ConnectTimeout != 7*time.Second {
		t.Fatalf("database connect timeout = %s, want 7s", cfg.Database.ConnectTimeout)
	}
}

func TestLoadSampleDatabaseConfiguration(t *testing.T) {
	t.Setenv("DB_ENABLED", "false")
	t.Setenv("DB_CONNECT_TIMEOUT", "5s")
	cfg, err := NewConfig(filepath.Join("..", "sample.application.yml"))
	if err != nil {
		t.Fatalf("NewConfig(sample.application.yml) error = %v", err)
	}
	if cfg.DatabaseEnabled {
		t.Fatal("sample configuration should keep PostgreSQL disabled")
	}
	if cfg.Database.ConnectTimeout != 5*time.Second {
		t.Fatalf("database connect timeout = %s, want 5s", cfg.Database.ConnectTimeout)
	}
}
