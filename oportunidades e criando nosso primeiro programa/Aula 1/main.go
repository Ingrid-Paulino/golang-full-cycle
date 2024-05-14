package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Carro struct {
	Nome   string `json:"nome"` //json é uma tag, ao rodar o projeto o atributo aparecera como o nome colocado na tag
	Modelo string `json:"modelo"`
	Ano    int    `json:"-"` // colocando um "-" na tag json, o campo não é retornado ao rodar o projeto
}

// Adicionando comportamento ao Carro
// Essa func é um Methodo, a diferença entre methodo e função é que o methodo faz referencia para a estrutura que colocada após func
func (c Carro) Andar() {
	fmt.Println("O carro, " + c.Nome + " está andando")
}

func (c Carro) Parar() {
	fmt.Println("O carro, " + c.Nome + " está Parando")
}

// go run main.go -> roda o projeto
func main() {
	//PARTE 1 DA AULA
	// print("Hello, Wold")

	//PARTE 2 DA AULA
	//http.ListenAndServe(":3333", nil) // WEB SERVER
	// go run main.go
	// curl http://localhost:3333 -> comando para subir o servidor web, rodar pelo terminal

	//PARTE 3 DA AULA
	// carro1 := Carro{Nome: "Ford", Modelo: "Mustang", Ano: 1969}
	// carro2 := Carro{Nome: "Chevrolet", Modelo: "Camaro", Ano: 1969}

	// fmt.Println(carro1.Nome)
	// fmt.Println(carro2.Nome)

	// carro1.Andar()
	// carro1.Parar()

	// carro2.Andar()
	// carro2.Parar()

	//PARTE 4 DA AULA
	carro1 := Carro{Nome: "Ford", Modelo: "Mustang", Ano: 1969}

	// funcão anonima
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// convert struct to json
		json.NewEncoder(w).Encode(carro1)
	})

	http.HandleFunc("/home", home)

	http.ListenAndServe(":8080", nil)
}

// Função normal
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Olá, Mundo!")
}
