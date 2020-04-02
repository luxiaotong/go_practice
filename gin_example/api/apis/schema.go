package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luxiaotong/go_practice/gin_example/api/models"
)

//GetDatabaseList is a function to list all databases
func GetDatabaseList(c *gin.Context) {
	result := models.Databases()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

//GetTableList is a function to list all tables of the database
func GetTableList(c *gin.Context) {
	dbname := c.Query("dbname")
	result := models.Tables(dbname)

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

//GetColumnList is a function to list all columns of the table
func GetColumnList(c *gin.Context) {
	dbname := c.Query("dbname")
	tablename := c.Query("tablename")
	result := models.Columns(dbname, tablename)

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}
