package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func callThirdPartyAPI(c context.Context, userID int) (bool, error) {
	time.Sleep(400 * time.Millisecond)

	if c.Err() == context.DeadlineExceeded {
		return false, errors.New("context timeout exceeded!")
	}

	return true, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	userId := 6851118

	// calling the 3rd party API

	isUserSubbed, err := callThirdPartyAPI(ctx, userId)
	if err != nil {
		log.Fatalf("error fetching user status for : %d", userId)
	}

	if isUserSubbed {
		fmt.Printf("This user is subbbed %d", userId)
	}
}
