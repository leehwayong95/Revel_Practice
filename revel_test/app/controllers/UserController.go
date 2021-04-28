package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"revel_test/app/Model"
	"revel_test/app/Model/Response"
)

type UserController struct {
	*revel.Controller
}

var User = Model.User{}

func (c UserController) GetUserById(id string) revel.Result {
	user := User.GetName(id)
	response := Response.JsonResponse{}
	if user != "" {
		response.Status = true
		response.Data = user
	} else {
		response.Status = false
		response.Error = "Can't find user"
	}
	return c.RenderJSON(response)
}

func (c UserController) AddUser () revel.Result {
	//link parameter in Body, on Post Request
	paramMap := make(map[string]string)
	c.Params.BindJSON(&paramMap)

	fmt.Println("Controller ? New User! > id : " + paramMap["id"] + " / name : " + paramMap["name"])
	User.AddUser(paramMap["id"], paramMap["name"])

	result := Response.JsonResponse{}
	result.Status = true
	defer fmt.Println(result)
	return c.RenderJSON(result)
}

func (c UserController) GetUserList (rang int) revel.Result {
	List := User.GetUserList(rang)
	response := Response.JsonResponse{}
	if List != nil {
		response.Status = true
		response.Data = List
	} else {
		response.Status = false
	}
	return c.RenderJSON(response)
}