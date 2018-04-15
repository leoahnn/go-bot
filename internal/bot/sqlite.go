package bot

import (
	"database/sql"
	"fmt"

	log "github.com/Sirupsen/logrus"
	// sqlite init needed
	_ "github.com/mattn/go-sqlite3"
)

// SQLite wraps sql for the bot
type SQLite struct {
	db *sql.DB
}

// Init opens/creates bot.db and creates a user table
func (o *SQLite) Init() {
	var err error
	o.db, err = sql.Open("sqlite3", "./bot.db")
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}
	o.createUserTable()
}

// InsertUser inserts a new user
func (o *SQLite) InsertUser(name string) {
	stmt := ` 
		INSERT INTO users (name, kudos)
		VALUES(?,0)
	`
	o.db.Exec(stmt, name)
}

// GetUser gets a user based on their name
func (o *SQLite) GetUser(name string) {
	stmt := `SELECT id FROM users WHERE ? = name`
	res, err := o.db.Query(stmt, name)
	if err != nil {
		log.Errorf("error getting user %v", err)
		return
	}
	var id int
	res.Scan(&id)
	fmt.Println(id)
}

func startPoll() {
	// Start a vote
}

func (o *SQLite) createUserTable() {
	stmt := `
		CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE ON CONFLICT IGNORE,
		kudos INT
		);
	`
	_, err := o.db.Exec(stmt)
	if err != nil {
		log.Errorf("error creating table %v", err)
	}
}

// PlusKudo adds kudos to a given user
func (o *SQLite) PlusKudo(name string) {
	stmt := `UPDATE users SET kudos = kudos + 1 WHERE name = ?`
	_, err := o.db.Exec(stmt, name)
	if err != nil {
		log.Error("could not modify kudos %v", err)
	}
}

// MinusKudo subtracts kudos from a given user
func (o *SQLite) MinusKudo(name string) {
	stmt := `UPDATE users SET kudos = kudos - 1 WHERE name = ?`
	_, err := o.db.Exec(stmt, name)
	if err != nil {
		log.Error("could not modify kudos %v", err)
	}
}

// GetKudos gets a given user's kudos
func (o *SQLite) GetKudos(name string) int {
	var score int
	stmt := `SELECT kudos FROM users WHERE name = ?`
	err := o.db.QueryRow(stmt, name).Scan(&score)
	if err != nil {
		log.Error("could not retrieve kudos")
	}
	return score
}
