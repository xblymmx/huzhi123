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
	ServerError              string
	UserSignedOut            string
	InvalidParameter         string
	InvalidFileName          string
	InvalidFileExtensionName string
	InvalidMimeType          string

	// user
	AccountOrPasswordWrong     string
	InvalidUserSignUpParameter string
	UserNameExists             string
	UserEmailExists            string
	CreateUserError            string

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
	SUCCESS:                    "success",
	InvalidParameter:           "parameter invalid",
	GinBindingError:            "gin context binding and validation error",
	InvalidCategoryParentID:    "invalid category parent id",
	InvalidCategoryName:        "invalid category name",
	InvalidFileName:            "invalid file name",
	InvalidFileExtensionName:   "invalid file extension name",
	InvalidMimeType:            "invalid mime type",
	QueryCategoryError:         "query category error",
	UpdateCategoryError:        "update category error",
	CategoryNotExist:           "category not exists",
	GinContextRequired:         "gin.Context is required",
	ServerError:                "server has encountered some error",
	UserSignedOut:              "user signed out",
	AccountOrPasswordWrong:     "user account or password wrong",
	InvalidUserSignUpParameter: "invalid user sign up parameter",
	UserEmailExists:            "user name already exists",
	UserNameExists:             "user email already exists",
	CreateUserError:            "creating user error",
}

// user constants
type user struct {
	// user account status
	StatusNotActive int
	StatusActive    int
	StatusFrozen    int

	// user gender
	GenderMale    int
	GenderFemale  int
	GenderUnknown int

	// role
	RoleBasic      int
	RoleEditor     int
	RoleAdmin      int
	RoleSuperAdmin int
}

var User = user{
	StatusNotActive: 1,
	StatusActive:    2,
	StatusFrozen:    3,

	GenderMale:    1,
	GenderFemale:  2,
	GenderUnknown: 3,

	RoleBasic:      1,
	RoleEditor:     2,
	RoleAdmin:      4,
	RoleSuperAdmin: 8,
}
