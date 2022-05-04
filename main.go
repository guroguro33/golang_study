package main

import (
	"database/sql"
	"fmt"
	"log"

	// コード内では使用しないが、initだけ使用してコンパイルに含めるときはアンスコを使用する
	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Name string
	Age  int
}

var DbConnection *sql.DB

func main() {
	// 第１引数にドライバ、第２引数にSQLを指定
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	// テーブルの生成
	cmd := `CREATE TABLE IF NOT EXISTS person(
				name STRING,
				age  INT)`
	// テーブル作成で結果が返ってこないため_を指定するのが一般的
	_, err := DbConnection.Exec(cmd) // Execは結果が返ってこないメソッド
	if err != nil {
		log.Fatalln(err)
	}

	// データの挿入
	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err = DbConnection.Exec(cmd, "MIke", 23)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// データの更新
	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, 23, "Mike")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// 複数のデータの取得（まずstructを定義しておく）
	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd) // Queryはrowsの数を返却するメソッド
	defer rows.Close()                 // 次のレコードを取得することを考え、deferとする
	var pp []Person                    // 複数のPersonを入れるスライスを定義
	// Nextでrowsを1つ1つ準備する（最初の呼び出しでも必要）
	for rows.Next() {
		var p Person
		// Scanで現在のrowのカラムを引数で指定したアドレスにコピーする
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		// appendでスライスppの最後にpを追加する
		pp = append(pp, p)
	}
	for key, p := range pp {
		fmt.Println(key, p.Name, p.Age)
	}

	// 1つのデータ取得
	cmd = "SELECT * FROM person WHERE age = ?"
	// 1つだけ取得する場合はQueryRowメソッドを使用
	row := DbConnection.QueryRow(cmd, 20)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		// rowがなかった場合、ScanはErrNoRowsを返却する
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	// データの削除
	// cmd = "DELETE FROM person WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, "Mike")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

}
