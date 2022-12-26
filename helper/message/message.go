package message

const (
	CODE_OK                 = 200
	CODE_CREATED            = 201
	CODE_BAD_REQUEST        = 400
	CODE_UNAUTHORIZED       = 401
	CODE_NOT_FOUND          = 401
	CODE_METHOD_NOT_ALLOWED = 401
)

// General error/desc message
const (
	MSG_SUCCESS         = "Success"
	MSG_NOTFOUND        = "Not found"
	MSG_ALREADYEXISTS   = "Data already exists"
	MSG_BADROUTING      = "Inconsistent mapping between route and handler"
	MSG_INVALID_REQUEST = "Invalid request parameter(s)"
	MSG_INTERNAL_ERROR  = "Error has been occured while processing request"
	MSG_NO_AUTH         = "No Authorization"
	MSG_INVALID_HEADER  = "Invalid header"
)

// General error/desc database message
const (
	MSG_ERR_DB        = "Error has been occured while processing database request"
	MSG_NO_DATA       = "Data is not found"
	MSG_ERR_SAVE_DATA = "Data cannot be saved, please check your request"
	MSG_ERR_DEL_DATA  = "Data cannot be deleted, please check your request"
)

// General error/desc field validation Message
const (
	MSG_ERR_CHAR_LENGTH = "Character length from %v until %v"
	MSG_ERR_REQUIRED    = "Required field"
)
