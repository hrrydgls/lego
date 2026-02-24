package controllers

import (
	"fmt"
	"net/http"
)


func AboutHandler (w http.ResponseWriter, _ *http.Request) {
	hobbies := []string {"Football", "Hiking", "Musics"}
	fmt.Fprint(w, "<h1>About us!</h1>")

	fmt.Fprintln(w, "<p>My hobbies are: </p>")

	fmt.Fprintln(w, "<ul>")

	for _, hobby := range hobbies {
		fmt.Fprintf(w, "<li>%s</li>", hobby)
	}

	fmt.Fprintln(w, "</ul>")

}