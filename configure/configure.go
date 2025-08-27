package configure

import(
	"gorm.io/gorm"
	"github.com/go-sql-driver/mysql"
)

var (db*gorm.DB)

func Connection ()*gorm.DB{
	// opens the db
	d,err := gorm.open("mysql", )
	if err != nil{
		panic(err)
	}
	db = d
	return
}