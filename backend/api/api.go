package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type ListItem struct {
	Id   string `json:"id"`
	Item string `json:"item"`
	Done bool   `json:"done"`
}

var db *sql.DB
var err error

func SetupPostgres() {
	// when running in Docker
	db, err = sql.Open("postgres", "postgres://postgres:password@postgres/todo?sslmode=disable")

	// when running locally
	// db, err = sql.Open("postgres", "postgres://postgres:password@localhost/todo?sslmode=disable")

	if err != nil {
		fmt.Println(err.Error())
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
	}

	log.Println("connected to postgres")
}

// CRUD: Create Read Update Delete API Format

// List all todo items
func TodoItems(c *gin.Context) {
	// 使用 SELECT Query 获取所有行
	rows, err := db.Query("SELECT * FROM list")
	if err != nil {
		c.Error(err) // 将错误添加到上下文中
		c.Abort()    // 中止当前请求，让中间件处理错误
	}
	// 获取所有行并添加到 items
	items := make([]ListItem, 0)

	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			item := ListItem{}
			if err := rows.Scan(&item.Id, &item.Item, &item.Done); err != nil {
				c.Error(err)
				c.Abort()
			}
			item.Item = strings.TrimSpace(item.Item)
			items = append(items, item)
		}
	}
	// 设置响应数据
	c.Set("response_data", items)
}

type CreateItemEntity struct {
	ItemName string `json:"item"`
}

// Create todo item and add to DB
func CreateTodoItem(c *gin.Context) {
	//打印请求路径
	fmt.Println("Request Path: ", c.Request.URL.Path)
	//
	//request body
	var newItem CreateItemEntity
	if err := c.ShouldBind(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	fmt.Println("New Item: ", newItem.ItemName)
	//db
	if newItem.ItemName == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter an item"})
	} else {
		_, err := db.Query("INSERT INTO list(item, done) VALUES($1, $2);", newItem.ItemName, false)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
			return
		}
	}
	c.Set(("response_data"), newItem)
}

// Update todo item
func UpdateTodoItem(c *gin.Context) {
	var listItem ListItem
	if err := c.ShouldBind(&listItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	} else {
		_, err := db.Query("update list set done = $2 where id = $1", listItem.Id, listItem.Done)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
			return
		}
	}
	c.Set(("response_data"), listItem)
}

// Delete todo item
func DeleteTodoItem(c *gin.Context) {
	var deleteItem ListItem

	if err := c.ShouldBind(&deleteItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	} else {
		var exists int
		err := db.QueryRow("SELECT COUNT(*) FROM list WHERE id=$1;", deleteItem.Id).Scan(&exists)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
			return
		}
		if exists == 0  {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}else {
			_, err := db.Exec("DELETE FROM list WHERE id=$1;", deleteItem.Id)
			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
				return
			}
		}

	}
	 	c.Set(("response_data"), deleteItem)
}

// Add Filter API
