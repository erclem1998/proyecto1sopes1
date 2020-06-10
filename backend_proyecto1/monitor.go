package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os/exec"
	"strings"
	"strconv"
	"github.com/rs/cors"
)

func getRam(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("sh", "-c", "cat /proc/meminfo | grep -e MemTotal -e MemFree -e MemAvailable")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

    //fmt.Println("Ram obtenida correctamente")
    output := string(out[:])
    //fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
    //MemTotal
	sMemTotal := strings.Fields(s[0])
	MemTotal := sMemTotal[1]
	//fmt.Println(s[0])

	//MemAvailable
	//fmt.Println(s[2])
	//MemFree
	sMemFree := strings.Fields(s[1])
	MemFree := sMemFree[1]
	//fmt.Println(s[1])

	//PorcertajeRam
	i1, err := strconv.Atoi(MemTotal)
	if err != nil {
		log.Fatal(err)
	}

	i2, err := strconv.Atoi(MemFree)
	if err != nil {
		log.Fatal(err)
	}
	resultado := (100-((i2*100)/i1))
	fmt.Fprintf(w, strconv.Itoa(resultado))

}


func getTotalRam(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("sh", "-c", "cat /proc/meminfo | grep -e MemTotal")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("Ram obtenida correctamente")
    output := string(out[:])
    //fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
    //MemTotal
	sMemTotal := strings.Fields(s[0])
	MemTotal := sMemTotal[1]
	//fmt.Println(s[0])

	//PorcertajeRam
	i1, err := strconv.Atoi(MemTotal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, strconv.Itoa(i1/1024))

}
func getConsumeRam(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("sh", "-c", "cat /proc/meminfo | grep -e MemTotal -e MemFree -e MemAvailable")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
    
    //fmt.Println("Ram obtenida correctamente")
    output := string(out[:])
    //fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
    //MemTotal
	sMemTotal := strings.Fields(s[0])
	MemTotal := sMemTotal[1]
	//MemAvailable
	//fmt.Println(s[2])
	//MemFree
	sMemFree := strings.Fields(s[1])
	MemFree := sMemFree[1]
	//fmt.Println(s[1])

	i1, err := strconv.Atoi(MemTotal)
	if err != nil {
		log.Fatal(err)
	}
	i2, err := strconv.Atoi(MemFree)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Fprintf(w, strconv.Itoa((i1-i2)/1024))
}


func getCPU(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("sh", "-c", "ps aux | grep -e STAT -e R")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("Procesos obtenidos correctamente")
    
    output := string(out[:])
	//fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
    cpuUsado := 0.0
    for i := 1; i < 51; i++ {

    	valor, err := strconv.ParseFloat(strings.Trim(s[i], " "), 64)
    	if err != nil {
    		fmt.Println("valorError ->" + s[i] + "<-" + strconv.Itoa(i))
    		fmt.Println(err)
		}
		
		cpuUsado += valor 
		//fmt.Println("valor ->" + s[i] + "<-" + strconv.Itoa(i))
		
	}

    
    fmt.Fprintf(w, "%f", cpuUsado)
}

func getRunningProcess(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sh", "-c", "ps aux | grep -e USER -e PID -e %MEM -e COMMAND -e STAT")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("CPU obtenido correctamente")
    
    output := string(out[:])
	//fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
	
	// for {key}, {value} := range {list}
	for _, process := range s {
		fmt.Println(process,"\n")
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Server Running on port: 8080")
	router.HandleFunc("/cpu", getCPU).Methods("GET")
	router.HandleFunc("/ram", getRam).Methods("GET")
	router.HandleFunc("/total", getTotalRam).Methods("GET")
	router.HandleFunc("/actualram", getConsumeRam).Methods("GET")
	// cors.Default() setup the middleware with default options being
    // all origins accepted with simple methods (GET, POST). See
    // documentation below for more options.
    handler := cors.Default().Handler(router)
    http.ListenAndServe(":8080", handler)
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
