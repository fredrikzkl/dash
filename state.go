package main

type programState string

var (
	MAIN_STATE    programState = "MAIN"
	ADD_STATE     programState = "ADD_INPUT"
	COMMAND_STATE programState = "COMMAND_INPUT"
)
