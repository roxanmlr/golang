package main
// le package main est le package qui contient la fonction main du programme
// Go lance le programme à partir de cette fonction main

import "os"

func main(){
	os.Stdout.Write([]byte{'c'})
}	
