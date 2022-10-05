package xgo

import (
	"database/sql"
	// "fmt"
	"log"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriverName = "sqlite3"
	dbName       = "./data.db3"
)

// type user struct {
// 	Username string
// 	Age      int
// 	Job      string
// 	Hobby    string
// }

type logins struct {
	Username string
	Times      int
	Ipaddr      string
	Logintime    string
}

// func Info() {
// 	db, err := sql.Open(dbDriverName, dbName)
// 	if checkErr(err) {
// 		return
// 	}
// 	err = createTable(db)
// 	if checkErr(err) {
// 		return
// 	}
// 	err = insertData(db, user{"zhangsan", 28, "engineer", "play football"})
// 	if checkErr(err) {
// 		return
// 	}
// 	err = insertData(db, user{"yeahs", 18, "meish", "play baseball"})
// 	if checkErr(err) {
// 		return
// 	}
// 	res, err := queryData(db, "zhangsan")
// 	if checkErr(err) {
// 		return
// 	}
// 	fmt.Println(len(res))
// 	for _, val := range res {
// 		fmt.Println(val)
// 	}
// 	// r, err := delByID(db, 1)
// 	// if checkErr(err) {
// 	// 	return
// 	// }
// 	// if r {
// 	// 	fmt.Println("delete row success")
// 	// }
// }

func InsertLog(username string,times int,logintime string,ipaddr string) string {
	db, err := sql.Open(dbDriverName, dbName)
	if checkErr(err) {
		return "db error."
	}
	err = ExistsLogTable(db)
	if checkErr(err) {
		return "table error."
	}
	err = insertLogData(db, logins{username, times, logintime, ipaddr})
	if checkErr(err) {
		return "insert error."
	}
	res, err := queryLogTotal(db)
	if checkErr(err) {
		return "query error."
	}
	return "times:" + strconv.Itoa(len(res))
}

func ExistsLogTable(db *sql.DB) error {
	sql := `create table if not exists "logs" (
		"id" integer primary key autoincrement,
		"username" text not null,
		"times" integer not null,
		"logintime" text,
		"ipaddr" text,
	)`
	_, err := db.Exec(sql)
	return err
}

func insertLogData(db *sql.DB, u logins) error {
	sql := `insert into logs (username, times, logintime, ipaddr) values(?,?,?,?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Username, u.Times, u.Logintime, u.Ipaddr)
	return err
}

func queryLogTotal(db *sql.DB) (l []logins, e error) {
	sql := `select * from logs`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var result = make([]logins, 0)
	for rows.Next() {
		var username, ipaddr, logintime string
		var times, id int
		rows.Scan(&id, &username, &times, &logintime, &ipaddr)
		result = append(result, logins{username, times, logintime, ipaddr})
	}
	return result, nil
}

// func createTable(db *sql.DB) error {
// 	sql := `create table if not exists "users" (
// 		"id" integer primary key autoincrement,
// 		"username" text not null,
// 		"age" integer not null,
// 		"job" text,
// 		"hobby" text
// 	)`
// 	_, err := db.Exec(sql)
// 	return err
// }

// func insertData(db *sql.DB, u user) error {
// 	sql := `insert into users (username, age, job, hobby) values(?,?,?,?)`
// 	stmt, err := db.Prepare(sql)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = stmt.Exec(u.Username, u.Age, u.Job, u.Hobby)
// 	return err
// }

// func queryData(db *sql.DB, name string) (l []user, e error) {
// 	sql := `select * from users`
// 	stmt, err := db.Prepare(sql)
// 	if err != nil {
// 		return nil, err
// 	}
// 	rows, err := stmt.Query()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var result = make([]user, 0)
// 	for rows.Next() {
// 		var username, job, hobby string
// 		var age, id int
// 		rows.Scan(&id, &username, &age, &job, &hobby)
// 		result = append(result, user{username, age, job, hobby})
// 	}
// 	return result, nil
// }

// func delByID(db *sql.DB, id int) (bool, error) {
// 	sql := `delete from users where id=?`
// 	stmt, err := db.Prepare(sql)
// 	if err != nil {
// 		return false, err
// 	}
// 	res, err := stmt.Exec(id)
// 	if err != nil {
// 		return false, err
// 	}
// 	_, err = res.RowsAffected()
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

func checkErr(e error) bool {
	if e != nil {
		log.Fatal(e)
		return true
	}
	return false
}