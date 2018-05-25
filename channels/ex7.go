package main

/*
*********Passing structs into a channel and logging channel status

NOTE: Order of execution is labeled

Only using one go routine so no waitgroups needed

*/
import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

// add buffer length as second argument

func main() {

	go logger()
	// 1. starts running logger concurrently

	defer func() {
		close(logCh)
		//6.channel will close after main is done executing
	}()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	//2.message is received by channel

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	//4.second message is received by channel

	time.Sleep(100 * time.Millisecond)
	//wait 1/10th of a sec to give way for the buffer created above otherwise logentries won't get printed
}

func logger() {
	for entry := range logCh {
		fmt.Printf(" %v - [%v]%v\n", entry.time, entry.severity, entry.message)
		// 3&5.logEntry field values are output

	}
}
