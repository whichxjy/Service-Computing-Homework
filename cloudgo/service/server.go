package service

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!\n")
}

// template for login page
const loginTemplate = `
<!DOCTYPE html>
<html>
	<head>
        <title>Login</title>
    </head>
    <body>
        <form action="/login" method="post">
            Username: <input type="text" name="username">
            Password: <input type="password" name="password">
            <input type="submit" value="Login">
        </form>
    </body>
</html>
`

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.New("webpage").Parse(loginTemplate)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Username:", r.FormValue("username"))
		fmt.Println("Password:", r.FormValue("password"))
	}
}

func NewServer() *negroni.Negroni {
	// create router
	router := mux.NewRouter()
	// register routes
	router.HandleFunc("/", sayHello).Methods("GET")
	router.HandleFunc("/login", login).Methods("GET", "POST")
	// add some default middlewares
	nc := negroni.Classic()
	// use router
	nc.UseHandler(router)
	return nc
}
