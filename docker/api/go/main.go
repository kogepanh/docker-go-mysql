package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"         // ginフレームワーク
	_ "github.com/go-sql-driver/mysql" // mysql用ドライバー
	"github.com/jinzhu/gorm"           // gorm
	"github.com/joho/godotenv"
)

type User struct {
	gorm.Model
	NickName string
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(db:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func dbInit() {
	db := gormConnect()
	defer db.Close() // コネクション解放

	db.AutoMigrate(&User{})
}

func dbInsert(nickname string) {
	db := gormConnect()
	defer db.Close()

	db.Create(&User{NickName: nickname})
}

func dbGetAll() []User {
	db := gormConnect()
	defer db.Close()

	var users []User
	// 取得した情報は引数で与えたモデルに格納される
	db.Find(&users)
	return users
}

func dbGetOneByID(id int) User {
	db := gormConnect()
	defer db.Close()

	var user User
	db.First(&user, id)
	return user
}

func main() {
	// 環境変数ファイルの読み込み
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	dbInit()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "dadsadasdsa",
		})
	})

	router.GET("/users", func(c *gin.Context) {
		users := dbGetAll()
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	router.GET("/users/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		user := dbGetOneByID(id)
		c.JSON(http.StatusOK, gin.H{
			"users": user,
		})
	})

	router.POST("/users/:nickname", func(c *gin.Context) {
		nickname := c.Param("nickname")
		dbInsert(nickname)
	})

	router.Run(":8080")
}
