package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetTasksRoute(c echo.Context) error {
	username := c.Get("username").(string)

	userTasks := []Task{}

	res, err := (*db).Query("SELECT taskid, content, completed FROM tasks WHERE owner=?", username)
	if err != nil {
		log.Println(err)
		return err
	}

	for res.Next() {
		task := Task{}
		err = res.Scan(&task.TaskId, &task.Content, &task.Completed)
		if err != nil {
			return err
		}
		userTasks = append(userTasks, task)
	}

	return c.JSON(http.StatusOK, userTasks)
}

func NewTaskRoute(c echo.Context) error {
	username := c.Get("username").(string)

	task := Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	res, err := (*db).Exec("INSERT INTO tasks(content, owner) VALUES (?, ?)", task.Content, username)

	if err != nil {
		return err
	}
	taskId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"taskId": taskId,
	})
}

func UpdateCompleteRoute(c echo.Context) error {
	username := c.Get("username").(string)

	task := Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	_, err := (*db).Exec(
		"UPDATE tasks SET completed=? WHERE owner=? AND taskid=?",
		task.Completed, username, task.TaskId)

	if err != nil {
		return c.String(http.StatusForbidden, "Action denied")
	}

	return c.String(http.StatusOK, "Updated")

}

func DeleteTaskRoute(c echo.Context) error {
	username := c.Get("username").(string)

	task := Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}

	_, err := (*db).Exec("DELETE FROM tasks WHERE taskid=? AND owner=?", task.TaskId, username)

	if err != nil {
		return c.String(http.StatusForbidden, "Action denied")
	}

	return c.String(http.StatusOK, "Deleted")
}
