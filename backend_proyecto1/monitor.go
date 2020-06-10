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
	"encoding/json"
	"bytes"
)

func getCPU(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("sh", "-c", "ps -eo pcpu | sort -k 1 -r | head -50")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("CPU obtenido correctamente")

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

type proc struct {
    PID string
    Username string
	Command string
	State string
	Ram string
}
type Salida struct{
	Output []proc
}

func getRunningProcess(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sh", "-c", "ps aux")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
    
    output := string(out[:])
	//fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
	
	// for {key}, {value} := range {list}	

	contador := 0
	for i := 1; i < len(s)-1; i++ {
		 procs := strings.Fields(s[i])
		 if strings.Contains(procs[7],"R") {
			 contador = contador + 1
		 }
	 }	
		fmt.Fprintf(w, strconv.Itoa(contador))
}

func getSuspendProcess(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sh", "-c", "ps aux")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
    
    output := string(out[:])
	//fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
	
	// for {key}, {value} := range {list}	

	contador := 0
	for i := 1; i < len(s)-1; i++ {
		 procs := strings.Fields(s[i])
		 if strings.Contains(procs[7],"S") {
			 contador = contador + 1
		 }
	 }	
		fmt.Fprintf(w, strconv.Itoa(contador))
}

func getStopProcess(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sh", "-c", "ps aux")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
    
    output := string(out[:])
	//fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
	
	// for {key}, {value} := range {list}	

	contador := 0
	for i := 1; i < len(s)-1; i++ {
		 procs := strings.Fields(s[i])

		if strings.Contains(procs[7],"T") {
			contador= contador+1
		}else if strings.Contains(procs[7],"t") {
			contador= contador+1
		}
	 }	
		fmt.Fprintf(w, strconv.Itoa(contador))
}

func getZombieProcess(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sh", "-c", "ps aux")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
    
    output := string(out[:])
	//fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
	
	// for {key}, {value} := range {list}	

	contador := 0
	for i := 1; i < len(s)-1; i++ {
		 procs := strings.Fields(s[i])
		 if strings.Contains(procs[7],"Z") {
			 contador = contador+1
		 }
	 }	
		fmt.Fprintf(w, strconv.Itoa(contador))
}

func getTotalProcess(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sh", "-c", "ps aux")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
    
    output := string(out[:])
	//fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
	
	// for {key}, {value} := range {list}	
		fmt.Fprintf(w, strconv.Itoa(len(s)))

}


func getAllProcess(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sh", "-c", "ps aux")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
    
    output := string(out[:])
	//fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
	
	// for {key}, {value} := range {list}	
	    arreglo := []proc{}
	   for i := 1; i < len(s)-1; i++ {
			procs := strings.Fields(s[i])
			pr := proc{procs[1],procs[0],procs[10],procs[7],procs[3]}
			arreglo = append(arreglo,pr)
		}	
		salida := Salida{arreglo}
		js, err := json.Marshal(salida)
		if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		  
}

func killProcess(w http.ResponseWriter, r *http.Request){
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	str := "kill "+newStr
	cmd := exec.Command("sh", "-c", str)
	out, err := cmd.CombinedOutput()
	if err != nil {
//		log.Fatal(err)
fmt.Println("error")
	}
    fmt.Println(out)
	fmt.Fprintf(w, "OK")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Server Running on port: 8080")
	router.HandleFunc("/cpu", getCPU).Methods("GET")
	router.HandleFunc("/ram", getRam).Methods("GET")
	router.HandleFunc("/total", getTotalRam).Methods("GET")
	router.HandleFunc("/actualram", getConsumeRam).Methods("GET")
	router.HandleFunc("/runningprocess", getRunningProcess).Methods("GET")
	router.HandleFunc("/suspendprocess", getSuspendProcess).Methods("GET")
	router.HandleFunc("/stopprocess", getStopProcess).Methods("GET")
	router.HandleFunc("/zombieprocess", getZombieProcess).Methods("GET")
	router.HandleFunc("/totalprocess", getTotalProcess).Methods("GET")
	router.HandleFunc("/allprocess", getAllProcess).Methods("GET")
	router.HandleFunc("/killprocess", killProcess).Methods("POST")
	// cors.Default() setup the middleware with default options being
    // all origins accepted with simple methods (GET, POST). See
    // documentation below for more options.
    handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
