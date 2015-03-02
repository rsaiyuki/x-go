package main
import (
	_"github.com/go-sql-driver/mysql"
	"testxorm/controllers"
	"testxorm/db"
//	"testxorm/model"
//	"net/http"
)

func main(){
//	book:=model.Book{
//		Name:"book2",
//	}
//	http.HandleFunc("/insert",controllers.InsertBook1)
	controllers.WriteIo("/book",db.GetJson)
	controllers.WriteIo("/insert",db.InsertBook1)
//	view.WriteIo("/hello",controllers.InsertBook(&book))
//	view.WriteIo("/bye","bye")
	controllers.Run()
}

