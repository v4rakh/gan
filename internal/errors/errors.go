package errors

var ErrorValidationNotBlank = New(IllegalArgument, "blank values are not allowed")
var ErrorValidationPageGreaterZero = New(IllegalArgument, "page has to be greater 0")
var ErrorValidationPageSizeGreaterZero = New(IllegalArgument, "pageSize has to be greater 0")

var ErrorAnnouncementNotFound = New(NotFound, "announcement not found")
var ErrorAnnouncementCreateFailed = New(GeneralError, "announcement creation failed")
var ErrorAnnouncementUpdateFailed = New(GeneralError, "announcement update failed")
var ErrorAnnouncementDeleteFailed = New(GeneralError, "announcement deletion failed")

var ErrorSubscriptionNotFound = New(NotFound, "subscription not found")
var ErrorSubscriptionAlreadyActive = New(Conflict, "state already active")
var ErrorSubscriptionForbiddenTokenMatch = New(Forbidden, "token doesn't match")
var ErrorSubscriptionCreateFailed = New(GeneralError, "subscription creation failed")
var ErrorSubscriptionUpdateFailed = New(GeneralError, "subscription update failed")
var ErrorSubscriptionDeleteFailed = New(GeneralError, "subscription deletion failed")
