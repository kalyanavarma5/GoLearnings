package main

import (
	"fmt"
	"net/http"
)

// Handler for the root path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welcome to the *** Home Page!")
}

// Handler for /hello route greets user by name
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Println(w, " *** Hello Kalyan, %s!", name)
}

// Handler for setting a cookie
func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	cookie := &http.Cookie{
		Name:  "username",
		Value: name,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	fmt.Println(w, " *** Cookie set for user: %s", name)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/setcookie", setCookieHandler)

	fmt.Println(" *** Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
