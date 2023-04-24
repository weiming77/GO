package main

//import "github.com/stretchr/testify/assert"
import log "github.com/sirupsen/logrus"

const NewValue = "ChangedValue"

func testUpdateArray(t *[2]string) {
	t[0] = NewValue
}

func main() {
	//var t *testing.T
	//testArray := [2]string{"Value1", "Value2"}
	//testUpdateArray(&testArray)
	//assert.Equal(t, NewValue, testArray[0])

	log.Info("This is a log line")
	log.Warn("Another warning log line")
	log.Fatal("This is really bad, exicting...")
}
