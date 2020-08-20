package api

import "github.com/labstack/echo"

func ApiHandler(api *echo.Group) {
	api.POST("/login", LoginRoute)
	api.POST("/register", RegisterRoute)

	api.GET("/getTasks", GetTasksRoute, UserAuth)
	api.POST("/newTask", NewTaskRoute, UserAuth)
	api.PUT("/updateComplete", UpdateCompleteRoute, UserAuth)
	api.DELETE("/deleteTask", DeleteTaskRoute, UserAuth)
}
