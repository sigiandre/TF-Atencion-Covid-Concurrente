package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var atencions []Atencion

type Atencion struct {
	id_persona               int    `json:"id_persona"`
	id_eess                  string `json:"id_eess"`
	fecha_ingreso            string `json:"fecha_ingreso"`
	hora_ingreso             string `json:"hora_ingreso"`
	es_recuperado            string `json:"es_recuperado"`
	fecha_alta               string `json:"fecha_alta"`
	es_recuperado_voluntario string `json:"es_recuperado_voluntario"`
	fecha_alta_voluntaria    string `json:"fecha_alta_voluntaria"`
	es_fallecido             string `json:"es_fallecido"`
	fecha_fallecido          string `json:"fecha_fallecido"`
	es_referido              string `json:"es_referido"`
	fecha_referido           string `json:"fecha_referido"`
	eess_destino_id          string `json:"eess_destino_id"`
}

func lineToStruc(lines [][]string) {
	// Recorre líneas y conviértete en objeto
	for _, line := range lines {
		id_persona, _ := strconv.Atoi(strings.TrimSpace(line[0]))

		atencions = append(atencions, Atencion{
			id_persona:               id_persona,
			id_eess:                  strings.TrimSpace(line[1]),
			fecha_ingreso:            strings.TrimSpace(line[2]),
			hora_ingreso:             strings.TrimSpace(line[3]),
			es_recuperado:            strings.TrimSpace(line[4]),
			fecha_alta:               strings.TrimSpace(line[5]),
			es_recuperado_voluntario: strings.TrimSpace(line[6]),
			fecha_alta_voluntaria:    strings.TrimSpace(line[7]),
			es_fallecido:             strings.TrimSpace(line[8]),
			fecha_fallecido:          strings.TrimSpace(line[9]),
			es_referido:              strings.TrimSpace(line[10]),
			fecha_referido:           strings.TrimSpace(line[11]),
			eess_destino_id:          strings.TrimSpace(line[12]),
		})
	}
}

//variables globales
var eschucha_funcion bool
var remotehost string
var n, min, valorUsuario int

func enviar(num int) { //enviar el numero mayor al host remoto
	conn, _ := net.Dial("tcp", remotehost)
	defer conn.Close()
	//envio el número
	fmt.Fprintf(conn, "%d\n", num)

}

func enviar_Principal(num int) { //enviar el numero mayor al host remoto
	conn, _ := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	//envio el número
	fmt.Fprintf(conn, "%d\n", num)

}

func manejador_respueta(conn net.Conn) bool {
	defer conn.Close()
	eschucha_funcion = false
	bufferIn := bufio.NewReader(conn)
	numStr, _ := bufferIn.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	numero, _ := strconv.Atoi(numStr)
	strNumero := strconv.Itoa(numero)
	if strNumero[1] == 49 {
		return true
	} else {
		return false
	}
}

func readFileUrl(filePathUrl string) ([][]string, error) {
	// Abrir archivo CSV
	f, err := http.Get(filePathUrl)
	if err != nil {
		return [][]string{}, err
	}

	// Leer archivo en una variable
	lines, err := csv.NewReader(f.Body).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	defer f.Body.Close()
	return lines, nil
}

// Get all Atencion
func getAtencions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atencions)
}

// Get single Atencion
func getAtencion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range atencions {
		idpersona, _ := strconv.Atoi(params["id_persona"])
		if item.id_persona == idpersona {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Atencion{})
}

func main() {
	bufferIn := bufio.NewReader(os.Stdin)

	filePathUrl := "https://raw.githubusercontent.com/sigiandre/TF-Atencion-Covid-Concurrente/master/dataset/TB_ATEN_COVID19.csv"
	lines, err := readFileUrl(filePathUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Leyo archivos")
	lineToStruc(lines)
	fmt.Println("Parseo Archivos")

	//tipo de nodo
	log.Print("Ingrese el tipo de nodo (i:inicio -n:intermedio - f:final): ")
	tipo, _ := bufferIn.ReadString('\n')
	tipo = strings.TrimSpace(tipo)

	r := mux.NewRouter()

	r.HandleFunc("/atencions", getAtencions).Methods("GET")
	r.HandleFunc("/atencions/{id}", getAtencion).Methods("GET")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Start server
	port := ":8000"
	fmt.Println("Escuchando en " + port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(r)))
}
