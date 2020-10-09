// +build !integration test

package contentdir

import "testing"

func TestFilterByExtension(t *testing.T) {

	fname := "test.yaml"
	extList := []string{".txt", ".yaml"}

	if !FilterByExtension(fname, extList) {
		t.Errorf("expected FilterByExtension fname=%s, extList=%v to return true", fname, extList)
	}

	fname = "test.mp3"

	if FilterByExtension(fname, extList) {
		t.Errorf("expected FilterByExtension fname=%s, extList=%v to return false", fname, extList)
	}

}
