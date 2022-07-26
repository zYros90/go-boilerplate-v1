package v1

import (
	_ "github.com/labstack/echo/v4"
	restapi "github.com/zYros90/protoc-gen-restapi/utils"
)

const TodoSvc_Create_Method = "POST"
const TodoSvc_Create_Path = "/todo/v1"

const TodoSvc_Update_Method = "PUT"
const TodoSvc_Update_Path = "/todo/v1"

const TodoSvc_Get_Method = "GET"
const TodoSvc_Get_Path = "/todo/v1/:todo_id"

const TodoSvc_Delete_Method = "DELETE"
const TodoSvc_Delete_Path = "/todo/v1/:todo_id"

var TodoSvcHTTP []*restapi.ApiAnnotations = []*restapi.ApiAnnotations{

	{
		Method: TodoSvc_Create_Method,
		Path:   TodoSvc_Create_Path,
	},

	{
		Method: TodoSvc_Update_Method,
		Path:   TodoSvc_Update_Path,
	},

	{
		Method: TodoSvc_Get_Method,
		Path:   TodoSvc_Get_Path,
	},

	{
		Method: TodoSvc_Delete_Method,
		Path:   TodoSvc_Delete_Path,
	},
}
