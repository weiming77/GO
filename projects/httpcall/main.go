package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// metric is an application metric
type Metric struct {
	Time   time.Time `json:"time"`
	CPU    float64   `json:"cpu"`
	Memory float64   `json:"memory"`
}

func postMetric(m Metric) error {
	// marshal the Metric into byte slices
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	// create the context with timeout upon request
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer the cancel of the request meaning that everytime the context
	// to be cancel when we exit the function
	defer cancel()

	const url = "https://httpbin.org/post"
	//PS: The body should be passed as io.Reader so we create a
	// bytes.NewReader on top of the data
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return err
	}

	// we set the headers of the content type
	req.Header.Set("Content-Type", "application/json")

	// now we call the service
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	// we also check the status code before ew parse the response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}

	// defer closing the request body
	defer resp.Body.Close()
	// You should not blindly read everyting from the network
	// so our maximum size is one megabyte
	const maxSize = 1 << 20 // 1MB
	// we set an io.limitReader on top of the body with this size
	r := io.LimitReader(resp.Request.Body, maxSize)
	// define the anonymous structure with JSON,
	var reply struct {
		JSON Metric
	}
	// then use the json.NewDecoder to decode the reply
	if err := json.NewDecoder(r).Decode(&reply); err != nil {
		return err
	}
	// finally, we log the reply on
	log.Printf("Got: %+v\n", reply.JSON)

	return nil
}

func main() {
	m := Metric{
		Time:   time.Now(),
		CPU:    0.23,
		Memory: 128,
	}

	if err := postMetric(m); err != nil {
		log.Fatal(err)
	}
}
