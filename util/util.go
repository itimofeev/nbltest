package util

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"
)

// Log some common log for tests
var Log = &logrus.Logger{
	Level:     logrus.DebugLevel,
	Out:       os.Stdout,
	Formatter: &logrus.JSONFormatter{},
}

// CheckErr check error is nil and if not panic with message
func CheckErr(err error, msg ...string) {
	if err != nil {
		log.Panicln(msg, err)
	}
}

// CheckOk check that ok and if not panic with message
func CheckOk(ok bool, msg string) {
	if !ok {
		log.Panicln(msg)
	}
}

// GetJSON returns json by input object
func GetJSON(v interface{}) string {
	data, err := json.Marshal(v)
	CheckErr(err)
	return string(data)
}

// PrintJSON prints object in json to console
func PrintJSON(v interface{}) {
	fmt.Println(GetJSON(v))
}


// RandInt64 generates random int64
func RandInt64() int64 {
	return rand.Int63()
}

// DefaultHTTPClient is default for AxxonEndpoint http client instance
var DefaultHTTPClient = NewDefaultHTTPClient()

// NewDefaultHTTPClient returns new default http client instance
func NewDefaultHTTPClient(timeoutOpt ...time.Duration) *http.Client {
	timeout := 30 * time.Second

	if len(timeoutOpt) > 0 {
		timeout = timeoutOpt[0]
	}
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          200,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   timeout,
		ExpectContinueTimeout: timeout,
		ResponseHeaderTimeout: timeout,
	}

	return &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}
}
