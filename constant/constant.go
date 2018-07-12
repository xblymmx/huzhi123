package constant

type errCode struct {
	SUCCESS      int
	ERROR        int
	NotFound     int
	LoginError   int
	LoginTimeout int
	NotInActive  int // account not in active
}

var ErrorCode = errCode{
	SUCCESS:      1,
	ERROR:        2,
	NotFound:     404,
	LoginError:   1000,
	LoginTimeout: 1001,
	NotInActive:  1002,
}

type errorMsg struct {
	// common
	InvalidParameter         string
	InvalidFileName          string
	InvalidFileExtensionName string
	InvalidMimeType          string

	// gin
	GinBindingError string

	// category
	InvalidCategoryParentID string
	InvalidCategoryName     string
}

var ErrorMsg = errorMsg{
	InvalidParameter:         "parameter invalid",
	GinBindingError:          "gin context binding and validation error",
	InvalidCategoryParentID:  "invalid category parent id",
	InvalidCategoryName:      "invalid category name",
	InvalidFileName:          "invalid file name",
	InvalidFileExtensionName: "invalid file extension name",
	InvalidMimeType:          "invalid mime type",
}
