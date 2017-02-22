package main

import (
	"fmt"
	"strings"
	"strconv"
    "bufio"
    "os"
	//"math/rand"
)
// arbol binario con valores enteros.
type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
	Padre *Arbol
}
func InicializarArbol(t *Arbol) *Arbol{
    return &Arbol{nil, "?",nil,nil}
}


func insertar(pos string,t *Arbol, v string) *Arbol {
	if t == nil {
		return &Arbol{nil, v, nil, nil}
	}
	if pos == "izq" {
		t.Izquierda = insertar(pos,t.Izquierda, v)
		t.Izquierda.Padre=t
		return t
	}
	t.Derecha = insertar(pos,t.Derecha, v)
	t.Derecha.Padre=t
	return t
}

func RecorrerInorden (cadena *string,t *Arbol){

	if t == nil {
		return
	}
	RecorrerInorden(cadena,t.Izquierda)
	*cadena=(*cadena+t.Valor+" ")
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

func pruebaLeerInorden(){
        var operacion = " "
        var po=&operacion
        var t *Arbol
		t = insertar("izq",t, "*")
		t = insertar("izq",t, "-")
		t = insertar("izq",t, "-")
		t = insertar("izq",t, "+")
		t = insertar("izq",t, "5")
		t.Izquierda.Izquierda = insertar("der",t.Izquierda.Izquierda, "2")
	    t.Izquierda.Derecha = insertar("der",t.Izquierda.Derecha, "1")
	    t= insertar("der",t,"+")
	    t.Derecha = insertar("der",t.Derecha,"6")
	    t.Derecha = insertar("izq",t.Derecha,"4")
		RecorrerInorden(po,t)
		RecorrerPreorden(t)
		//fmt.Println()
		//fmt.Println(operacion)
}
func guardarEnArbol(t *Arbol, info []string){
    var actual *Arbol
    actual=t
    for i := 0; i < len(info); i++ {
        val, err := strconv.Atoi(info[i])
        if(err!=nil){
            //fmt.Print("hubo error: ")
            actual.Valor=info[i]
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
	return


}


func guardarEnArbo(t *Arbol, info []string) *Arbol{
    var actual *Arbol
    actual=t
    for i := 0; i < len(info); i++ {
        val, err := strconv.Atoi(info[i])
        if(err!=nil){
            //fmt.Print("hubo error: ")
            actual.Valor=info[i]
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
    guardarEnArbol(ar,caracteres)
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



func main() {

	var text string
	var scanner=bufio.NewScanner(os.Stdin)

	fmt.Println("\narbol de prueba\n")
	text="- - + 5 3 - 1 2 * 5 - 8 / + 2 16 * + 8 7 - + 1 2 5";
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





}
