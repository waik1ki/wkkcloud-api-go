package postgresql

import "fmt"

func (c *connectDB) InsertFile(name string, size int64, ftype string, author string, container string, downloadUrl string) bool {

	sql_statement := "INSERT INTO files (name, size, ftype, author, container, downloadurl, createdAt) VALUES ($1, $2, $3, $4, $5, $6, now());"
	_, err := c.db.Exec(sql_statement, name, size, ftype, author, container, downloadUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted blob")

	return err == nil
}

func (c *connectDB) DeleteFile(name string) bool {

	var file FileInfo

	sql_statement := "SELECT * FROM files WHERE name=$1"

	err := c.db.QueryRow(sql_statement, name).Scan(&file.Name, &file.Size, &file.Ftype, &file.Author, &file.Container, &file.DownloadURL, &file.CreateAt)
	if err != nil {
		panic(err)
	}

	c.UpdateOverview(file.Ftype, file.Size, "DELETE")

	sql_statement = "DELETE FROM files WHERE name=$1"
	_, err = c.db.Exec(sql_statement, name)

	fmt.Println("Deleted blob")

	return err == nil
}

func (c *connectDB) GetRecentUploadFiles() []*FileInfo {

	files := []*FileInfo{}

	sql_statement := "SELECT * FROM files ORDER BY createdAt DESC LIMIT 6;"

	rows, err := c.db.Query(sql_statement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var file FileInfo
		rows.Scan(&file.Name, &file.Size, &file.Ftype, &file.Author, &file.Container, &file.DownloadURL, &file.CreateAt)
		files = append(files, &file)
	}
	return files
}

func (c *connectDB) GetFiles() []*FileInfo {

	files := []*FileInfo{}

	sql_statement := "SELECT * FROM files"

	rows, err := c.db.Query(sql_statement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var file FileInfo
		rows.Scan(&file.Name, &file.Size, &file.Ftype, &file.Author, &file.Container, &file.DownloadURL, &file.CreateAt)
		files = append(files, &file)
	}
	return files
}

func (c *connectDB) GetFileByContainer(containerId string) []*FileInfo {

	files := []*FileInfo{}

	sql_statement := "SELECT * FROM files WHERE container=$1"

	rows, err := c.db.Query(sql_statement, containerId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var file FileInfo
		rows.Scan(&file.Name, &file.Size, &file.Ftype, &file.Author, &file.Container, &file.DownloadURL, &file.CreateAt)
		files = append(files, &file)
	}
	return files
}
