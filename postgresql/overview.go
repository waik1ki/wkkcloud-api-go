package postgresql

func ClassifyType(ftype string) string {
	var category string

	switch ftype {
	case "png", "jpg", "jpeg", "bmp":
		category = "이미지"
	case "pdf", "hwp", "docx", "doc":
		category = "문서"
	case "zip", "egg":
		category = "압축파일"
	}

	return category
}

func Calculator(oldSize int, newSize int, file_cnt int, option string) (int, int) {
	var resultSize int
	var resultFiles int
	if option == "DELETE" {
		resultSize = oldSize - newSize
		resultFiles = file_cnt - 1
	} else if option == "INSERT" {
		resultSize = oldSize + newSize
		resultFiles = file_cnt + 1
	}
	return resultSize, resultFiles
}

func (c *connectDB) UpdateOverview(ftype string, fileSize int64, option string) {
	name := ClassifyType(ftype)

	var category CategoryInfo

	sql_statement := "SELECT * FROM overview WHERE name=$1"

	err := c.db.QueryRow(sql_statement, name).Scan(&category.Index, &category.Name, &category.Size, &category.Count)
	if err != nil {
		panic(err)
	}
	resultSize, resultCount := Calculator(category.Size, int(fileSize), category.Count, option)

	sql_statement = "UPDATE overview SET size=$1 WHERE name=$2"
	_, err = c.db.Exec(sql_statement, resultSize, name)
	if err != nil {
		panic(err)
	}

	sql_statement = "UPDATE overview SET count=$1 WHERE name=$2"
	_, err = c.db.Exec(sql_statement, resultCount, name)
	if err != nil {
		panic(err)
	}
}

func (c *connectDB) GetOverview() []*CategoryInfo {

	data := []*CategoryInfo{}

	sql_statement := "SELECT * FROM overview ORDER BY index ASC"

	rows, err := c.db.Query(sql_statement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var category CategoryInfo
		rows.Scan(&category.Index, &category.Name, &category.Size, &category.Count)
		data = append(data, &category)
	}
	return data
}

func (c *connectDB) InitOverview() {

	overviewCategory := []string{"문서", "이미지", "압축파일", "동영상"}

	sql_statement := "INSERT INTO overview (name, size, count) VALUES ($1, 0, 0);"

	for _, v := range overviewCategory {

		_, err := c.db.Exec(sql_statement, v)
		if err != nil {
			panic(err)
		}

	}
}
