package valueobjects

const (
	fileFormatJPEG  string = "JPEG"
	fileFormatKnown string = ""
)

type FileFormat struct {
	s string
}

func NewFileFormat(ext string) *FileFormat {
	switch ext {
	case fileFormatJPEG:
		return FileFormatJPEG()
	}
	return &FileFormat{
		s: fileFormatKnown,
	}
}

func FileFormatJPEG() *FileFormat {
	return &FileFormat{
		s: fileFormatJPEG,
	}
}

func (f *FileFormat) String() string {
	return f.s
}

func (f *FileFormat) FilenameExtension() string {
	switch f.s {
	case fileFormatJPEG:
		return ".jpg"
	}
	return ""
}
