package contentdir

import (
	"path"
)

// UserDirInfo stores a user dir along with its files
type UserDirInfo struct {
	*Filelist
	DirInfo *FileInfo
}

// IndexFile stores the full path of a file and maintains a reference to the
// file info
type IndexFile struct {
	name     string
	FileInfo *FileInfo
	DirInfo  *FileInfo
}

// NewUserDirInfo creates a new user dir info struct.
func NewUserDirInfo(userDir *FileInfo) *UserDirInfo {
	return &UserDirInfo{
		Filelist: NewFilelist(),
		DirInfo:  userDir,
	}
}

// PathInfo stores path information
type PathInfo struct {
	rootPath   string
	contentTag string
	userDirs   map[string]*UserDirInfo
	// pathlist contains a sorted list of all paths
	pathList *Filelist
}

// NewPathInfo returns a  path information struct
func NewPathInfo(rootPath string, contentTag string) PathInfo {
	return PathInfo{
		rootPath:   rootPath,
		contentTag: contentTag,
		pathList:   NewFilelist(),
		userDirs:   make(map[string]*UserDirInfo),
	}
}

// UserDirKey returns a unique key for hashing the user directory.
func (p PathInfo) UserDirKey(userDir FileInfo) string {
	return path.Join(p.ContentPath(), userDir.Name())
}

// FileKey returns a unique key for hashing a file
func (p PathInfo) FileKey(userDir FileInfo, file FileInfo) string {
	return path.Join(p.ContentPath(), userDir.Name(), file.Name())
}

// Root returns the root path of the content dir
func (p PathInfo) Root() string { return p.rootPath }

// ContentPath returns the content path of the content dir
func (p PathInfo) ContentPath() string { return path.Join(p.rootPath, p.contentTag) }

// AddFile adds a file to the directory
func (p PathInfo) AddFile(userDir *FileInfo, file *FileInfo) {

	dirKey := p.UserDirKey(*userDir)
	if v, found := p.userDirs[dirKey]; found {
		v.Append(file)
	} else {
		ui := NewUserDirInfo(userDir)
		p.userDirs[dirKey] = ui
		ui.Append(file)
	}
	indexFile := NewFile(p.FileKey(*userDir, *file))
	indexFile.DirInfo = userDir
	p.pathList.Append(indexFile)
}
