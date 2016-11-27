package main

import (
    "fmt"
    "github.com/jambola2/bip38"
	"bufio"
	"os"
	"strings"
	"log"
    "runtime"
)

func main() {
    //Modify as needed
    fileName := "completeFilev3-2-78600.txt"
    passPhrase := "6PfTUmvB7PWj5KwUjoJeU3S51Fu1AU86URJfSetnotzZ8BBrCbghG7i84C"
    //How frequently you want to be updated on tries completed
    triesPrintFreq := 100 
    
    //Maximize processor usage
    runtime.GOMAXPROCS(runtime.NumCPU())
	tries := 0
    pwd, _ := os.Getwd()
    f, _ := os.Open(pwd + "/" + fileName)
    bf := bufio.NewReader(f)
    
    fullLine := ""

    for {
    
        line, isPrefix, err := bf.ReadLine()
        fullLine += string(line[:])
        
        if !isPrefix || err != nil{
            break
        }
    }	
    //Split the line on commas.
    parts := strings.Split(fullLine, ",")
    
    //Loop over the parts from the string.
	for i := range parts {
		privKey := bip38.DecryptWithPassphrase(passPhrase,parts[i])
		//fmt.Println(parts[i])
        tries = tries + 1
		if tries%triesPrintFreq == 0{
            fmt.Println(tries + "tries completed")
        }
        if privKey != "" {
            fmt.Println(parts[i])
            log.Fatal(privKey)
        }
    }
}