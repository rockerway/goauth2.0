package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const dbUsername string = "root"
const dbPassword string = "password"
const dbName string = "go"
const dbHost string = "172.19.0.2"

// Model struct
type Model struct {
	db *sql.DB
}

// Init is the method that connect to target
func (model *Model) Init() *sql.DB {
	model.db, _ = sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":3306)/"+dbName)
	return model.db
}

// Query is the method that select statement
func (model *Model) Query(statement string) *sql.Rows {
	rows, _ := model.db.Query(statement)
	return rows
}

// Exec is the method that insert statement
func (model *Model) Exec(statement string) int64 {
	var id int64
	result, err := model.db.Exec(statement)
	if err != nil {
		id, _ = result.LastInsertId()
	}
	return id
}

// InitModel to init model
func InitModel() *Model {
	return &Model{db: &sql.DB{}}
}

// Close close
func (model *Model) Close() {
	model.db.Close()
}
