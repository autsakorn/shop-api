package types

// ResponseCode defines map response code
/*
	Success        200
	CreatedSuccess 201
	BadRequest     400
	Forbidden      403
*/
var ResponseCode = map[string]int{
	"Success":        200,
	"CreatedSuccess": 201,
	"BadRequest":     400,
	"Forbidden":      403,
}
