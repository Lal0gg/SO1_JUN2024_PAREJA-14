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

type Process struct {
	PID      int            `json:"pid"`
	Name     string         `json:"name"`
	User     int            `json:"user"`
	State    int            `json:"state"`
	RAM      int            `json:"ram"`
	Children []ChildProcess `json:"children"`
}

type ChildProcess struct {
	PID      int    `json:"pid"`
	Name     string `json:"name"`
	State    int    `json:"state"`
	PIDPadre int    `json:"pidPadre"`
}

type CPU struct {
	CpuUsed    int       `json:"CpuUsed"`
	CpuPercent int       `json:"CpuPercent"`
	Processes  []Process `json:"processes"`
	Running    int       `json:"running"`
	Sleeping   int       `json:"sleeping"`
	Zombie     int       `json:"zombie"`
	Stopped    int       `json:"stopped"`
	Total      int       `json:"total"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ram", getRamInfo).Methods("GET")
	router.HandleFunc("/cpu", getCpuInfo).Methods("GET")
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
		if strings.Contains(ramSplit2[0], "RamUsed") {
			ramInfo.RamUsed = round(val, 2)
			ramInfo.RamFree = round(100-val, 2)
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

func readCpuInfo() CPU {
	var cpuInfo CPU
	cmd := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
	outCpu, err := cmd.Output()
	if err != nil {
		fmt.Println("Error al ejecutar el comando:", err)
		return cpuInfo
	}
	output := string(outCpu)
	fmt.Println("Output Success")
	err = json.Unmarshal([]byte(output), &cpuInfo)
	if err != nil {
		fmt.Println("Error al parsear JSON:", err)
		return cpuInfo
	}

	return cpuInfo

}

func getCpuInfo(w http.ResponseWriter, r *http.Request) {
	cpuInfo := readCpuInfo()

	response, err := json.Marshal(cpuInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Escribir la respuesta JSON
	w.Write(response)
}
