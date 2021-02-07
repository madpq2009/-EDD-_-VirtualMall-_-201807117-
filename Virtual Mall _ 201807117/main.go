package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
)
//Estructuras JSON----------------------------------------------------------------------------------
type TiendasS struct{
	Nombre string `json:Nombre`
	Descripcion string `json:Descripcion`
	Contacto string `json:Contacto`
	Calificacion int `json:Calificacion`
}
type DepartamentosS struct{
	Nombre string `json:Nombre`
	Tiendas []TiendasS `json:Tiendas`
}
type DatosS struct{
	Indice string `json:Indice`
	Departamentos []DepartamentosS `json:Departamentos`
}
type ArchivoGlobal struct {
	Datos []DatosS `json:Datos`
}

//main----------------------------------------------------------------------------------
func main() {
	r := mux.NewRouter()
    r.HandleFunc("/Cargartienda", Cargartienda).Methods("POST")
    //r.HandleFunc("/getArreglo", getArreglo).Methods("GET")
    //r.HandleFunc("/TiendaEspecifica", TiendaEspecifica).Methods("POST")
	//r.HandleFunc(":/id/:numero", IdNumero).Methods("GET")
    log.Fatal(http.ListenAndServe(":3000",r))
	http.Handle("/", r)
}
//Funciones
func Cargartienda(w http.ResponseWriter, r *http.Request){
	var ag ArchivoGlobal
	reqBody,err:= ioutil.ReadAll(r.Body)
	if err !=nil{
		fmt.Fprintf(w,"Error al insertar")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &ag)
	json.NewEncoder(w).Encode(ag)
}