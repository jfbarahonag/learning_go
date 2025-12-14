package main

import "fmt"

type ServerStatus uint8

const (
	StatusIdle ServerStatus = iota
	StatusConnected
	StatusError
	StatusRetrying 
)

var statusName = map[ServerStatus] string {
	StatusIdle: "Idle",
	StatusConnected: "Connected",
	StatusError: "Error",
	StatusRetrying: "Retrying",
}

func main() {
	serverStaus := networkValidation(StatusIdle)
	fmt.Println("Server status: ", serverStaus)
	
	nextServerStatus := networkValidation(serverStaus)
	fmt.Println("Server status: ", nextServerStatus)

}

func (status ServerStatus) String() string {
	return statusName[status]
}

func networkValidation(serverStatus ServerStatus) ServerStatus {
	switch serverStatus {
	case StatusIdle:
		return StatusConnected
	case StatusConnected, StatusRetrying:
		return StatusIdle
	case StatusError:
		return StatusError
	default:
		panic(fmt.Errorf("Unknown status %d", serverStatus))
	}
}