package main
import ("net/http"
		"html/template"
)

type newPage struct{
	Title string
}

func index_handler(w http.ResponseWriter, r *http.Request ){
	p := newPage{Title: "Proyecto 2"}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w,p)
}
func main(){
	fs := http.FileServer(http.Dir("static"))
  	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000",nil)
}
