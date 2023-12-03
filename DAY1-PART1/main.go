package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	// Abrir el archivo
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	//Leer el archivo linea por linea
	var lines []string

	scanner := bufio.NewScanner(file)

	//Meter cada linea en un array
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	re := regexp.MustCompile("a|b|c|d|e|f|g|h|i|j|k|l|m|n|o|p|q|r|s|t|u|v|w|x|y|z")

	acum := 0
	for i := 0; i < len(lines); i++ {
		//Eliminar el texto, me quedo solo con los numeros
		newTxt := re.ReplaceAllString(lines[i], "")

		//Cojo el primer y ultimo digito
		first := string(newTxt[0])
		lenght := len(newTxt)
		last := string(newTxt[lenght-1])

		//Uno el primer y ultimo dígito y lo paso de string a number
		totalString := first + last
		totalNum, _ := strconv.Atoi(totalString)

		//Sumo el numero anterior, con lo que se lleve acumulado
		acum = acum + totalNum
	}

	fmt.Println("TOTAL")
	fmt.Println(acum)

}
