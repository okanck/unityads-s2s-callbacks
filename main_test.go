package unityads_s2s_callbacks

import (
	"testing"
	"net/http"
	"strings"
	"bufio"
)

func TestValidateUnityadsCallback(t *testing.T) {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader("GET /?productid=1234&sid=1234567890&oid=0987654321&hmac=106ed4300f91145aff6378a355fced73 HTTP/1.1\r\n\r\n")))
	if err != nil {
		t.Fatalf("Error while creating request: %v", err)
	}

	req.Proto = "HTTP/1.1"
	req.Close = false
	req.ContentLength = -1

	result := ValidateUnityadsCallback(req)

	if !result {
		t.Fatalf("Excepted true but got %v", result)
	}
}