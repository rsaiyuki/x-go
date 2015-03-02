package db

import (
	"log"
	"testxorm/model"
	"encoding/json"
	"net/http"
)

func GetJson(r *http.Request) string {
	engine := GetEngine()
	//	err1 := engine.CreateTables(model.Book{})
	//	if err1 != nil {
	//		log.Println(err1)
	//	}
	//	var book1 model.Book
	//	book1.Name="book1"
	//	_,err2:=engine.Insert(book1,book1)
	//	if err2!=nil {
	//		log.Println(err2)
	//	}
	books := make([]model.Book, 0, 20)
	r.ParseForm()
	if r.Method=="GET" {
		err5 := engine.Find(&books)
		if err5 != nil {
			log.Println(err5)
		}
	}
	b, _ := json.Marshal(&books)
	log.Println(string(b))
	return string(b)
}
func InsertBook1(r *http.Request)string  {

	mr := new(model.Response)
	book:=new(model.Book)
	r.ParseForm()
	if r.Method=="GET"&&len(r.Form["book"])>0{
		book.Name = r.Form["book"][0]
	}else{
		mr.IsSuccess = false
		mr.Msg = "书名不能为空"
		b, _ := json.Marshal(&mr)
//		io.WriteString(w, string(b))
		return string(b)
	}
	engine := GetEngine()
	count, err2 := engine.Insert(book, book)
	if err2 != nil {
		log.Println(err2)
	}
	if count>0{
		mr.IsSuccess = true
		mr.Msg = "成功"
	}else{
		mr.IsSuccess = false
		mr.Msg = err2.Error()
	}

	b, _ := json.Marshal(&mr)
	log.Println(string(b))
//	io.WriteString(w, string(b))
	return string(b)
}
func InsertBook(book *model.Book) string {
	r := new(model.Response)
	engine := GetEngine()
	count, err2 := engine.Insert(book, book)
	if err2 != nil {
		log.Println(err2)
	}
	if count>0{
		r.IsSuccess = true
		r.Msg = "成功"
	}else{
		r.IsSuccess = false
		r.Msg = err2.Error()
	}

	b, _ := json.Marshal(&r)
	log.Println(string(b))
	return string(b)
}

