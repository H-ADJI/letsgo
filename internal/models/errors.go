package models

import "errors"

var ErrNoRecords = errors.New("models: no record match found")
var ErrInvalidCreds = errors.New("models: invalid  credentials")
var ErrDuplicateEmail = errors.New("models: duplicate email")
