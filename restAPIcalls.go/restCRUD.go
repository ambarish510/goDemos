package restAPIcalls

import (
	"fmt"
	"net/http"
	"time"
)

func testPackage() {
	fmt.Println("Hai")
}

func GetAPICall(TimeOut int, getURL string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Duration(TimeOut) * time.Second,
	}
	req, err := http.NewRequest("GET", getURL, nil)
	for k, v := range headers {
		fmt.Println("k:", k, "v:", v)
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	return resp, err
}
