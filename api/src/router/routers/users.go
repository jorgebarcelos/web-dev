package routers

import (
	"api/src/controllers"
	"net/http"
)

var usersRouter = []Router{
	{
		URI: "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		AuthRequire: false,
	},

	{
		URI: "/users",
		Method: http.MethodGet,
		Function: controllers.RetrieveUsers,
		AuthRequire: false,
	},

	{
		URI: "/users/{userID}",
		Method: http.MethodGet,
		Function: controllers.RetrieveUser,
		AuthRequire: false,
	},

	{
		URI: "/users/{userID}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		AuthRequire: false,
	},

	{
		URI: "/users/{userID}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		AuthRequire: false,
	},
}