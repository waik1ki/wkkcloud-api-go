package postgresql

import "fmt"

func (c *connectDB) CreateContainer(containerName string, containerId string) bool {

	sql_statement := "INSERT INTO containers (name, id) VALUES ($1, $2);"
	_, err := c.db.Exec(sql_statement, containerName, containerId)

	fmt.Println("Inserted container")

	return err == nil
}

func (c *connectDB) DeleteContainer(containerId string) bool {

	sql_statement := "DELETE FROM containers WHERE id=$1"
	_, err := c.db.Exec(sql_statement, containerId)

	fmt.Println("Deleted container")

	return err == nil
}

func (c *connectDB) GetContainers() []*ContainerInfo {

	containers := []*ContainerInfo{}

	sql_statement := "SELECT * FROM containers"

	rows, err := c.db.Query(sql_statement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var container ContainerInfo
		rows.Scan(&container.Index, &container.Name, &container.Id)
		containers = append(containers, &container)
	}
	return containers
}

func (c *connectDB) FindContainer(containerName string) ContainerInfo {

	var container ContainerInfo

	sql_statement := "SELECT * FROM containers WHERE name=$1"

	err := c.db.QueryRow(sql_statement, containerName).Scan(&container.Index, &container.Name, &container.Id)
	if err != nil {
		panic(err)
	}

	return container
}
