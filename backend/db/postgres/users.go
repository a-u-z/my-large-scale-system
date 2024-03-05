package pg

import (
	"database/sql"
	"fmt"
	"log"
)

func createTable(db *sql.DB) error {
	// 創建表格
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL
	);
	`
	_, err := db.Exec(createTableQuery)
	return err
}

func (d *Database) InsertUser(username, password string) error {
	// 建立查詢
	query := `INSERT INTO users (username, password)VALUES ($1, $2);`

	// 準備查詢
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 執行查詢
	_, err = stmt.Exec(username, password)

	return err
}

func (d *Database) UpdateData(newPassword, usernameToUpdate string) error {
	// 更新資料
	updateDataQuery := "UPDATE users SET password = $1 WHERE username = $2"
	_, err := d.db.Exec(updateDataQuery, newPassword, usernameToUpdate)
	return err
}

func (d *Database) DeleteData(usernameToDelete string) error {
	// 刪除資料
	deleteDataQuery := "DELETE FROM users WHERE username = $1"
	_, err := d.db.Exec(deleteDataQuery, usernameToDelete)
	return err
}

func (d *Database) QueryData() {
	// 查詢資料
	query := "SELECT id, username, password FROM users"
	rows, err := d.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var username, password string
		err := rows.Scan(&id, &username, &password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Username: %s, Password: %s\n", id, username, password)
	}
}

func (d *Database) FindOne(username, password string) (bool, error) {
	// 建立查詢
	query := `
SELECT
    EXISTS (
        SELECT
        *
        FROM
        users
        WHERE
        username = $1
        AND
        password = $2
    );
`

	// 準備查詢
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	// 執行查詢
	row := stmt.QueryRow(username, password)

	// 掃描結果
	var exists bool
	err = row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (d *Database) FindUserByUsername(username string) (bool, error) {
	// 建立查詢
	query := `
        SELECT EXISTS (
            SELECT username FROM users WHERE username = $1
        );
    `

	// 準備查詢
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	// 執行查詢
	row := stmt.QueryRow(username)

	// 掃描結果
	var exists bool
	err = row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
