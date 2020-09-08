package domain

import "errors"

var ErrorAnnouncementNotFound = errors.New("no announcement found")
var ErrorAnnouncementDeleteFailed = errors.New("announcement deletion failed")
var ErrorAnnouncementCreateFailed = errors.New("announcement creation failed")
var ErrorAnnouncementUpdateFailed = errors.New("announcement update failed")
