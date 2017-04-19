package main

import "fmt"
import "regexp"
import "bufio"
import "os"
import "strings"

func evaluar(text string){
    r, _ := regexp.Compile("^([0-9]+$)")
    s, _ := regexp.Compile("^(false|true)")
    t, _ := regexp.Compile("^([<=]|[>=]|[<]|[>]|[=]|[!=])")
    u, _ := regexp.Compile("^([&]|[|])")

    if(r.MatchString(text)){//confirma si es numero
        fmt.Println("Valor Numerico: ",text)   
    }else if(s.MatchString(text)){//confirma que sea palabra
        fmt.Println("valor logico: ",text) 
    }else if(t.MatchString(text)){//confirma que sea palabra
        fmt.Println("comparador logico: ",text) 
    }else if(u.MatchString(text)){//confirma que sea palabra
        fmt.Println("operacion logica: ",text) 
    }
}

func main() {
    var text string
	var scanner=bufio.NewScanner(os.Stdin)
	var caracteres[] string
	scanner.Scan()
    text = scanner.Text()
	caracteres = strings.Split(text," ")
	
	for i:=0;i<len(caracteres);i++{
	   evaluar(caracteres[i])
	    
	}



}
