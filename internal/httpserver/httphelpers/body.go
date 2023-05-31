package httphelpers

import (
	"encoding/json"
	"net/http"
)

func ReadBody(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
