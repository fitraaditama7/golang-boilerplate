package controllers

import (
	"net/http"

	"../structs"
	"../helper"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		response structs.Response
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		response.Message = "Data Tidak Ditemukan"
		response.Code = http.StatusNotFound
		result = helper.ErrorCustomStatus(response)
		c.JSON(response.Code, result)
	} else {
		response.Message = "success"
		response.Result = person
		result = helper.Responses(response)
		c.JSON(http.StatusOK, result)
	}

}

func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result gin.H
		response structs.Response
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		response.Message = "Belum ada data"
		result = helper.Responses(response)
	} else {
		response.Message = "success"
		response.Results = persons
		result = helper.Responses(response)
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	person.First_Name = first_name
	person.Last_Name = last_name
	idb.DB.Create(&person)

	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")

	var (
		person structs.Person
		newPerson structs.Person
		result gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result= gin.H {
			"result": "Data not found",
		}
	}

	newPerson.First_Name = first_name
	newPerson.Last_Name = last_name
	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H {
			"result": "Successfully updated data",
		}
	}

	c.JSON(http.StatusOK, result)
}


func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "successfully deleted data",
		}
	}

	c.JSON(http.StatusOK, result)
}
