package v1

import (
	_ "github.com/labstack/echo/v4"
	restapi "github.com/zYros90/protoc-gen-restapi/utils"
)

const UserSvc_Create_Method = "POST"
const UserSvc_Create_Path = "/user/v1"

const UserSvc_Update_Method = "PUT"
const UserSvc_Update_Path = "/user/v1"

const UserSvc_Get_Method = "GET"
const UserSvc_Get_Path = "/user/v1"

const UserSvc_Delete_Method = "DELETE"
const UserSvc_Delete_Path = "/user/v1"

var UserSvcHTTP []*restapi.ApiAnnotations = []*restapi.ApiAnnotations{

	{
		Method: UserSvc_Create_Method,
		Path:   UserSvc_Create_Path,
	},

	{
		Method: UserSvc_Update_Method,
		Path:   UserSvc_Update_Path,
	},

	{
		Method: UserSvc_Get_Method,
		Path:   UserSvc_Get_Path,
	},

	{
		Method: UserSvc_Delete_Method,
		Path:   UserSvc_Delete_Path,
	},
}
