package model

type ZipFile struct {
	Name string `json:"filename"`
	ZipSize float64 `json:"archive_size"`
	TotalSize float64 `json:"total_size"`
	FilesNum int `json:"total_files"`
	Files []*File `json:"files"`
}

type File struct {
	Path string `json:"file_path"`
	Size float64 `json:"size"`
	Type string `json:"mimetype"`
}																										