package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic("o arquivo nao foi criado")
	}
	tamanho, err := f.Write([]byte("Bora amigos"))
	// tamanho, err := f.WriteString("Bora amigos") // escreve apenas textos
	if err != nil {
		panic("o arquivo nao foi criado")
	}
	fmt.Printf("O arquivo foi escrito com sucesso \n tamanho: %d bytes\n", tamanho)

	f.Close()
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic("falha ao abrir o arquivo")
	}

	fmt.Println(string(arquivo))

	// usando buffer

	file, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
