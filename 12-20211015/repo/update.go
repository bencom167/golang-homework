package repo

import (
	"fmt"
	"homework/model"

	"gorm.io/gorm"
)

/*
	Sửa tên thành viên theo ID
*/
func UpdateMemberNameByID(db *gorm.DB, memberID string, memberNewName string) (err error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&model.Member{}).Where("id = ?", memberID).Update("name", memberNewName).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

/*
	Sửa tên Club theo ID
*/
func UpdateClubNameByID(db *gorm.DB, clubID string, clubNewName string) (err error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&model.Club{}).Where("id = ?", clubID).Update("name", clubNewName).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

/*
	Xoá 1 thành viên theo ID
*/
func DeleteMemberByID(db *gorm.DB, memberID string) (err error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("id = ?", memberID).Delete(&model.Member{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

/*
	Xoá 1 thành viên theo Tên
*/
func DeleteMemberByName(db *gorm.DB, memberName string) (err error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("name = ?", memberName).Delete(&model.Member{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

/*
	Trả về danh sách tên thành viên và số lượng club mà thành viên đó tham gia
*/
func PrintMemberListByClub(db *gorm.DB) (err error) {
	query := `SELECT m.name AS Name, mcc.club_count AS Clubs FROM member m
	JOIN (SELECT member_id, count(member_id) AS club_count FROM member_club GROUP BY member_id) AS mcc
	ON m.id = mcc.member_id`

	type MemberCountClubs struct {
		Name  string
		Clubs int
	}

	var members []MemberCountClubs
	db.Raw(query).Scan(&members)

	fmt.Printf("NAME            CLUBS\n")
	for _, member := range members {
		fmt.Printf("%-16v%v\n", member.Name, member.Clubs)
	}

	return nil
}

/*
	Trả về danh sách tên club cùng số lượng các thành viên đã tham gia
*/
func PrintClubListByMember(db *gorm.DB) (err error) {
	query := `SELECT c.name AS Name, mcc.member_count AS Members FROM club c
	JOIN (SELECT club_id, count(club_id) AS member_count FROM member_club GROUP BY club_id) AS mcc
	ON c.id = mcc.club_id`

	type ClubCountMembers struct {
		Name    string
		Members int
	}

	var clubs []ClubCountMembers
	db.Raw(query).Scan(&clubs)

	fmt.Printf("NAME            MEMBERS\n")
	for _, club := range clubs {
		fmt.Printf("%-16v%v\n", club.Name, club.Members)
	}

	return nil
}

/*
	Chỉ ra thông tin những club có số lượng thành viên tham gia nhiều nhất
*/
func FindTheBiggestClubs(db *gorm.DB) (err error) {

	query := `SELECT c.name AS Name, mcc.member_count AS Members FROM club c
	JOIN (SELECT club_id, COUNT(club_id) AS member_count FROM member_club GROUP BY club_id) AS mcc
	ON c.id = mcc.club_id
	ORDER BY mcc.member_count DESC`

	type ClubMaxMembers struct {
		Name    string
		Members int
	}

	var clubs []ClubMaxMembers
	db.Raw(query).Scan(&clubs)

	max := clubs[0].Members

	fmt.Printf("NAME            MEMBERS\n")
	for _, club := range clubs {
		if club.Members == max {
			fmt.Printf("%-16v%v\n", club.Name, club.Members)
		} else {
			break
		}
	}

	return nil
}

/*
	Chỉ ra thông tin 1 club có số lượng thành viên tham gia nhiều nhất
*/
func FindTheBiggestOneClub(db *gorm.DB) (err error) {

	query := `SELECT cmc.Name AS Name, MAX(cmc.Clubs) AS Members FROM
	(SELECT c.name AS Name, mcc.member_count AS Clubs FROM club c
	JOIN (SELECT club_id, COUNT(club_id) AS member_count FROM member_club GROUP BY club_id) AS mcc
	ON c.id = mcc.club_id) AS cmc
	GROUP BY cmc.Name
	ORDER BY MAX(cmc.Clubs) DESC 
	LIMIT 1`

	type ClubMaxMembers struct {
		Name    string
		Members int
	}

	var clubs []ClubMaxMembers
	db.Raw(query).Scan(&clubs)

	fmt.Printf("NAME            MEMBERS\n")
	for _, club := range clubs {
		fmt.Printf("%-16v%v\n", club.Name, club.Members)
	}

	return nil
}
