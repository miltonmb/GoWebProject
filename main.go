package main

import (
	//"fmt"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	//"sort"
)

type newPage struct {
	Title string
}
type adminPage struct {
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
	return Factura{ID, nombreCliente, Libros, Total}
}

func makeLibro(ISBN string, name string, inStock int, editora string) Libro {
	return Libro{ISBN, name, inStock, editora}
}

func bodyFactura(Libros []Libro) string {

	return " "
}

func calcularTotal(n []int) float64 {
	var temp = n
	sort.Slice(temp[:], func(i, j int) bool {
		return temp[i] < temp[j]
	})
	fmt.Println(temp)
	var total = 0.00
	//Se toma el mayor elemento de la lista para
	var iter = temp[len(temp)-1]
	z := temp
	x := []int{}

	var acum = 0
	for j := 0; j < iter; j++ {

		if z[0] > 0 {
			acum++
			z[0]--
		}

		if z[1] > 0 {
			acum++
			z[1]--
		}

		if z[2] > 0 {
			acum++
			z[2]--
		}

		if z[3] > 0 {
			acum++
			z[3]--
		}

		if z[4] > 0 {
			acum++
			z[4]--
		}
		x = append(x, acum)
		fmt.Println(x)
		acum = 0
	}

	for k := 0; k < len(x); k++ {
		if x[k] == 1 {
			total += 800
		}
		if x[k] == 2 {
			total += 2 * (800 - 800*0.05)
		}
		if x[k] == 3 {
			total += 3 * (800 - 800*0.10)
		}
		if x[k] == 4 {
			total += 4 * (800 - 800*0.20)
		}
		if x[k] == 5 {
			total += 5 * (800 - 800*0.25)
		}
	}
	return total
}

func calcularTotalEsp(n []int) float64 {
	var temp = n
	y := []int{}

	sort.Slice(temp[:], func(i, j int) bool {
		return temp[i] < temp[j]
	})

	if temp[0] == temp[1] && temp[2] == temp[0]*2 {
		y = append(y, temp[0]*2, temp[2], temp[3], temp[4])
	}

	fmt.Println(n)
	var total = 0.00
	//Se toma el mayor elemento de la lista para
	var iter = y[len(y)-1]
	z := y
	x := []int{}

	var acum = 0
	for j := 0; j < iter; j++ {

		if z[0] > 0 {
			acum++
			z[0]--
		}

		if z[1] > 0 {
			acum++
			z[1]--
		}

		if z[2] > 0 {
			acum++
			z[2]--
		}

		if z[3] > 0 {
			acum++
			z[3]--
		}

		x = append(x, acum)
		fmt.Println(x)
		acum = 0
	}

	for k := 0; k < len(x); k++ {
		if x[k] == 1 {
			total += 800
		}
		if x[k] == 2 {
			total += 2 * (800 - 800*0.05)
		}
		if x[k] == 3 {
			total += 3 * (800 - 800*0.10)
		}
		if x[k] == 4 {
			total += 4 * (800 - 800*0.20)
		}
		if x[k] == 5 {
			total += 5 * (800 - 800*0.25)
		}
	}
	return total
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := newPage{Title: "Proyecto 2"}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, p)
}
func adminHandler(w http.ResponseWriter, r *http.Request) {
	p := newPage{Title: "Administrador"}
	t, _ := template.ParseFiles("templates/admin.html")
	t.Execute(w, p)
}

func buttonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You called me!")
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/callme", buttonHandler)
	http.ListenAndServe(":8000", nil)
}
