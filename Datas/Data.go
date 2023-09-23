package Datas

import (
	funcs "learning-golang/golang-first-api/Functions"
	"learning-golang/golang-first-api/Model"
)

var Todos = []Model.Todo{
	{Item: "clean room", Owner: funcs.Address("raphael olaiyapo"), Completed: false},
	{Item: "read book", Owner: funcs.Address("alade tunji"), Completed: false},
	{Item: "record video", Owner: funcs.Address("aduragbemi adegbite"), Completed: false},
}
