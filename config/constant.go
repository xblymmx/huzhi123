package config

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
