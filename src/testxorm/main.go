package main
import (
	_"github.com/go-sql-driver/mysql"
	"testxorm/controllers"
	"testxorm/db"
)

func main(){
	db.InitDBEngine()
	controllers.WriteIo("/book",db.GetJson)
	controllers.WriteIo("/",db.InsertBook1)
	controllers.Run()
}

