package main

import "errors"

var (
	ErrStatusCode = errors.New("received invalid code")
	ErrInvalidRecords = errors.New("invalid records")
	ErrInvalidSubject = errors.New("invalid subject")
	ErrInvalidMessage = errors.New("invalid message")
	ErrTimeParse = errors.New("convert utc to kst error")
)
