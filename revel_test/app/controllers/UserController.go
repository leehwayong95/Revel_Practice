package controllers

import (
	"github.com/revel/revel"
	"revel_test/app/Model"
	"revel_test/app/Model/Response"
)

type UserController struct {
	*revel.Controller
}

func (c UserController) GetUserById(id string) revel.Result {
	var user = Model.GetName(id)

	response := Response.JsonResponse{}
	if user != "" {
		response.Status = true
		response.Data = user
	} else {
		response.Status = false
		response.Data = nil
	}
	return c.RenderJSON(response)
}

func (c UserController) AddUser (id string, name string) revel.Result {
	Model.AddUser(id, name)
	result := Response.JsonResponse{}
	result.Status = true
	return c.RenderJSON(result)
}

func (c UserController) GetUserList (rang int) revel.Result {
	List := Model.GetUserList(rang)
	response := Response.JsonResponse{}
	if List != nil {
		response.Status = true
		response.Data = List
	} else {
		response.Status = false
	}
	return c.RenderJSON(response)
}