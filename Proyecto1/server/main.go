package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type RAM struct {
	RamUsed float64 `json:"RamUsed"`
	RamFree float64 `json:"RamFree"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ram", getRamInfo).Methods("GET")
	fmt.Println("Servidor escuchando en el puerto 8080...")

	// Wrap the router with CORS support
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(router)

	http.ListenAndServe(":8080", corsHandler)
}

func readRamInfo() RAM {
	var ramInfo RAM

	cmd := exec.Command("sh", "-c", "cat /proc/ram_so1_jun2024")
	outRam, err := cmd.Output()
	if err != nil {
		fmt.Println("Error al ejecutar el comando:", err)
		return ramInfo
	}

	ramOutput := string(outRam)
	ramSplit := strings.Split(ramOutput, "\n")
	for _, line := range ramSplit {
		ramSplit2 := strings.Split(line, ",")
		if len(ramSplit2) < 2 {
			continue // Si la línea no tiene el formato esperado, ignorarla
		}

		// Extraer el valor numérico después de la coma
		valStr := strings.TrimSpace(ramSplit2[1])
		fmt.Println("Valor:", valStr)
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			fmt.Println("Error al convertir el valor:", err)
			continue
		}

		// Asignar el valor a la estructura RAM según el nombre de la variable
		if strings.Contains(ramSplit2[0], "RAMused") {
			ramInfo.RamUsed = round(val/100000, 2)
		} else if strings.Contains(ramSplit2[0], "RAMfree") {
			ramInfo.RamFree = round(100-ramInfo.RamUsed, 2)
		}
	}

	return ramInfo
}

func round(num float64, decimals int) float64 {
	pow := math.Pow(10, float64(decimals))
	return math.Round(num*pow) / pow
}

func getRamInfo(w http.ResponseWriter, r *http.Request) {
	ramInfo := readRamInfo()

	response, err := json.Marshal(ramInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Escribir la respuesta JSON
	w.Write(response)
}
