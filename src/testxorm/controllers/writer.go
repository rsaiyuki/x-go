package controllers
import (
	"net/http"
	"io"
	"log"
//	"testxorm/controllers"
)
func WriteIo(path string,f interface{}){
	http.HandleFunc(path,func (w http.ResponseWriter, r *http.Request){
			str:=f.(func(*http.Request)string)(r)
			io.WriteString(w, str)
		})
}
func Run(){
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println(err)
	}
}
//func WriteIoBySelect(path string){
//	http.HandleFunc(path,func (w http.ResponseWriter, r *http.Request){
//			io.WriteString(w, controllers.GetJson())
//		})
//}
