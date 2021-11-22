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
	Id_persona               int    `json:"id_persona"`
	Id_eess                  string `json:"id_eess"`
	Fecha_ingreso            string `json:"fecha_ingreso"`
	Hora_ingreso             string `json:"hora_ingreso"`
	Es_recuperado            string `json:"es_recuperado"`
	Fecha_alta               string `json:"fecha_alta"`
	Es_recuperado_voluntario string `json:"es_recuperado_voluntario"`
	Fecha_alta_voluntaria    string `json:"fecha_alta_voluntaria"`
	Es_fallecido             string `json:"es_fallecido"`
	Fecha_fallecido          string `json:"fecha_fallecido"`
	Es_referido              string `json:"es_referido"`
	Fecha_referido           string `json:"fecha_referido"`
	Eess_destino_id          string `json:"eess_destino_id"`
}

func lineToStruc(lines [][]string) {
	// Recorre líneas y conviértete en objeto
	for _, line := range lines {
		id_persona, _ := strconv.Atoi(strings.TrimSpace(line[0]))

		atencions = append(atencions, Atencion{
			Id_persona:               id_persona,
			Id_eess:                  strings.TrimSpace(line[1]),
			Fecha_ingreso:            strings.TrimSpace(line[2]),
			Hora_ingreso:             strings.TrimSpace(line[3]),
			Es_recuperado:            strings.TrimSpace(line[4]),
			Fecha_alta:               strings.TrimSpace(line[5]),
			Es_recuperado_voluntario: strings.TrimSpace(line[6]),
			Fecha_alta_voluntaria:    strings.TrimSpace(line[7]),
			Es_fallecido:             strings.TrimSpace(line[8]),
			Fecha_fallecido:          strings.TrimSpace(line[9]),
			Es_referido:              strings.TrimSpace(line[10]),
			Fecha_referido:           strings.TrimSpace(line[11]),
			Eess_destino_id:          strings.TrimSpace(line[12]),
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

//Lee archivo de url
func readFileUrl(filePathUrl string) ([][]string, error) {
	// Abrir archivo CSV
	resp, err := http.Get(filePathUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
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
		idpersona, _ := strconv.Atoi(params["id"])
		if item.Id_persona == idpersona {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Atencion{})
}

func createAtencion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var atencion Atencion
	_ = json.NewDecoder(r.Body).Decode(&atencion)
	atencions = append(atencions, atencion)
	json.NewEncoder(w).Encode(atencion)
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
	r.HandleFunc("/atencions", createAtencion).Methods("POST")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Start server
	port := ":8000"
	fmt.Println("Escuchando en " + port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(r)))
}
