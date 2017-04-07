package httpHelper

// ContentType response content type
type ContentType int

const (
	// ContentTypeJSON ContentType_JSON
	ContentTypeJSON ContentType = iota
	// ContentTypeText ContentType_Text
	ContentTypeText = iota
)

var contentTypes = []string{
	"JSON",
	"Text",
}

func (c ContentType) String() string {
	return contentTypes[c]
}

// Is is this content type
func (c ContentType) Is(c1 ContentType) bool {
	return c == c1
}
