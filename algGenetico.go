package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Genotipo struct {
	genes        []float64
	predominante []float64
}

type Entrada struct {
	//eess           int
	FechaIngreso   int
	DiasRecuperado int
	DiasFallecido  int
	DiasReferido   int
}

type Salida struct {
	DiasRecuperado int
	DiasFallecido  int
	DiasReferido   int
}

/**
 */
func InitGenetico(atenciones []Atencion) []Genotipo {
	// Cantidad de genes
	var genesCantidad = 10

	// Iniciar genes
	genotipos := initGenotipos(genesCantidad, 3*3)

	// Ordenar genes por euristica
	var heuristicas []float64
	genotipos, heuristicas = ordPorHeuristica(genotipos, atenciones)

	// Imprimir mejores 5 genes con su heuristica
	for i := 0; i < 3; i++ {
		Info("Mejor gen ["+strconv.Itoa(i)+"]:", genotipos[i].genes, " | heuristica: ", heuristicas[i])
	}

	// Retornar genes
	return genotipos
}

func SiguienteGeneracion(genotipos []Genotipo, atenciones []Atencion) []Genotipo {
	// Cruzar genes
	genesCruzados := cruceGenes(genotipos, 10)
	genesMutados := mutarGenes(genotipos, 2.0)
	// genesMutados = append(genesMutados, mutarGenes(genotipos, 2.0)...)

	// Append genes
	genotipos = append(genotipos, genesCruzados...)
	genotipos = append(genotipos, genesMutados...)

	// Ordenar
	var heuristicas []float64
	genotipos, heuristicas = ordPorHeuristica(genotipos, atenciones)

	// Imprimir mejores 3 genes con su heuristica
	for i := 0; i < 3; i++ {
		Info("Mejor gen ["+strconv.Itoa(i)+"]:", genotipos[i].genes, " | heuristica: ", heuristicas[i])
	}

	// Retornar nuevos 10 mejores genes
	return genotipos[0:10]
}

func ObtenerV(atencion Atencion) float64 {
	// From 1 ene 2020
	from := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	to := FechaStr2Time(atencion.Fecha_ingreso)
	return float64(to.Sub(from).Hours() / 24)
}

func Ejecutar(v float64, genotipo Genotipo) Salida {
	// Formula "(x*g1+g2)^g3"
	// x <- variable
	// g <- gen
	g := genotipo.genes
	return Salida{
		DiasRecuperado: int(math.Round(math.Pow((v*g[0] + g[1]), g[2]))),
		DiasFallecido:  int(math.Round(math.Pow((v*g[3] + g[4]), g[5]))),
		DiasReferido:   int(math.Round(math.Pow((v*g[6] + g[7]), g[8]))),
	}
	// Info("ABC EJE", v, int(math.Round((v*g[0]+g[1]))-v))
	// return Salida{
	// 	DiasRecuperado: int(math.Round((v*g[0] + g[1])) - v),
	// 	DiasFallecido:  int(math.Round((v*g[3] + g[4])) - v),
	// 	DiasReferido:   int(math.Round((v*g[6] + g[7])) - v),
	// }
}

func EjecutarCsv(atencion Atencion) Salida {
	// Obtener fecha de ingreso
	fechaIngreso := FechaStr2Time(atencion.Fecha_ingreso)

	// Si esta recuperado
	var diasRecuperado int = -1
	if atencion.Es_recuperado == "1" {
		fechaNueva := FechaStr2Time(atencion.Fecha_alta)
		diasRecuperado = int(fechaIngreso.Sub(fechaNueva).Hours() / 24)
	}

	// Si esta fallecido
	var diasFallecido int = -1
	if atencion.Es_recuperado == "1" {
		fechaNueva := FechaStr2Time(atencion.Fecha_fallecido)
		diasFallecido = int(fechaIngreso.Sub(fechaNueva).Hours() / 24)
	}

	// Si esta referido
	var diasReferido int = -1
	if atencion.Es_recuperado == "1" {
		fechaNueva := FechaStr2Time(atencion.Fecha_referido)
		diasReferido = int(fechaIngreso.Sub(fechaNueva).Hours() / 24)
	}

	// Retornar salida
	return Salida{
		DiasRecuperado: diasRecuperado,
		DiasFallecido:  diasFallecido,
		DiasReferido:   diasReferido,
	}
}

func FechaStr2Time(fecha string) time.Time {
	var year int
	var month time.Month
	var day int
	fmt.Sscanf(fecha, "%d/%d/%d", &day, &month, &year)
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func crearGenVacio(longitud int) Genotipo {
	// Crear genes y predominancia
	genes := make([]float64, longitud)
	predominancia := make([]float64, longitud)

	// Init vals
	for idGen := 0; idGen < longitud; idGen++ {
		genes[idGen] = 1
		predominancia[idGen] = 1
	}

	// Retornar gen
	return Genotipo{
		genes:        genes,
		predominante: make([]float64, longitud),
	}
}

func initGenotipos(cantidad int, longitud int) []Genotipo {

	// Crear "n" genotipos
	genotipos := make([]Genotipo, cantidad)

	// Para cada genotipo
	for id := 0; id < cantidad; id++ {
		// Crear genes y predominancia
		genes := make([]float64, longitud)
		predominancia := make([]float64, longitud)

		// Crear "l" genes
		for idGen := 0; idGen < longitud; idGen++ {
			genes[idGen] = rand.Float64()
			predominancia[idGen] = 1
		}

		// Asignar genotipo
		genotipos[id] = Genotipo{
			genes:        genes,
			predominante: predominancia,
		}
	}

	// Retornar genotipos
	return genotipos
}

func cruceGenes(genotipos []Genotipo, nCruces int) []Genotipo {
	// Obtener los 3 mejores
	nuevosGenes := genotipos[0:3]

	// Obtener longitud de los genes
	lenGenes := len(genotipos[0].genes)

	// Generar aleatoriamente 10 de los 3 mejores
	for i := 0; i < 10; i++ {
		// Nuevo genotipo
		nuevoGenotipo := crearGenVacio(lenGenes)

		// Seleccionar el valor del cruce entre los 3 primeros
		for iGen := 0; iGen < lenGenes; iGen++ {
			nuevoGenotipo.genes[iGen] = nuevosGenes[rand.Intn(3)].genes[iGen]
		}

		// Añadir nuevo genotipo
		nuevosGenes = append(nuevosGenes, nuevoGenotipo)
	}

	// Generar aleatoriamente nCruces de todos los genes
	cantGenes := len(genotipos)
	for i := 0; i < nCruces; i++ {
		// Nuevo genotipo
		nuevoGenotipo := crearGenVacio(lenGenes)

		// Seleccionar el valor del cruce entre los 3 primeros
		for iGen := 0; iGen < lenGenes; iGen++ {
			nuevoGenotipo.genes[iGen] = nuevosGenes[rand.Intn(cantGenes)].genes[iGen]
		}

		// Añadir nuevo genotipo
		nuevosGenes = append(nuevosGenes, nuevoGenotipo)
	}

	// Retornar nuevo set de genes
	return nuevosGenes
}

func mutarGenes(genotipos []Genotipo, probMutacion float64) []Genotipo {
	var nuevosGenes []Genotipo

	// Obtener longitud de los genes
	lenGenes := len(genotipos[0].genes)

	// Mutar genes
	for iGenotipo := 0; iGenotipo < len(genotipos); iGenotipo++ {
		nuevoGenotipo := crearGenVacio(lenGenes)

		// Mutar cada gen
		for iGen := 0; iGen < lenGenes; iGen++ {
			// Setear valor inicial
			nuevoGenotipo.genes[iGen] = float64(genotipos[iGenotipo].genes[iGen])

			// Probabilidad de mutacion?
			// Info("MUT?", rand.Float64() < probMutacion)
			// if rand.Float64() >= probMutacion {
			// 	continue
			// }

			// Valor de mutacion
			// valMut := math.Pow(rand.Float64(), 2)
			valMut := rand.Float64()

			// Cambiar signo?
			if rand.Intn(2) == 0 {
				valMut *= -1
			}

			// Mutar gen
			nuevoGenotipo.genes[iGen] += valMut
		}

		// Añadir nuevo  gen
		nuevosGenes = append(nuevosGenes, nuevoGenotipo)
	}

	// Retornar nuevo set de genes mutados
	return nuevosGenes
}

func heuristica_h(entrada Salida, salida Salida) float64 {
	// Respuesta perfecta (0 de error)
	var res float64 = 0.0

	// Calcular error en diasRecuperado
	if salida.DiasRecuperado != -1 {
		diff64 := float64(entrada.DiasRecuperado - salida.DiasRecuperado)
		res += math.Pow(diff64, 2)
	}
	// Calcular error en diasFallecido
	if salida.DiasFallecido != -1 {
		diff64 := float64(entrada.DiasFallecido - salida.DiasFallecido)
		res += math.Pow(diff64, 2)
	}
	// Calcular error en diasReferido
	if salida.DiasReferido != -1 {
		diff64 := float64(entrada.DiasReferido - salida.DiasReferido)
		res += math.Pow(diff64, 2)
	}

	// Retorna res
	return res
}

// Calcula la suma total de las eurisiticas de las entradas y salidas
func heuristica(entradas []Salida, salidas []Salida) float64 {
	// Respuesta perfecta (0 de error)
	var res float64 = 0.0

	// Sumar error para cada una entrada<->salida
	for i := 0; i < len(entradas); i++ {
		res += heuristica_h(entradas[i], salidas[i])
	}

	// Retorna res
	return res
}

func obtenerEntradas(genotipo Genotipo, atenciones []Atencion) []Salida {
	// Entradas
	var entradas = make([]Salida, len(atenciones))
	for i := 0; i < len(atenciones); i++ {
		// Obtener v
		// v := float64(FechaStr2Time(atenciones[i].Fecha_ingreso).Unix())
		v := ObtenerV(atenciones[i])

		// v, _ := strconv.ParseFloat(atenciones[i].Id_eess, 64)
		entradas[i] = Ejecutar(v, genotipo)

	}

	return entradas
}

// Ordenar genes por euristica
func ordPorHeuristica(genotipos []Genotipo, atenciones []Atencion) ([]Genotipo, []float64) {
	// Crear arreglo de heuristicas por genotipo
	heuristicas := make([]float64, len(genotipos))

	// Procesar csv
	salidas := make([]Salida, len(atenciones))
	for i := 0; i < len(atenciones); i++ {
		salidas[i] = EjecutarCsv(atenciones[i])
	}

	// Procesar cada genotipo
	for i := 0; i < len(genotipos); i++ {
		entradas := obtenerEntradas(genotipos[i], atenciones)
		heuristicas[i] = heuristica(entradas, salidas)
		// Info("ENTRADAS", entradas)
		// Info("SALIDAS", salidas)
	}

	// Ordenar genotipos por heuristica
	sort.SliceStable(genotipos, func(i, j int) bool {
		return heuristicas[i] > heuristicas[j]
	})

	// Ordenar heuristicas
	sort.SliceStable(heuristicas, func(i, j int) bool {
		return heuristicas[i] > heuristicas[j]
	})

	// Retornar genotipos ordenados
	return genotipos, heuristicas
}
