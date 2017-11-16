package main

import (
	"flag"
	"net/http"
	"os"
	"fmt"
	"net/url"
	"crypto/tls"
	"time"
)

type config struct {
	url               *url.URL
	retry             int
	skipVerify        bool
	retryInterval     time.Duration
	connectionTimeout time.Duration
	operationTimeout  time.Duration
}

var lastMessage string

func main() {

	cfg := processFlags()

	printStartUp(cfg)

	wait(cfg)

}

func processFlags() *config {

	var err error
	c := config{}

	flag.DurationVar(&c.operationTimeout, "t", 60 * time.Second, "Operation timeout")
	flag.BoolVar(&c.skipVerify, "s", false, "SSH skip host verification")
	flag.DurationVar(&c.connectionTimeout, "c", 5 * time.Second, "Connection timeout")
	flag.DurationVar(&c.retryInterval, "r", 1 * time.Second, "Retry interval")

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}


	c.url, err = url.ParseRequestURI(flag.Arg(0))
	if err != nil {
		fmt.Printf("Error:\n    %s\n", err)
		os.Exit(1)
	}

	return &c
}

func printStartUp(cfg *config) {
	fmt.Printf("Wait status 200 OK from:\n    %s\n", cfg.url)
}
func wait(cfg *config) {

	c := newHttpClient(cfg)

	start := time.Now()
	for {

		t := time.Now()

		if t.Sub(start) > cfg.operationTimeout {
			s := fmt.Sprintf("\nOperation timeout %s", cfg.operationTimeout)
			done(1, s)
		}

		resp, err := c.Get(cfg.url.String())
		if err != nil {
			printMessage(err.Error())
		} else {

			if resp.StatusCode == http.StatusOK {
				s := fmt.Sprintf("OK\nTime: %s\n", t.Sub(start))
				done(0, s)
			} else {
				s := fmt.Sprintf("Response code is: %d", resp.StatusCode)
				printMessage(s)
			}
		}
		time.Sleep(cfg.retryInterval)
	}

}

func done(code int, msg string) {
	fmt.Printf("\n%s\n", msg)
	os.Exit(code)
}

func newHttpClient(cfg *config) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.skipVerify},
	}
	return &http.Client{
		Timeout:   time.Duration(cfg.connectionTimeout),
		Transport: tr,
	}
}

func printMessage(s string) {
	if s == lastMessage {
		fmt.Print(".")
	} else {
		lastMessage = s
		fmt.Printf("\n%s ", s)
	}
}
