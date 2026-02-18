package controllers 
import (
	"fmt"
	"net/http"
	
)
func HomeHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello world from Controller!</h1>")
}