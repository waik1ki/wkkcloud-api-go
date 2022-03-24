package postgresql

func (c *connectDB) SessionPASS(sessionID string) (bool, UserInfo) {
	var user UserInfo

	if sessionID == "" {
		return false, user
	}

	sql_statement := "SELECT * FROM users WHERE sessionid=$1"

	err := c.db.QueryRow(sql_statement, sessionID).Scan(&user.Index, &user.Sessionid, &user.Id, &user.Password, &user.CreateAt)

	return err == nil, user
}
