package core

import "regexp"

type Sizer interface {
	Size() int64
}

const (
	LocalFileDir   = "static/uploads/file"
	MinFileSize    = 1       // bytes
	MaxFileSize    = 5000000 // bytes
	FileType       = "(jpg|gif|p?jpeg|(x-)?png)"
	AcceptFileType = FileType
)

var (
	imageTypes      = regexp.MustCompile(FileType)
	acceptFileTypes = regexp.MustCompile(AcceptFileType)
)

type FileInfo struct {
	Url          string `json:"url,omitempty"`
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Size         int64  `json:"size"`
	Error        string `json:"error,omitempty"`
	DeleteUrl    string `json:"deleteUrl,omitempty"`
	DeleteType   string `json:"deleteType,omitempty"`
}

func (fi *FileInfo) ValidateType() (valid bool) {
	if acceptFileTypes.MatchString(fi.Type) {
		return true
	}
	fi.Error = "FileType not allowed"
	return false
}

func (fi *FileInfo) ValidateSize() (valid bool) {
	if fi.Size < MinFileSize {
		fi.Error = "File is too small"
	} else if fi.Size > MaxFileSize {
		fi.Error = "File is too big"
	} else {
		return true
	}
	return false
}
