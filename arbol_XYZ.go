package main

import (
	"fmt"
	"strings"
	"strconv"
    "bufio"
    "os"
    "regexp"

)
// arbol binario con valores enteros.
type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Expresion string
	Derecha *Arbol
	Padre *Arbol
}
func InicializarArbol() *Arbol{
    return &Arbol{nil, "?", "?", nil,nil}
}

func RecorrerInorden (t *Arbol){

	if t == nil {
		return
	}
	RecorrerInorden(t.Izquierda)
	//*cadena=(*cadena+ t.Expresion +": "+ t.Valor+" ")
	if(t.Valor != "?"){
	fmt.Println(t.Expresion +": "+ t.Valor+" ")
	}
	RecorrerInorden(t.Derecha)


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

func GuardarTextoEnArbolPreorden(t *Arbol, info []string) *Arbol{
    var actual *Arbol
    var exp string
    actual=t
    for  i := 0;i<len(info);i++{
        exp=evaluar(info[i])
        if(actual.Izquierda==nil){
            actual.Izquierda=InicializarArbol()
            actual.Izquierda.Padre=actual
        }
        if(actual.Derecha==nil){
            actual.Derecha=InicializarArbol()
            actual.Derecha.Padre=actual
        }
        if(exp=="Variable "){
            if(actual.Padre==nil){
                actual.Izquierda.Valor=info[i]
                actual.Izquierda.Expresion=exp
            }else if (actual.Padre.Expresion=="Operador "){
                actual=actual.Padre
                if(actual.Izquierda.Expresion=="?"){
                    actual.Izquierda.Valor=info[i]
                    actual.Izquierda.Expresion=exp
                    actual=actual.Izquierda
                }else{
                    actual.Derecha.Valor=info[i]
                    actual.Derecha.Expresion=exp
                }
                
            }
        }else if(exp=="Expresion "){
            actual.Valor=info[i]
            actual.Expresion=exp
            actual=actual.Derecha
        }else if(exp=="Operador "){
            actual.Valor=info[i]
            actual.Expresion=exp
            actual=actual.Izquierda
        }else if(exp=="Valor "){
            actual.Valor=info[i]
            actual.Expresion=exp
            actual=actual.Padre
            for {
                if(actual.Derecha.Valor=="?"){
                    actual=actual.Derecha
                    break
                }else{
	                if(actual!=t){
                        actual=actual.Padre
                    }else{
                        break
                    }
                }
            }
        }
    }
    return t
}


// evalua exprecion

func expresiones(t *Arbol) int{
    if(t.Valor==":=" && evaluar(t.Izquierda.Valor)=="Variable "){
        return expresiones(t.Derecha)
    }else if t.Valor == "+" {
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
        //fmt.Println("Valor: ",text) 
        exp = "Valor "
      
    }else if(t.MatchString(text)){
        //fmt.Println("Operador: ",text) 
        exp = "Operador "
        
    }else if(u.MatchString(text)){
        //fmt.Println("Expresion: ",text) 
        exp = "Expresion "
    }else if(s.MatchString(text)){
        //fmt.Println("Variable: ",text) 
        exp = "Variable "
    }
    return exp
}



func main() {
    var operacion string
	var scanner=bufio.NewScanner(os.Stdin)
    
	fmt.Println("\narbol de prueba\n")
	operacion="X := - - + 5 3 - 1 2 * 5 - 8 / + 2 16 * + 8 7 - + 1 2 5";
	fmt.Println(operacion)



	var ar *Arbol
	ar=InicializarArbol()


	var caracteres = strings.Split(operacion," ")

    
	ar= GuardarTextoEnArbolPreorden(ar,caracteres)


//	fmt.Println(" \nEcuacion: \n")
//	text=pruebaLeerOperacion(text[5:])
//	fmt.Println(text)


	fmt.Println("\nresultado: ",expresiones(ar))

    var X *Arbol
    var Y *Arbol
    var Z *Arbol
   
    Y=InicializarArbol()
    Z=InicializarArbol()
     X=InicializarArbol()
    
    fmt.Println("\nIngrese 3 arboles en preorden\n")

	fmt.Println("\nArbol 1\n")
	scanner.Scan()
	operacion = scanner.Text()
//	fmt.Println(" \nEcuacion: \n")
 // text=pruebaLeerOperacion(text)
    fmt.Println(operacion)
    caracteres = strings.Split(operacion," ")
	X= GuardarTextoEnArbolPreorden(X,caracteres)
	RecorrerInorden(X)
	fmt.Println("\n")

	fmt.Println("\nArbol 2.\n")
	scanner.Scan()
	operacion = scanner.Text()
	
//	fmt.Println(" \nEcuacion: \n")
 // text=pruebaLeerOperacion(text)
    fmt.Println(operacion)
    caracteres = strings.Split(operacion," ")
	Y= GuardarTextoEnArbolPreorden(Y,caracteres)
	RecorrerInorden(Y)
	fmt.Println("\n")
	
	fmt.Println("\nArbol 3\n")
	scanner.Scan()
	operacion = scanner.Text()

    fmt.Println(operacion)
    caracteres = strings.Split(operacion," ")
	Z= GuardarTextoEnArbolPreorden(Z,caracteres)
	RecorrerInorden(Z)
	fmt.Println("\n")
		

	
	if(Z.Derecha.Izquierda.Valor==X.Izquierda.Valor){
	    Z.Derecha.Izquierda=X
	}
	if(Z.Derecha.Derecha.Valor==X.Izquierda.Valor){
	    Z.Derecha.Derecha=X

	}
    if(Z.Derecha.Izquierda.Valor==Y.Izquierda.Valor){
	    Z.Derecha.Izquierda=Y
	}

	if(Z.Derecha.Derecha.Valor==Y.Izquierda.Valor){
	    Z.Derecha.Derecha=Y
	}
    
	

	fmt.Println("\nResultados\n\n"+X.Izquierda.Valor+X.Valor,expresiones(X))
	fmt.Println("\n"+Y.Izquierda.Valor+Y.Valor,expresiones(Y))
	fmt.Println("\n"+Z.Izquierda.Valor+Z.Valor,expresiones(Z))
	

}
