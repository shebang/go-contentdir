// +build !integration test

package contentdir

import (
	"reflect"
	"testing"
)

func TestNewFilelist(t *testing.T) {

	p := NewFilelist()
	expectedInt := 0
	gotInt := p.Count()
	if expectedInt != gotInt {
		t.Errorf("expected value=%d, got=%d", expectedInt, gotInt)
	}
	expectedInt = 0
	gotInt = cap(p.files)
	if expectedInt != gotInt {
		t.Errorf("expected value=%d, got=%d", expectedInt, gotInt)
	}

	p = NewFilelist(10)
	expectedInt = 10
	gotInt = p.Count()
	if expectedInt != gotInt {
		t.Errorf("expected value=%d, got=%d", expectedInt, gotInt)
	}
}

func TestFilelistAppend(t *testing.T) {

	p := NewFilelist()
	p.Append(NewFile("/path1"))
	p.Append(NewFile("/path2"))
	expectedInt := 2
	gotInt := p.Count()
	if expectedInt != gotInt {
		t.Errorf("expected value=%d, got=%d", expectedInt, gotInt)
	}

	expectedValues := []string{
		"/path1",
		"/path2",
	}
	gotValues := make([]string, 0)
	for _, v := range p.Values() {
		gotValues = append(gotValues, v.Name())
	}

	if !reflect.DeepEqual(expectedValues, gotValues) {
		t.Errorf("expected value=%+v, got=%+v", expectedValues, gotValues)
	}

}

func TestFilelistSortByName(t *testing.T) {

	p := NewFilelist()
	p.Append(NewFile("/zpath"))
	p.Append(NewFile("/kpath"))
	p.Append(NewFile("/apath"))
	p.Append(NewFile("/2path"))
	p.Append(NewFile("/0path"))
	p.Append(NewFile("/9path"))

	p.SortByName()
	expectedValues := []string{
		"/0path",
		"/2path",
		"/9path",
		"/apath",
		"/kpath",
		"/zpath",
	}
	gotValues := make([]string, 0)
	for _, v := range p.Values() {
		gotValues = append(gotValues, v.Name())
	}

	if !reflect.DeepEqual(expectedValues, gotValues) {
		t.Errorf("expected value=%v, got=%v", expectedValues, gotValues)
	}

}
