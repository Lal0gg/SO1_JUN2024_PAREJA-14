package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"server/Controller"
	"server/Database" // Asegúrate de importar el paquete Database

	"server/Model"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RAM struct {
	RamUsed float64 `json:"RamUsed"`
	RamFree float64 `json:"RamFree"`
}

type Process struct {
	PID      int       `json:"pid"`
	Name     string    `json:"name"`
	User     int       `json:"user"`
	State    string    `json:"state"`
	RAM      int       `json:"ram"`
	PIDPadre int       `json:"pidPadre"`
	Child    []Process `json:"child"`
}

type CPU struct {
	CpuUsed    int       `json:"CpuUsed"`
	CpuPercent float64   `json:"CpuPercent"`
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
	router.HandleFunc("/create-process", CreateProcess).Methods("POST")
	router.HandleFunc("/kill-process", KillProcess).Methods("POST")
	fmt.Println("Servidor escuchando en el puerto 8080...")

	// Wrap the router with CORS support
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(router)

	if err := Database.Connect(); err != nil { // Llamar a Database.Connect() en lugar de Instance.Connect()
		log.Fatal(err)
	}

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

	// Insertar datos en la base de datos
	Controller.InsertData("ram", strconv.Itoa(int(ramInfo.RamUsed)), primitive.NewDateTimeFromTime(time.Now()))
}

func readCpuInfo() CPU {
	var cpuInfo CPU

	// Ejecutar el comando cat para obtener la salida del archivo /proc/cpu_so1_1s2024
	cmd := exec.Command("cat", "/proc/cpu_so1_1s2024")
	outCpu, err := cmd.Output()
	if err != nil {
		fmt.Println("Error al ejecutar el comando:", err)
		return cpuInfo
	}

	// Decodificar la salida del comando en la estructura CPU
	err = json.Unmarshal(outCpu, &cpuInfo)
	if err != nil {
		fmt.Println("Error al parsear JSON:", err)
		return cpuInfo
	}

	return cpuInfo
}

func ReadPercentCpu() (float64, error) {
	cmd, err := exec.Command("sh", "-c", "mpstat 1 1").Output()
	if err != nil {
		return 0, fmt.Errorf("error al ejecutar el comando: %v", err)
	}

	output := string(cmd)

	lines := strings.Split(output, "\n")
	var idleLine string
	for _, line := range lines {
		if strings.Contains(line, "all") {
			idleLine = line
			break
		}
	}

	if idleLine == "" {
		return 0, fmt.Errorf("no se encontró la línea con 'all'")
	}

	fields := strings.Fields(idleLine)

	idleStr := fields[len(fields)-1]

	idle, err := strconv.ParseFloat(idleStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error al convertir %%idle: %v", err) // Corrección aquí para escapar el %
	}

	cpuUsed := 100.0 - idle
	return cpuUsed, nil
}

func getCpuInfo(w http.ResponseWriter, r *http.Request) {
	cpuInfo := readCpuInfo()
	cpuPercent, err := ReadPercentCpu()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cpuInfo.CpuPercent = cpuPercent

	response, err := json.Marshal(cpuInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Escribir la respuesta JSON
	w.Write(response)

	timestamp := primitive.NewDateTimeFromTime(time.Now())
	// Insertar procesos principales
	for _, process := range cpuInfo.Processes {
		procData := Model.ProcessData{
			ID:        primitive.NewObjectID(),
			PID:       process.PID,
			Name:      process.Name,
			State:     process.State,
			PIDPadre:  0, // Proceso principal no tiene padre
			Timestamp: timestamp,
		}
		Controller.InsertProcessData(procData)

		// Insertar procesos hijos
		insertChildProcesses(process.Child, process.PID, timestamp)
	}
}

func insertChildProcesses(children []Process, parentPID int, timestamp primitive.DateTime) {
	for _, child := range children {
		procData := Model.ProcessData{
			ID:        primitive.NewObjectID(),
			PID:       child.PID,
			Name:      child.Name,
			State:     child.State,
			PIDPadre:  parentPID,
			Timestamp: timestamp,
		}
		Controller.InsertProcessData(procData)

		// Recursivamente insertar hijos de los hijos
		insertChildProcesses(child.Child, child.PID, timestamp)
	}
}

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	// Crear un comando que ejecute un sleep infinito
	cmd := exec.Command("sleep", "infinity")

	// Asignar la salida estándar y de error del comando al mismo que el del proceso principal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Iniciar el comando
	if err := cmd.Start(); err != nil {
		http.Error(w, fmt.Sprintf("Error al iniciar el comando: %v", err), http.StatusInternalServerError)
		return
	}

	// Obtener el PID del proceso
	pid := cmd.Process.Pid
	fmt.Printf("El PID del proceso es: %d\n", pid)

	// Responder con el PID del proceso
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("El PID del proceso es: %d\n", pid)))
}

func KillProcess(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud
	var requestData map[string]int
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Obtener el PID del proceso de la solicitud
	pid, exists := requestData["pid"]
	if !exists {
		http.Error(w, "PID no proporcionado", http.StatusBadRequest)
		return
	}

	// Matar el proceso
	if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
		http.Error(w, fmt.Sprintf("Error al matar el proceso: %v", err), http.StatusInternalServerError)
		return
	}

	// Responder con el mensaje de éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Proceso eliminado"}`))
}
