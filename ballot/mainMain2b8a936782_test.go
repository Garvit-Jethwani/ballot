package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

const port = "8080"

func serveRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the ballot system!")
}

func runTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Running tests...")
}

func main() {
	log.Println("ballot is ready to store votes")
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/tests/run", runTest)
	log.Println(http.ListenAndServe(net.JoinHostPort("", port), nil))
}
```

Now, let's write test cases for the `serveRoot` and `runTest` handlers:

```go
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServeRoot(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveRoot)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("serveRoot returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Welcome to the ballot system!"
	response, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(string(response)) != expected {
		t.Errorf("serveRoot returned unexpected body: got %v want %v", response, expected)
	}
}

func TestRunTest(t *testing.T) {
	req, err := http.NewRequest("GET", "/tests/run", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(runTest)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("runTest returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Running tests..."
	response, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(string(response)) != expected {
		t.Errorf("runTest returned unexpected body: got %v want %v", response, expected)
	}
}