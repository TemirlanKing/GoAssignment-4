package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID        int
	Name      string
	Completed bool
}

func createTask(db *sql.DB, name string) error {
	_, err := db.Exec("INSERT INTO tasks (name) VALUES (?)", name)
	return err
}

func readTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, name, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func updateTask(db *sql.DB, id int, completed bool) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.Exec("UPDATE tasks SET completed = ? WHERE id = ?", completed, id)
	return err
}

func deleteTask(db *sql.DB, id int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func main() {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/task_management")
	if err != nil {
		fmt.Println("Error connecting to MySQL:", err)
		return
	}
	defer db.Close()

	err = createTask(db, "Complete assignment")
	if err != nil {
		fmt.Println("Error creating task:", err)
		return
	}

	tasks, err := readTasks(db)
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}
	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("%d. %s (Completed: %t)\n", task.ID, task.Name, task.Completed)
	}

	err = updateTask(db, 1, true)
	if err != nil {
		fmt.Println("Error updating task:", err)
		return
	}

	err = deleteTask(db, 1)
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return
	}
}
