package test

import (
	"fmt"
	"homework/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddMemberToClub(t *testing.T) {
	db := repo.GetInstance()
	err := repo.AddMemberToClub(db)
	assert.Nil(t, err)
}

func Test_GetClubByName(t *testing.T) {
	db := repo.GetInstance()
	club, err := repo.GetClubByName(db, "Sport")
	assert.Nil(t, err)

	fmt.Println(club)

	fmt.Println(club.Name)
	for _, m := range club.Members {
		fmt.Println("	" + m.Name)
	}
}

func Test_GetMemberByName(t *testing.T) {
	db := repo.GetInstance()
	member, err := repo.GetMemberByName(db, "Bob")
	assert.Nil(t, err)

	fmt.Println(member)

	fmt.Println(member.Name)
	for _, c := range member.Clubs {
		fmt.Println("	" + c.Name)
	}
}

/*
	Các Unit Test được viết thêm ở đây
*/
func Test_UpdateMemberName(t *testing.T) {
	db := repo.GetInstance()
	err := repo.UpdateMemberNameByID(db, "6sxdPMtA", "Bob")
	assert.Nil(t, err)
}

func Test_UpdateClubName(t *testing.T) {
	db := repo.GetInstance()
	err := repo.UpdateClubNameByID(db, "inVTY0ur", "Sport")
	assert.Nil(t, err)
}

func Test_DeleteMemberByID(t *testing.T) {
	db := repo.GetInstance()
	err := repo.DeleteMemberByID(db, "6sxdPMtA")
	assert.Nil(t, err)
}

func Test_DeleteMemberByName(t *testing.T) {
	db := repo.GetInstance()
	err := repo.DeleteMemberByName(db, "Bob")
	assert.Nil(t, err)
}

func Test_PrintMemberListByClub(t *testing.T) {
	db := repo.GetInstance()
	repo.PrintMemberListByClub(db)
}

func Test_PrintClubListByMember(t *testing.T) {
	db := repo.GetInstance()
	repo.PrintClubListByMember(db)
}

func Test_FindTheBiggestClubs(t *testing.T) {
	db := repo.GetInstance()
	repo.FindTheBiggestClubs(db)
}
