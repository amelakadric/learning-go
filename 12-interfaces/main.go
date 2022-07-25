package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logger interface {
	log()
}

type logData struct {
	message  string
	fileName string
}

func (lData *logData) log() {
	err := ioutil.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Println("Storing the log data failed!")
		message := err.Error()
		fmt.Println(message)
	}
}

type loggableString string

func (text loggableString) log() {
	fmt.Println(text)
}

func main() {
	userLog := &logData{"User logged in", "user-log.txt"}
	// do more work ...
	userLog.log()

	message := loggableString("[INFO] User created!")
	// do more work ...
	message.log()

	createLog(message)

	outputValue(message)
	outputValue(userLog)
	// reader := bufio.NewReader(os.Stdin)
}

func createLog(data logger) {
	data.log()
}

func outputValue(value interface{}) {
	fmt.Println(value)
}
