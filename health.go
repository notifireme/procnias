package procnias

import (
	"fmt"
	"net/http"
)

func (a App) health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ok!")
}
