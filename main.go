package unityads_s2s_callbacks

import (
	"net/http"
	"strings"
	"os"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

var UNITYADSSECRET = os.Getenv("UNITYADSSECRET")

func ValidateUnityadsCallback(r *http.Request) bool {
	return generateHash(r.URL.Query().Encode()) == r.FormValue("hmac")
}

func generateHash(params string) string {
	parameters := strings.Split(params, "&")

	params = ""

	for _, p := range parameters {
		kv := strings.Split(p, "=")

		if kv[0] == "hmac" {
			continue
		}

		params = params + p + ","
	}

	params = params[:len(params) - 1]

	secret := UNITYADSSECRET
	if secret == "" {
		secret = "xyzKEY"
	}

	sig := hmac.New(md5.New, []byte(secret))
	sig.Write([]byte(params))

	return hex.EncodeToString(sig.Sum(nil))
}