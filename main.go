package main

import (
	"html/template"
	"net/http"
	 "sort"

)

type newPage struct {
	Title string
}

type Libro struct {
	ISBN    string
	name    string
	inStock int
	editora string
}

type Factura struct {
	ID            string
	nombreCliente string
	Libros        []Libro
	Total         float64
}

func makeFactura(ID string, nombreCliente string, Libros []Libro, Total float64) Factura {
	return Factura{ID,nombreCliente,Libros,Total}
}


func makeLibro(ISBN string, name string, inStock int, editora string) Libro {
	return Libro{ISBN, name, inStock, editora}
}



func bodyFactura(Libros []Libro ) string {
	
	
	
	
	return " "
}
/*
func length_encoding (arr [] string) [] string {

  var counter = 1 ;
  for k := 0; k < len(arr) ; k++ {
   if(k < len(arr)-1 ){ 
    if(arr[k] == arr[k+1] ) {
      counter++      
    }else{
      t := strconv.Itoa(counter)
      z = append(z,"(" +t+"," +arr[k]+")")
      counter = 1
    }
   }  
   }
  return z 
}
*/

func calcularTotal(Libros []Libro) float64 {
	var temp = 0.00

	sort.Slice(Libros[:], func(i, j int) bool {
		return Libros[i].ISBN < Libros[j].ISBN
	})

	var z = []int{}
	
	var counter = 1;
	for i := 0; i < len(Libros); i++ {
		if(Libros[i].ISBN != Libros[i+1].ISBN){
			counter++;
		}else{
			z = append(z,counter)
			counter = 1;
		}
	}
	

	for j := 0; j < len(z); j++ {
		/*	
		z[0] = 3;
		z[1] = 2;
		z[2] = 1;
		z[3] = 4;	
		*/


	}



	return temp
}




func index_handler(w http.ResponseWriter, r *http.Request) {
	p := newPage{Title: "Proyecto 2"}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, p)
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
