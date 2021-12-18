package database

func getAddUserQuery() string {
	return `INSERT INTO users (id, name, login, password) 
			VALUES ($1, $2, $3, $4) 
			ON CONFLICT (id) DO UPDATE 
			SET 
				name = $2, 
				login = $3
				password = $4;`
}

func getRemoveUserQuery() string {
	return `DELETE FROM users 
			WHERE id=$1`
}

func getUserQuery() string {
	return `SELECT id, name, login, password 
			FROM users 
			WHERE id=$1`
}

func getUpdateUserQuery() string {
	return `UPDATE users 
			SET 
				name = $1, 
				login = $2, 
				password = $3 
			WHERE id = $4`
}
