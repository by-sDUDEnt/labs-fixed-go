package swagger

import (
	"net/http"
)

var fileServer = http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./docs/dist/")))

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/swaggerui/swagger.yaml" {
		http.ServeFile(w, r, "docs/swagger.yaml")
		return
	}

	fileServer.ServeHTTP(w, r)

	return
}
