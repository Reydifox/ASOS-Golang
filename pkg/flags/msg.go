package flags

var MsgFlags = map[int]string{
	SUCCESS:                  "ok",
	ERROR:                    "fail",
	INVALID_PARAMS:           "invalid request parameter",
	ERROR_EXIST_EXAMPLE:      "example already exist",
	ERROR_EXIST_EXAMPLE_FAIL: "failed to get existing example",
	ERROR_NOT_EXIST_EXAMPLE:  "failed to get example",
	ERROR_EXIST_MESSAGE:      "message already exist",
	ERROR_EXIST_MESSAGE_FAIL: "failed to get existing message",
	ERROR_NOT_EXIST_MESSAGE:  "failed to get message",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
