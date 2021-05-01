package domain

import "errors"

var ErrorNotFound = errors.New("not found")
var ErrorDeleteFailed = errors.New("deletion failed")
var ErrorCreateFailed = errors.New("creation failed")
var ErrorUpdateFailed = errors.New("update failed")

var ErrorForbiddenTokenMatch = errors.New("token doesn't match")
var ErrorConflict = errors.New("already existing")

var ErrorValidationNotBlank = errors.New("blank values are not allowed")
var ErrorValidationPageGreaterZero = errors.New("page has to be greater 0")
var ErrorValidationPageSizeGreaterZero = errors.New("pageSize has to be greater 0")
