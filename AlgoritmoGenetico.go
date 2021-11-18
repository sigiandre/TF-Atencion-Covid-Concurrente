package main

type Atencion struct {
	id_persona            int
	EESS                  int
	fecha_ingreso         string
	hora_ingreso          string
	recuperado            bool
	fecha_alta            string
	recuperado_voluntario bool
	fecha_alta_voluntaria string
	fallecido             bool
	fecha_fallecido       string
	referido              bool
	fecha_referido        string
	EESS_destino_id       int
}
