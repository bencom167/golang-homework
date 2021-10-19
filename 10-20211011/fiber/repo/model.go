package repo

import (
	"context"
	"fmt"
	"homework/goccy"
	"math/rand"
	"time"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/pgjson"
)

type User struct {
	//tableName struct{} `pg:"public.users"`
	Id       int `pg:"id,pk"`
	Fullname string
	Email    string
	Phone    string
	Age      int
	Sex      string
}

type dbLogger struct{}

var (
	numUsers int        = 10 // Lượng users sẽ khởi tạo nếu chưa có dữ liệu
	DB       *pg.DB          // Kết nối vào CSDL Postgresql
	random   *rand.Rand      // Đối tượng dùng để tạo random number
	users    []User
)

func init() {
	//Mở kết nối vào CSDL Postgresql
	DB = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "123",
		Database: "postgres",
	})

	pgjson.SetProvider(goccy.NewJSONProvider()) //Sử dụng goccy json
	//Log các câu lệnh SQL thực thi để debug
	DB.AddQueryHook(dbLogger{}) //Log query to console

	//Khởi động engine sinh số ngẫu nhiên
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Hàm hook (móc câu vào lệnh truy vấn) để in ra câu lệnh SQL query
func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

// Hàm hook chạy sau khi query được thực thi
func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	bytes, _ := q.FormattedQuery()
	fmt.Println(string(bytes))
	return nil
}

// Hàm sinh ngẫu nhiên danh sách users, sau đó chèn vào DB
func InitUsersList() error {
	transaction, err := DB.Begin()
	if err != nil {
		return err
	}

	users = []User{}
	for i := 0; i < numUsers; i++ {
		user := User{
			Id:       10000 + random.Intn(89999), // Sinh ID: 10000 <= id <= 99998
			Fullname: fake.Person().FirstName + " " + fake.Person().LastName,
			Email:    fake.Email(),
			Phone:    fake.Phone(),
			Age:      22 + random.Intn(44), // Sinh tuổi: 22 <= age <= 65
			Sex:      fake.Person().Gender,
		}
		_, err = transaction.Model(&user).Insert()

		users = append(users, user)

		if !check_err(err, transaction) {
			return err
		}
	}
	return transaction.Commit()
}

/*
Kiểm tra err khác nil thì rollback transaction
*/
func check_err(err error, trans *pg.Tx) bool {
	if err != nil {
		_ = trans.Rollback()
		return false
	}
	return true
}
