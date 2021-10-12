package repo

import (
	"fmt"
	"log"
	"net/http"
)

/*
	Hàm handle với path: /users
*/
func usersHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listUsers(w, r)
	} else if r.Method == "POST" {
		createUsers(w, r)
	} else {
		defaultHandle(w, r)
	}
}

/*
	Hàm handle với path: /users/id
*/
func usersIDHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getUser(w, r)
	} else if r.Method == "PUT" {
		updateUser(w, r)
	} else if r.Method == "DELETE" {
		deleteUser(w, r)
	} else {
		defaultHandle(w, r)
	}
}

/*
	Hàm handle mặc định
*/
func defaultHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Invalid method!")
}

/*
	Hàm quản lý user repository
*/
func RepoUsers(port string) {
	http.HandleFunc("/users", usersHandle)
	http.HandleFunc("/users/id", usersIDHandle)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

/*
	Middleware ghi log request
*/
func middlewareUsersHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] - %v", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

/*
	Hàm quản lý user repository với middleware ghi log request
*/
func RepoUsersWithMiddleware(port string) {
	mux := http.NewServeMux()

	mux.Handle("/users", middlewareUsersHandle(http.HandlerFunc(usersHandle)))
	mux.Handle("/users/id", middlewareUsersHandle(http.HandlerFunc(usersIDHandle)))

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
