package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"ts-server/db"
	"ts-server/models"
)

func main() {

	db.PgConnect()

	db.Pgdb.AutoMigrate(
		&models.Exam{},
		&models.Answer{},
		&models.Question{})

	router := NewRouter()
	var port string
	if port = "4000"; os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	fmt.Println("Server will start at: " + port)
	router.Run(":" + port)

}

func OrmChaining() {

	var questions []models.Question
	db.Pgdb.
		Where("index > ?", 2).
		Where("index < ?", 10).
		Find(&questions)

	fmt.Print(questions)
}

func OrmScopes() {

	var questions []models.Question
	db.Pgdb.
		Scopes(FromIndex5, ToIndex10).
		Find(&questions)

	fmt.Println(questions)
}

func OrmScopesWithArrayParams() {

	var questions []models.Question
	db.Pgdb.
		Scopes(IndexOf([]int{2, 3, 4})).
		Find(&questions)

	fmt.Println(questions)
}

func OrmPreload() {

	var question models.Question
	db.Pgdb.Preload("Answers").
		Where("index = ?", 2).
		Find(&question)

	fmt.Println(question)
}

func IndexOf(indexes []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("index in (?)", indexes)
	}
}

func FromIndex5(db *gorm.DB) *gorm.DB {
	return db.Where("index > ?", 5)
}

func ToIndex10(db *gorm.DB) *gorm.DB {
	return db.Where("index < ?", 10)
}
