package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 创建数据库连接
func CreateDataBase() *sql.DB {
	//"localhost"        数据库地址
	//"3306"             数据库端口号
	//"root"             数据库用户名
	//"123456"           数据库密码
	//"databaseproject"  数据库名
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/gindemo")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("连接数据库成功")
	}
	//defer db.Close()
	return db
}

// 创建表
func CreateUsersTable(db *sql.DB) {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT,
            username TEXT NOT NULL,
            password INT NOT NULL,
            PRIMARY KEY (id)
        )
    `
	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("创建表成功")
	}
}

// 插入数据
func InsertData(db *sql.DB, username string, password string) {
	insertSQL := "INSERT INTO users(username, password) VALUES(?, ?)"
	result, err := db.Exec(insertSQL, username, password)
	if err != nil {
		panic(err.Error())
	} else {
		id, _ := result.LastInsertId()
		fmt.Println("插入数据成功，id为", id)
	}
}

// 删除数据
func deleteData(db *sql.DB, username string) {
	deleteSQL := "DELETE FROM users WHERE username=?"
	_, err := db.Exec(deleteSQL, username)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("删除数据成功")
	}
}

// 更新数据
func updateData(db *sql.DB, username string, password int) {
	updateSQL := "UPDATE users SET password=? WHERE username=?"
	_, err := db.Exec(updateSQL, password, username)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("更新密码成功")
	}
}

// 查询数据
func selectData(db *sql.DB) {
	selectSQL := "SELECT id, username, password FROM users"
	rows, err := db.Query(selectSQL)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		var password string
		err := rows.Scan(&id, &username, &password)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("id: %d, username: %s, password: %d\n", id, username, password)
	}

}

// 条件查询
func SelectData1(db *sql.DB, username string) *sql.Row {
	var id int
	var password string
	selectSQL := "SELECT * FROM users WHERE username=?"
	user := db.QueryRow(selectSQL, username)
	err := user.Scan(&id, &username, &password)
	if err != nil {
		fmt.Println("查询数据失败")
		panic(err.Error())
	}
	return user
	//else {
	//	fmt.Println("查询数据成功，结果为：", user)
	//}
	//var user User
	//queryData := `SELECT * FROM users WHERE name=?`
	//err = db.QueryRow(queryData, "张三").Scan(&user.ID, &user.Name, &user.Age)
	//if err != nil {
	//	fmt.Println("查询数据失败")
	//	return
	//}
	//fmt.Println("查询数据成功，结果为：", user)
}

func CloseDataBase(db *sql.DB) {
	defer db.Close()

}
