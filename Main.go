package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

const CONSTANTE_GLOBAL float32 = 3.141592
const PORCENTAJE_IMPUESTO_1 float32 = 0.1
const PORCENTAJE_IMPUESTO_2 float32 = 0.15
const PORCENTAJE_IMPUESTO_3 float32 = 0.2
const PORCENTAJE_IMPUESTO_4 float32 = 0.25

type pagarImpuestos func(salario float32) float32

func main() {

	//Variables
	fmt.Println("\n*****************")
	fmt.Println("* Variables")
	fmt.Println("*****************")

	var i int = 10
	var n int = 5
	fmt.Println(i, "%", n, "=", i%n)

	var hola string = "Hola"
	var mundo string = " mundo!!"
	fmt.Println(hola + mundo)

	var numero int = 100
	var frase string = "Capacidad de disco duro: "
	frase += strconv.Itoa(numero) + "%"
	fmt.Println(frase)

	//Funciones
	fmt.Println("\n*****************")
	fmt.Println("* Funciones")
	fmt.Println("*****************")

	callFunction()
	parametros(i, n)

	//Arrays
	fmt.Println("\n*****************")
	fmt.Println("* Arrays")
	fmt.Println("*****************")

	var numeros = []int{1, 2, 3}
	fmt.Println(numeros)
	numeros = append(numeros, 4, 5)
	fmt.Println(numeros)
	fmt.Println(numeros[0])

	//Loops
	fmt.Println("\n*****************")
	fmt.Println("* Loops")
	fmt.Println("*****************")

	fmt.Println("Loop for")
	for count := 0; count < len(numeros); count++ {
		fmt.Println(numeros[count])
	}

	fmt.Println("Loop while")
	count := 0
	for count < len(numeros) {
		fmt.Println(numeros[count])
		count++
	}

	//Lists
	fmt.Println("\n*****************")
	fmt.Println("* Lists")
	fmt.Println("*****************")

	var lista list.List
	lista.PushBack("e1")
	lista.PushBack("e2")
	lista.PushBack("e3")
	lista.PushBack("e4")
	lista.PushFront("e0")
	pointerToFront := lista.Front()
	for pointerToFront != nil {
		fmt.Println(pointerToFront.Value)
		pointerToFront = pointerToFront.Next()
	}

	//Mapa
	fmt.Println("\n*****************")
	fmt.Println("* Mapa")
	fmt.Println("*****************")

	var mapa = make(map[string]string)
	mapa["e1"] = "124"
	mapa["e2"] = "/etc/openvpn/config"
	mapa["e3"] = "Error: archivo no encontrado"
	fmt.Println(mapa["e1"])
	fmt.Println(mapa["e2"])
	fmt.Println(mapa["e3"])

	//Json
	fmt.Println("\n*****************")
	fmt.Println("* JSON")
	fmt.Println("*****************")

	var jsonText string = ""
	jsonText += "{\n"
	jsonText += "  \"e1\": \"124\",\n"
	jsonText += "  \"e2\": \"/etc/openvpn/config\",\n"
	jsonText += "  \"e3\": \"Error: archivo no encontrado\"\n"
	jsonText += "}"
	fmt.Println("\n" + jsonText + "\n")
	my_json, err := json.Marshal(jsonText)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(my_json))

	jsonText = "{ \"e1\": \"124\", \"e2\": \"/etc/openvpn/config\", \"e3\": \"Error: archivo no encontrado\"}"
	var json_map map[string]interface{}
	json.Unmarshal([]byte(jsonText), &json_map)
	fmt.Println(json_map["e1"])
	fmt.Println(json_map["e2"])
	fmt.Println(json_map["e3"])

	//Struct
	fmt.Println("\n*****************")
	fmt.Println("* Struct")
	fmt.Println("*****************")
	type Coche struct {
		color     string
		velocidad float32
		matricula string
	}

	var coche = Coche{"negro", 163.4, "48124CVF"}
	fmt.Println(coche)
	coche.color = "blanco"
	fmt.Println(coche)

	type Hacienda struct {
		ingresos float32
		salario  pagarImpuestos
	}

	var h = Hacienda{
		ingresos: 35000.0,
		salario: func(ingresos float32) float32 {
			var franja1 float32 = 15000
			var franja2 float32 = 25000
			var franja3 float32 = 35000
			var out float32 = 0
			var empty bool

			//franja 1
			if (ingresos - franja1) > 0 {
				out += franja1 - (franja1 * PORCENTAJE_IMPUESTO_1)
				ingresos -= franja1
			} else {
				out = (ingresos * PORCENTAJE_IMPUESTO_1) - ingresos
				empty = true
			}

			//franja 2
			if ((ingresos - franja2) > 0) && (!empty) {
				out += franja2 - (franja2 * PORCENTAJE_IMPUESTO_2)
				ingresos -= franja2
			} else if !empty {
				out += ingresos - (ingresos * PORCENTAJE_IMPUESTO_2)
				empty = true
			}

			//franja 3
			if ((ingresos - franja3) > 0) && (!empty) {
				out += franja3 - (franja3 * PORCENTAJE_IMPUESTO_3)
				ingresos -= franja3
			} else if !empty {
				out += ingresos - (ingresos * PORCENTAJE_IMPUESTO_3)
				empty = true
			}

			//franja 4
			if !empty {
				out += ingresos - (ingresos * PORCENTAJE_IMPUESTO_4)
			}

			return out
		},
	}
	fmt.Println(h.salario(h.ingresos))

	type Hacienda2 struct {
		ingresos float32
		salario  func(float32) float32
	}
	var h2 = Hacienda2{
		ingresos: 35000.0,
		//Puntero a funcion
		salario: pagarImpuestos2,
	}
	fmt.Println(h2.salario(h2.ingresos))

	//Local package
	fmt.Println("\n*****************")
	fmt.Println("* Local package")
	fmt.Println("*****************")
	fmt.Println("Agregar el archivo al package main y ejecutar con \"golang main.go mysql.go\"")

	//MySQL
	fmt.Println("\n*****************")
	fmt.Println("* MySQL")
	fmt.Println("*****************")
	var mysql = mysql_crear("root", "dos+3=CINCO", "192.168.1.33", "30306", "golang")
	mysql.conexion(mysql)

}

func callFunction() {
	fmt.Println("Hola desde la función")
}

func parametros(numero1 int, numero2 int) {
	fmt.Println(numero1, "+", numero2, "=", numero1+numero2)
}

func pagarImpuestos2(ingresos float32) float32 {
	var franja1 float32 = 15000
	var franja2 float32 = 25000
	var franja3 float32 = 35000
	var out float32 = 0
	var empty bool

	//franja 1
	if (ingresos - franja1) > 0 {
		out += franja1 - (franja1 * PORCENTAJE_IMPUESTO_1)
		ingresos -= franja1
	} else {
		out = (ingresos * PORCENTAJE_IMPUESTO_1) - ingresos
		empty = true
	}

	//franja 2
	if ((ingresos - franja2) > 0) && (!empty) {
		out += franja2 - (franja2 * PORCENTAJE_IMPUESTO_2)
		ingresos -= franja2
	} else if !empty {
		out += ingresos - (ingresos * PORCENTAJE_IMPUESTO_2)
		empty = true
	}

	//franja 3
	if ((ingresos - franja3) > 0) && (!empty) {
		out += franja3 - (franja3 * PORCENTAJE_IMPUESTO_3)
		ingresos -= franja3
	} else if !empty {
		out += ingresos - (ingresos * PORCENTAJE_IMPUESTO_3)
		empty = true
	}

	//franja 4
	if !empty {
		out += ingresos - (ingresos * PORCENTAJE_IMPUESTO_4)
	}

	return out
}
