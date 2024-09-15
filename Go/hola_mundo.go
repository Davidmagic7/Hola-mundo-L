package main

import (
	"fmt"
	"sort"
)

func main() {
    
    fmt.Println("Hola Mundo")

    unaLista := []string{"e", "a", "f", "c", "b", "d"}
    sort.Strings(unaLista)
    fmt.Println("Orden ascendente:", unaLista)

    unaListaReves := []string{"e", "a", "f", "c", "b", "d"}
    sort.Sort(sort.Reverse(sort.StringSlice(unaListaReves)))
    fmt.Println("Orden descendente:", unaListaReves)
}