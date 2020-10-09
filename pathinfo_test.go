// +build !integration test

package contentdir

import (
	"reflect"
	"testing"
)

func TestPathInfo(t *testing.T) {

	mock := CreateReaderWriterMock("default")
	dir, err := NewDirectory(testDataRoot, "schemes", mock)
	if err != nil {
		t.Errorf("expected no error, got:  %v", err)
	}

	gotString := dir.PathInfo.Root()
	expectedString := "/contentdirs"
	if expectedString != gotString {
		t.Errorf("expected value=%v, got=%v", expectedString, gotString)
	}

	gotString = dir.PathInfo.ContentPath()
	expectedString = "/contentdirs/schemes"
	if expectedString != gotString {
		t.Errorf("expected value=%v, got=%v", expectedString, gotString)
	}
}

func TestNewPathInfo(t *testing.T) {

	pi := NewPathInfo(testDataRoot, "schemes")

	gotString := pi.Root()
	expectedString := "/contentdirs"
	if expectedString != gotString {
		t.Errorf("expected value=%v, got=%v", expectedString, gotString)
	}

	gotString = pi.ContentPath()
	expectedString = "/contentdirs/schemes"
	if expectedString != gotString {
		t.Errorf("expected value=%v, got=%v", expectedString, gotString)
	}
}

func TestPathInfoAddFile(t *testing.T) {

	pi := NewPathInfo(testDataRoot, "schemes")
	userDir1 := NewFile("dir1")
	userDir2 := NewFile("0dir")
	userDir3 := NewFile("bdir")
	userFile1 := NewFile("9file")
	userFile2 := NewFile("0file")
	userFile3 := NewFile("afile")
	userFile4 := NewFile("test1")
	userFile5 := NewFile("datafile.txt")
	pi.AddFile(userDir1, userFile1)
	pi.AddFile(userDir1, userFile2)
	pi.AddFile(userDir1, userFile3)
	pi.AddFile(userDir2, userFile4)
	pi.AddFile(userDir3, userFile5)

	expectedValues := []string{
		"9file",
		"0file",
		"afile",
	}
	gotValues := make([]string, 0)
	dirKey := pi.UserDirKey(*userDir1)
	ui := pi.userDirs[dirKey]

	for _, v := range ui.Values() {
		gotValues = append(gotValues, v.Name())
	}

	if !reflect.DeepEqual(expectedValues, gotValues) {
		t.Errorf("expected value=%v, got=%v", expectedValues, gotValues)
	}

	if ui.DirInfo != userDir1 {
		t.Errorf("expected value=%v, got=%v", ui.DirInfo, userDir1)
	}

	expectedIndexValues := []string{
		"/contentdirs/schemes/0dir/test1",
		"/contentdirs/schemes/bdir/datafile.txt",
		"/contentdirs/schemes/dir1/0file",
		"/contentdirs/schemes/dir1/9file",
		"/contentdirs/schemes/dir1/afile",
	}
	pi.pathList.SortByName()
	gotIndexValues := make([]string, 0)
	for _, v := range pi.pathList.Values() {
		gotIndexValues = append(gotIndexValues, v.Name())
	}
	if !reflect.DeepEqual(expectedIndexValues, gotIndexValues) {
		t.Errorf("expected value=%v, got=%v", expectedValues, gotValues)
	}

}
