// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

type Comments []string

type FBWall struct {
	Name       string
	DateJoined time.Time
	comment    *Comments
}

func main() {
	newWallMessage := &FBWall{
		Name:       "Lee Wei Ming",
		DateJoined: time.Now(),
		comment:    &Comments{},
	}
	fmt.Println(*newWallMessage)
}
