package modls

import(
	"db_mysql/configure"
	"gorm.io/gorm"
	
)


var db*gorm.DB

type Class struct{
	gorm.Model
	Name string  `gorm:"" json:"name"`
	lesson string `json:"name"`
	Attendance string `json:"name"`
}

func init (){
	db = configure.Connection()
	db.AutoMigrate(&Class{})
}

func GetClassesDb() []Class {
	var classes []Class
	db.Find(&classes)
	return classes
}

func GetclassDb(Id int64) (*Class ,*gorm.DB){
	var getClass Class
	// gets the specified ID and finds the row of that id
	db := db.Where("ID = ?",Id).Find(&getClass)
	return &getClass, db
}

func (c*Class) CreateClassDb() *Class{
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func DeleteClassDb(Id int64) Class{
	var cass Class
	db.Where("ID=?",Id).Delete(class)
	return class
}