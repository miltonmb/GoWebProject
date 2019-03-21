package main

import (
	//"fmt"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	//"sort"
)

var libro1 = makeLibro("libro1", "Learning Python, 4th Edition", 10, "O'Reilly Media, Inc.")
var libro2 = makeLibro("libro2", "Learning Ruby, 1st Edition", 10, "O'Reilly Media, Inc.")
var libro3 = makeLibro("libro3", "C# 7.0 in a Nutshell", 10, "O'Reilly Media, Inc.")
var libro4 = makeLibro("libro4", "21st Century C, 2nd Edition", 10, "O'Reilly Media, Inc.")
var libro5 = makeLibro("libro5", "Learning JavaScript, 3rd Edition", 10, "O'Reilly Media, Inc.")
var Arr = [5]Libro{libro1, libro2, libro3, libro4, libro5}
var cart = [5]int{0, 0, 0, 0, 0}

type newPage struct {
	Title string
	Name1 string
	Name2 string
	Name3 string
	Name4 string
	Name5 string
}
type adminPage struct {
	Title string
	Name1 string
	Name2 string
	Name3 string
	Name4 string
	Name5 string
}

type Libro struct {
	ISBN    string
	name    string
	inStock int
	editora string
}

func (l Libro) Stock() int {
	return l.inStock
}

func (l Libro) SetStock(stock int) {
	l.inStock = stock
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
	var test = []int{0, 0, 0, 0, 0}
	copy(test, n)
	var temp = test

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
	var test = []int{0, 0, 0, 0, 0}
	copy(test, n)
	var temp = test

	sort.Slice(temp[:], func(i, j int) bool {
		return temp[i] < temp[j]
	})

	var y = []int{}

	if temp[0] == temp[1] && temp[2] == temp[0]*2 {
		y = append(y, temp[0]*2, temp[2], temp[3], temp[4])
	}

	if len(y) < 3 {
		return 0.0
	}

	fmt.Println(y)
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
	p := newPage{Title: "Proyecto 2", Name1: Arr[0].name, Name2: Arr[1].name, Name3: Arr[2].name, Name4: Arr[3].name, Name5: Arr[4].name}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, p)
}
func adminHandler(w http.ResponseWriter, r *http.Request) {
	p := newPage{Title: "Administración", Name1: Arr[0].name, Name2: Arr[1].name, Name3: Arr[2].name, Name4: Arr[3].name, Name5: Arr[4].name}
	t, _ := template.ParseFiles("templates/admin.html")
	t.Execute(w, p)
}

func buttonHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "You called me!")
	amount := r.FormValue("amount")
	id := r.FormValue("id")
	fmt.Fprintln(w, amount)
	fmt.Fprintln(w, id)
	i1, err := strconv.Atoi(amount)
	if err == nil {
		fmt.Println(i1)
	}
	for i := 0; i < 5; i++ {
		if id == Arr[i].ISBN {
			Arr[i].inStock = i1
			fmt.Fprintln(w, Arr[i].inStock)
			break
		}
	}
}
func bookHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "You called me!")
	amount := r.FormValue("amount")
	id := r.FormValue("id")
	fmt.Fprintln(w, amount)
	fmt.Fprintln(w, id)
	i1, err := strconv.Atoi(amount)
	if err == nil {
		fmt.Println(i1)
	}
	for i := 0; i < 5; i++ {
		if id == Arr[i].ISBN {
			if Arr[i].inStock > 0 {
				Arr[i].inStock = Arr[i].inStock - i1
				cart[i] += i1
				fmt.Fprintln(w, Arr[i].inStock)
				break
			}
		}
	}
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/callme", buttonHandler)
	http.HandleFunc("/addLibs", bookHandler)
	http.ListenAndServe(":8000", nil)
}
