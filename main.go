package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func welcomeApi(c *gin.Context) {
	c.HTML(200, "welcomeApi.html", nil)
}

func crateRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("pages/*.html")
	// Simple group: v1
	v1 := router.Group("/api")
	{
		v1.GET("/", welcomeApi)

	}

	// Simple group: v2
	v2 := router.Group("/")
	{
		v2.GET("/login", func(c *gin.Context) {
			c.String(200, "succes :)")
		})

	}
}

func setupDatabase() *sql.DB {
	database, error := sql.Open("postgres", "user=kantemir password=kantemir dbname=projects sslmode=disable")

	if error != nil {
		log.Fatal(error)
	}

	return database
}

type Project struct {
	Id          int
	Name        string
	Gif_path    string
	Description string
	Homepage    string
	Repository  string
	Image_paths []uint8
}

func getAllProjects(db *sql.DB) (projects []Project, err error) {
	rows, err := db.Query(`SELECT * FROM projects`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	/* var projects []Project */
	for rows.Next() {
		var (
			id          int
			name        string
			gif_path    string
			description string
			homepage    string
			repository  string
			image_paths []uint8
		)

		if err := rows.Scan(&id, &name, &gif_path, &description, &homepage, &repository, &image_paths); err != nil {
			return nil, err
		}
		projects = append(projects, Project{Id: id, Name: name, Gif_path: gif_path, Description: description, Homepage: homepage, Repository: repository, Image_paths: image_paths})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func main() {
	router := gin.Default()

	crateRoutes(router)

	db := setupDatabase()

	/* data, err :=  */
	projects, err := getAllProjects(db)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(projects); i++ {
		fmt.Printf("%d has the name %s\n", projects[i].Id, projects[i].Name)
	}
	/* if err == nil {
		fmt.Print(data)
	} */

	router.Run(":5000")
}
