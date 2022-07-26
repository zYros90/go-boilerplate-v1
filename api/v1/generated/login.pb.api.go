package v1

import (
	_ "github.com/labstack/echo/v4"
	restapi "github.com/zYros90/protoc-gen-restapi/utils"
)

const LoginSvc_Login_Method = "POST"
const LoginSvc_Login_Path = "/login/v1"

var LoginSvcHTTP []*restapi.ApiAnnotations = []*restapi.ApiAnnotations{

	{
		Method: LoginSvc_Login_Method,
		Path:   LoginSvc_Login_Path,
	},
}
