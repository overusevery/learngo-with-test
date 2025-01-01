package playerhandler

import (
	"fmt"
	"net/http"
)

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
