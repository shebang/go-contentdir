// +build !integration test

package contentdir

import (
	"testing"
)

func TestFileInfo(t *testing.T) {

	fi := FileInfo{name: "test"}
	expectedString := "test"
	gotString := fi.Name()

	if expectedString != gotString {
		t.Errorf("expected value=%s, got value=%s", expectedString, gotString)
	}
}
