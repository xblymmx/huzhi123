package constant

type code struct {
	SUCCESS      int
	ERROR        int
	NotFound     int
	LoginError   int
	LoginTimeout int
	NotInActive  int // account not in active
}

var Code = code{
	SUCCESS:      1,
	ERROR:        2,
	NotFound:     404,
	LoginError:   1000,
	LoginTimeout: 1001,
	NotInActive:  1002,
}

type msg struct {
	// common
	SUCCESS                  string
	InvalidParameter         string
	InvalidFileName          string
	InvalidFileExtensionName string
	InvalidMimeType          string

	// gin
	GinBindingError    string
	GinContextRequired string

	// category
	InvalidCategoryParentID string
	InvalidCategoryName     string
	QueryCategoryError      string
	UpdateCategoryError     string
	CategoryNotExist        string
}

var Msg = msg{
	SUCCESS:                  "success",
	InvalidParameter:         "parameter invalid",
	GinBindingError:          "gin context binding and validation error",
	InvalidCategoryParentID:  "invalid category parent id",
	InvalidCategoryName:      "invalid category name",
	InvalidFileName:          "invalid file name",
	InvalidFileExtensionName: "invalid file extension name",
	InvalidMimeType:          "invalid mime type",
	QueryCategoryError:       "query category error",
	UpdateCategoryError:      "update category error",
	CategoryNotExist:         "category not exists",
	GinContextRequired:       "gin.Context is required",
}
