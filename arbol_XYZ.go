package main

import (
	"fmt"
	"strings"
	"strconv"
    "bufio"
    "os"
    "regexp"
	//"math/rand"
)
// arbol binario con valores enteros.
type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Expresion string
	Derecha *Arbol
	Padre *Arbol
}
func InicializarArbol(t *Arbol) *Arbol{
    return &Arbol{nil, "?", "?", nil,nil}
}


func insertar(pos string,t *Arbol, v string, w string) *Arbol {
	if t == nil {
		return &Arbol{nil, v, w, nil, nil}
	}
	if pos == "izq" {
		t.Izquierda = insertar(pos,t.Izquierda, v , w)
		t.Izquierda.Padre=t
		return t
	}
	t.Derecha = insertar(pos,t.Derecha, v, w)
	t.Derecha.Padre=t
	return t
}

func RecorrerInorden (cadena *string,t *Arbol){

	if t == nil {
		return
	}
	RecorrerInorden(cadena,t.Izquierda)
	*cadena=(*cadena+ t.Expresion +" "+ t.Valor+" ")
	//fmt.Print(t.Valor)
    //fmt.Print(" ")
	RecorrerInorden(cadena,t.Derecha)


}

func RecorrerPreorden(t *Arbol) {
	if t == nil {
		return
	}
  fmt.Print(t.Valor)
  fmt.Print(" ")
	RecorrerPreorden(t.Izquierda)
	RecorrerPreorden(t.Derecha)
}

func RecorrerPosorden(t *Arbol) {
	if t == nil {
		return
	}
	RecorrerPosorden(t.Izquierda)
	RecorrerPosorden(t.Derecha)
  fmt.Print(t.Valor)
  fmt.Print(" - ")
}



func guardarEnArbo(t *Arbol, info []string) *Arbol{
    var actual *Arbol
    actual=t
    var exp string
    
    for i := 0; i < len(info); i++ {
        exp=evaluar(info[i])
        val, err := strconv.Atoi(info[i])
        if(err!=nil){
            //fmt.Print("hubo error: ")
            actual.Valor=info[i]
            actual.Expresion=exp
            //actual.Valor=info[i]
            actual.Izquierda=InicializarArbol(actual.Izquierda)
            actual.Izquierda.Padre=actual
            actual.Derecha=InicializarArbol(actual.Derecha)
            actual.Derecha.Padre=actual
            //fmt.Println(actual.Izquierda.Padre.Valor)
            actual=actual.Izquierda

        }else{
            //fmt.Print("no hubo error: ")
                   actual.Valor=info[i]
                    actual.Expresion=exp
                //fmt.Println(actual.Valor)
                //fmt.Println(actual.Padre.Valor)
                actual=actual.Padre
                //fmt.Println("val",actual.Izquierda.Valor)
                //fmt.Println("val",actual.Derecha.Valor)

            for {
                if(actual.Derecha.Valor=="?"){
                  //  fmt.Println("No hay por derecha")

                        actual=actual.Derecha
                        break

                }else{
                    //fmt.Println("Si hay por derecha")
    //                for {
  //                      if((actual.Derecha.Valor!="?")&&(actual!=t)){
//  		                actual=actual.Padre
//                        }else{
                        //    break;
                      //  }
	                //}
	                if(actual!=t){
                        actual=actual.Padre
                    }else{
                        break
                    }

                }
            }
            if(i==len(info)-1){
                break
            }
        }
        if(val==0){

        }
        //fmt.Println("Finalizo ciclo:, Actual:",actual.Valor," padre:",actual.Padre.Valor,"izq: ",
        //actual.Padre.Izquierda.Valor,"der: ",actual.Padre.Derecha.Valor)

	}
	return actual
}



func pruebaLeerOperacion(op string) string{
    var operacion2 = " "
    var po=&operacion2

    var ar *Arbol
    ar=InicializarArbol(ar)
    var operacion=op
    var caracteres = strings.Split(operacion," ")
    ar= guardarEnArbo(ar,caracteres)
    RecorrerInorden(po,ar)
    return operacion2
}






// evalua exprecion

func expresiones(t *Arbol) int{

	if t.Valor == "+" {
		return expresiones(t.Izquierda)+expresiones(t.Derecha)
	}else if t.Valor == "-" {
		return expresiones(t.Izquierda)-expresiones(t.Derecha)
	}else if t.Valor == "*" {
		return expresiones(t.Izquierda)*expresiones(t.Derecha)
	}else if t.Valor == "/"{
		return expresiones(t.Izquierda)/expresiones(t.Derecha)
	}else {
		i,_:=strconv.Atoi(t.Valor)
		return i
	}
}

func evaluar(text string) string{
    
    r, _ := regexp.Compile("^([0-9]+$)")
    s, _ := regexp.Compile("^([A-Z]+$)")
    t, _ := regexp.Compile("^([+]|[-]|[/]|[*])")
    u, _ := regexp.Compile("^([:=])")
   
    var exp string

    if(r.MatchString(text)){
        fmt.Println("Valor: ",text) 
        exp = "Valor: "
      
    }else if(t.MatchString(text)){
        fmt.Println("Operador: ",text) 
        exp = "Operador: "
        
    }else if(u.MatchString(text)){
        fmt.Println("Expresion: ",text) 
        exp = "Expresion: "
    }else if(s.MatchString(text)){
        fmt.Println("Variable: ",text) 
        exp = "Variable: "
    }
    return exp
}



func main() {

	var text string
	var scanner=bufio.NewScanner(os.Stdin)

	fmt.Println("\narbol de prueba\n")
	text="X := - - + 5 3 - 1 2 * 5 - 8 / + 2 16 * + 8 7 - + 1 2 5";
	fmt.Println(text)

	var operacion = text


	var ar *Arbol
	ar=InicializarArbol(ar)

	var caracteres = strings.Split(operacion," ")
	ar= guardarEnArbo(ar,caracteres)


	fmt.Println(" \nEcuacion: \n")
	text=pruebaLeerOperacion(text)
	fmt.Println(text)

	fmt.Println("\nresultado: ",expresiones(ar))


	fmt.Println("\nIngrese en preorden (R - I - D) la operacion con un espacio entre cada caracter o numero.\n")
	scanner.Scan()
	text = scanner.Text()
	operacion=text
	fmt.Println(" \nEcuacion: \n")
  text=pruebaLeerOperacion(text)
  fmt.Println(text)

	ar=InicializarArbol(ar)

	caracteres = strings.Split(operacion," ")
	ar= guardarEnArbo(ar,caracteres)
	fmt.Println("\nresultado: ",expresiones(ar))

// GUARDAR ESTRUCTURA DE DATOS (EXPRESION, "")



}
