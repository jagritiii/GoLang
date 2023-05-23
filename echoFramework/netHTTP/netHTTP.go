package netHTTP

import (
	"fmt"
	"net/http"
)

//type MyWebserverType bool
//
//func (m MyWebserverType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "hellooo", r)
//}

type login int
type welcome int

func (l login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is login")
}

func (wl welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is welcome")
}

func myLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<html>
<head>Hi</head>
<bodyy>
<h1>This is login page</h1>
</body>
</html>
`)
}

func myWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<html>
<head>Hi</head>
<bodyy>
<h1>This is Welcome page</h1>
</body>
</html>
`)
}

func main() {
	//1--------
	//http.HandleFunc("/login", myLogin)
	//http.HandleFunc("/welcome", myWelcome)

	//2-----------
	//http.Handle("/login", http.HandlerFunc(myLogin))
	//http.Handle("/welcome", http.HandlerFunc(myWelcome))

	//3-----------
	//var i login
	//var j welcome
	//http.Handle("/login", i)
	//http.Handle("/welcome", j)

	fmt.Println("listening on port 8080")
	http.ListenAndServe("localhost:8080", nil)

}
