package database

import "database/sql"

func prepare(db *sql.DB) error {
	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS accounts (
			id TEXT PRIMARY KEY,
			user_name TEXT,
			email TEXT,
			hashed_password TEXT,
			image_url TEXT DEFAULT '',
			salt TEXT,
			is_active BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}

	// statement, err = db.Prepare(`
	// 	CREATE TABLE IF NOT EXISTS posts (
	// 		id TEXT PRIMARY KEY,
	// 		user_id TEXT,
	// 		title TEXT,
	// 		content TEXT,
	// 		image_url TEXT,
	// 		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// 		FOREIGN KEY (user_id) REFERENCES users(id)
	// )`)
	// if err != nil {

	// 	return err
	// }
	// if _, err := statement.Exec(); err != nil {

	// 	return err
	// }

	return nil
}
