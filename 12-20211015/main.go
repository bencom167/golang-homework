package main

import (
	"homework/repo"
)

func main() {
	// Connect database
	repo.DB = repo.ConnectDB()

	// repo.DemoCreate()
	// repo.DemoSelect()
	// repo.DemoUpdate()
	// repo.DemoDelete()
	repo.PrintMemberListByClub(repo.DB)
	repo.PrintClubListByMember(repo.DB)
}
