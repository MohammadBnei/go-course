package main

import (
	"encoding/json"
	"flag"
	"go-course/task"
	"io"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/go-course?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&task.Task{})

	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)

	commad := flag.String("cmd", "add", "command to execute, Options : [ add, find ]")
	name := flag.String("name", "", "name of the task")
	// id := flag.Int("id", -1, "id of the task to retrieve, if none specified will retrieve all")
	description := flag.String("desc", "", "description of the task")

	flag.Parse()

	switch *commad {
	case "add":
		if *name == "" || *description == "" {
			log.Fatalln("You have to put a name and a description")
		}

		newTask := task.InputTask{
			Name:        *name,
			Description: *description,
		}

		task, err := taskService.Store(newTask)
		if err != nil {
			log.Fatalln(err)
		}

		prettyEncode(task, log.Default().Writer())
	case "find":
		tasks, err := taskService.ListAll()
		if err != nil {
			log.Fatalln(err)
		}
		prettyEncode(tasks, log.Default().Writer())
	default:
		log.Fatalln("Specify add or find command, followed by the correct elements")
	}

}

func prettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}
