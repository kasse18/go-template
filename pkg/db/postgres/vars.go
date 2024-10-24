package postgres

const (
	queryUser        = "SELECT * FROM users WHERE id=$1"
	insertUser       = "INSERT INTO users (id, username) values ($1, $2)"
	updateUser       = "UPDATE users SET balance=$1 WHERE id=$2"
	queryLeaderboard = "SELECT id, username FROM users ORDER BY id DESC"
	queryUsers       = "SELECT username FROM users"
)

const (
	queryInitUsers = `CREATE TABLE IF NOT EXISTS users (
		id bigint NOT NULL,
		username text NOT NULL,
		PRIMARY KEY (id)
	  )`
)
