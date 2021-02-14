package domain

import "errors"

var ErrorNotFound = errors.New("not found")
var ErrorDeleteFailed = errors.New("deletion failed")
var ErrorCreateFailed = errors.New("creation failed")
var ErrorUpdateFailed = errors.New("update failed")

var ErrorForbiddenTokenMatch = errors.New("token doesn't match")
var ErrorConflict = errors.New("already existing")

var ErrorPageGreaterZero = errors.New("page has to be greater 0")
var ErrorPageSizeGreaterZero = errors.New("pageSize has to be greater 0")
