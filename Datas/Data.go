package Datas

import (
	funcs "learning-golang/golang-first-api/Functions"
	"learning-golang/golang-first-api/Model"
)

var Todos = []Model.Todo{
	{ID: 1, Item: "clean room", Owner: funcs.Address("raphael olaiyapo"), Completed: false},
	{ID: 2, Item: "read book", Owner: funcs.Address("alade tunji"), Completed: false},
	{ID: 3, Item: "record video", Owner: funcs.Address("aduragbemi adegbite"), Completed: false},
}
