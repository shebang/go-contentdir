// +build integration

package contentdir

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func TestIntegrationNewDirectory(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)
	rootPath := path.Join(path.Dir(filename), "./testdata/contentdir")

	dir, err := NewDirectory(rootPath, "schemes")
	if err != nil {
		t.Errorf("expected no error, got=%v", err)
	}

	contentPath := path.Join(rootPath, "templates")
	if _, err = os.Stat(contentPath); os.IsExist(err) {
		t.Errorf("expected non existant path=%s", contentPath)
	}

	dir, err = NewDirectory(rootPath, "templates")
	if err != nil {
		t.Errorf("expected no error, got=%v", err)
	}

	defer os.Remove(dir.PathInfo.ContentPath())
}

func TestIntegrationReadDir(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)
	rootPath := path.Join(path.Dir(filename), "./testdata/contentdir")

	dir, err := NewDirectory(rootPath, "schemes")
	if err != nil {
		t.Errorf("expected no error, got=%v", err)
	}
	dir.ReadDir()
}
