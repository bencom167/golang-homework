package main

import (
	"web/repo"
)

func main() {
	// Yêu cầu 1
	// repo.RepoUsers("8080")

	// Yêu cầu 2
	repo.RepoUsersWithMiddleware("8080")
}
