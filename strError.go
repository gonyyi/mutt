package mutt 

type StrError string

func (e StrError) Error() string {
	return string(e)
}

const (
	ERR_USERID_EXIST            StrError = "user exist"
	ERR_USERID_NOT_EXIST        StrError = "user not exist"
	ERR_MISSING_REQUIRED_FIELDS StrError = "missing required field"
	ERR_BAD_CREDENTIAL          StrError = "bad credential"
	ERR_DISABLED_ID             StrError = "disabled user account"
	ERR_USER_NOT_IN_GROUP       StrError = "user not in group"
	ERR_ENCRYPTION_CIPHER_SHORT StrError = "cipher text too short"
	ERR_KEY_NOT_EXIST           StrError = "key not exists"
)

