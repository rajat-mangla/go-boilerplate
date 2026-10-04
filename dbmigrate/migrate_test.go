package dbmigrate

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCreateWritesTimestampedMigrationPair(t *testing.T) {
	directory := t.TempDir()
	startTime := time.Now().Unix()

	upPath, downPath, err := Create(directory, "add_users")
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	pattern := regexp.MustCompile(`^[0-9]+_add_users\.(up|down)\.sql$`)
	if !pattern.MatchString(filepath.Base(upPath)) {
		t.Fatalf("up migration filename = %q, want epoch-based up migration", filepath.Base(upPath))
	}
	if !pattern.MatchString(filepath.Base(downPath)) {
		t.Fatalf("down migration filename = %q, want epoch-based down migration", filepath.Base(downPath))
	}
	versionText, _, _ := strings.Cut(filepath.Base(upPath), "_")
	version, err := strconv.ParseInt(versionText, 10, 64)
	if err != nil {
		t.Fatalf("parse migration epoch from %q: %v", filepath.Base(upPath), err)
	}
	if version < startTime || version > time.Now().Unix()+999 {
		t.Fatalf("migration version %d is not a Unix epoch timestamp", version)
	}
	if filepath.Dir(upPath) != directory || filepath.Dir(downPath) != directory {
		t.Fatalf("migration files are not in the requested directory")
	}
	for _, path := range []string{upPath, downPath} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("stat migration %q: %v", path, err)
		}
	}

	secondUpPath, secondDownPath, err := Create(directory, "add_sessions")
	if err != nil {
		t.Fatalf("create second migration pair: %v", err)
	}
	if secondUpPath == upPath || secondDownPath == downPath {
		t.Fatal("consecutive migrations received duplicate paths")
	}
}

func TestCreateAcceptsMigrationNamesOutsideOldWhitelist(t *testing.T) {
	directory := t.TempDir()
	upPath, downPath, err := Create(directory, "add users")
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	for _, path := range []string{upPath, downPath} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("stat migration %q: %v", path, err)
		}
	}
}
