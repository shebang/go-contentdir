// +build !integration test

package contentdir

import (
	"os"
	"path"
	"testing"
)

var testDataRoot = "/contentdirs"
var testDataDirs = map[string]string{
	"schemes": "/contentdirs/schemes",
}

type FSMockReturn func() ([]FileInfo, error)

func MockReturnFiles(args []string, files []FileInfo) FSMockReturn {
	return func() ([]FileInfo, error) {
		return files, nil
	}
}
func MockReturnError(args []string, err error) FSMockReturn {
	return func() ([]FileInfo, error) {
		return nil, err
	}
}
func MockReturnNoError(args []string) FSMockReturn {
	return func() ([]FileInfo, error) {
		return nil, nil
	}
}

var testDataDirectoryTree = map[string]map[string]map[string]FSMockReturn{
	"default": map[string]map[string]FSMockReturn{
		"/contentdirs": map[string]FSMockReturn{
			"stat": MockReturnFiles([]string{}, []FileInfo{
				{name: "/contentdirs", isdir: true},
			}),
			"readdir": MockReturnFiles([]string{}, []FileInfo{
				{name: "schemes", isdir: true},
			}),
		},
		"/contentdirs/schemes": map[string]FSMockReturn{
			"readdir": MockReturnFiles([]string{}, []FileInfo{
				{name: "default", size: 64, mode: 2147484141, isdir: true},
			}),
		},
		"/contentdirs/schemes/default": map[string]FSMockReturn{
			"readdir": MockReturnFiles([]string{}, []FileInfo{
				{name: "default-dark.yam", size: 64, mode: 0, isdir: false},
			}),
		},
		"/contentdirs/templates": map[string]FSMockReturn{
			"stat":  MockReturnError([]string{}, &os.PathError{Err: os.ErrNotExist}),
			"mkdir": MockReturnNoError([]string{"/contentdirs/templates"}),
		},
		"/contentdirs/non-existant": map[string]FSMockReturn{
			"stat": MockReturnError([]string{}, &os.PathError{Err: os.ErrNotExist}),
		},
	},
}

type ReaderWriterMock struct {
	// presets is used to select a test data preset
	testData map[string]map[string]FSMockReturn
}

func CreateReaderWriterMock(preset string) *ReaderWriterMock {

	return &ReaderWriterMock{
		testData: testDataDirectoryTree[preset],
	}
}

func (mock *ReaderWriterMock) returnMockValues(path string, mockValue FSMockReturn) (fi []FileInfo, err error) {
	var retval []FileInfo
	retval, err = mockValue()
	if err == nil {
		fi = retval
	}
	return
}

func (mock *ReaderWriterMock) ReadDir(path string) (fi []os.FileInfo, err error) {
	if _, ok := mock.testData[path]; ok {
		if mockdata, ok := mock.testData[path]["readdir"]; ok {
			var retval []FileInfo
			retval, err = mockdata()
			fi = make([]os.FileInfo, len(retval)-1)
			for _, k := range retval {
				fi = append(fi, &k)
			}

		}
	}
	return
}
func (mock *ReaderWriterMock) Mkdir(path string, perm os.FileMode) (err error) {
	if _, ok := mock.testData[path]; ok {
		if mockdata, ok := mock.testData[path]["mkdir"]; ok {
			_, err = mockdata()
		}
	}
	return
}
func (mock *ReaderWriterMock) Rmdir(path string) error {
	return nil
}

func (mock *ReaderWriterMock) Stat(name string) (fi os.FileInfo, err error) {
	if _, ok := mock.testData[name]; ok {
		if mockdata, ok := mock.testData[name]["stat"]; ok {
			var fileInfo []FileInfo
			fileInfo, err = mock.returnMockValues(name, mockdata)
			if len(fileInfo) > 0 {
				fi = &fileInfo[0]
			}
		}
	}
	return
}

func TestNewDirectory(t *testing.T) {

	mock := CreateReaderWriterMock("default")
	optsNew := Options{readData: true}
	filter := Filter{}
	dir, err := NewDirectory(testDataRoot, "schemes", filter, mock, optsNew)
	if err != nil {
		t.Errorf("expected no error, got: %v (dir: %v)", err, dir)
	}

	gotString := dir.contentTag
	expectedString := dir.contentTag
	if expectedString != gotString {
		t.Errorf("expected value=%v, got=%v", expectedString, gotString)
	}

}
func TestNewDirectoryNoContentDir(t *testing.T) {

	mock := CreateReaderWriterMock("default")
	optsNew := Options{readData: true}
	filter := Filter{}
	dir, err := NewDirectory(testDataRoot, "templates", filter, mock, optsNew)
	if err != nil {
		t.Errorf("expected no error, got: %v (dir: %v)", err, dir)
	}

}

func TestNewDirectoryError(t *testing.T) {

	mock := CreateReaderWriterMock("default")
	optsNew := Options{readData: true}
	filter := Filter{}
	_, err := NewDirectory(path.Join(testDataRoot, "non-existant"), "schemes", filter, mock, optsNew)
	if err == nil {
		t.Errorf("expected error, got=nil")
	}

}

func TestReadDir(t *testing.T) {

	mock := CreateReaderWriterMock("default")
	optsNew := Options{readData: true}
	filter := Filter{}
	dir, err := NewDirectory(testDataRoot, "schemes", filter, mock, optsNew)
	if err != nil {
		t.Errorf("expected no error, got=%v", err)
	}
	err = dir.ReadDir()

}
