package mysql

const (
	queryGetUserByUsername = `SELECT id, username, full_name, password, picture FROM users WHERE username = ?`
	queryInsertUser        = `INSERT INTO users (username, password, full_name, picture) VALUES (?,?,?,?)`
	queryGetUserByID       = `SELECT id, username, full_name, password, picture FROM users WHERE id = ?`
	queryUpdateUser        = `UPDATE users SET picture = ?, modified_at = NOW() WHERE id = ?`
)
