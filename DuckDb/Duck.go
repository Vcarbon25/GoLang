package main

// para executar o código primeiro deve estar no linux, e no terminal nessa pasta e rodar
// go mod init
//go mod tidy  esses dois comandos criarão os arquivos .mod e .sum
//depois execute o programa com go run Duck.go
//o banco de dados será criado na pasta
import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/duckdb/duckdb-go/v2"
)

func main() {
	db, err := sql.Open("duckdb", "./teste.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("COnectou ao banco")
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE people (id INTEGER, name VARCHAR)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Criou a tabela")
	_, err = db.Exec(`INSERT INTO people VALUES (42, 'John')`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inseriu os dados")
	var (
		id   int
		name string
	)
	row := db.QueryRow(`SELECT id, name FROM people`)
	err = row.Scan(&id, &name)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println("no rows")
	} else if err != nil {
		log.Fatal(err)
	}
	db.Exec("COMMIT")
	fmt.Printf("id: %d, name: %s\n", id, name)
}
