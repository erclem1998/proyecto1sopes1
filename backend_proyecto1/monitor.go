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
	"io/ioutil"
	"regexp"
	"time"
)

func getCPU(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("sh", "-c", "ps -eo pcpu | sort -k 1 -r | head -50")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

    go fmt.Println("CPU obtenido correctamente")

    output := string(out[:])
	//fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
    cpuUsado := 0.0
    for i := 1; i < 51; i++ {

    	valor, err := strconv.ParseFloat(strings.Trim(s[i], " "), 64)
    	if err != nil {
    		go fmt.Println("valorError ->" + s[i] + "<-" + strconv.Itoa(i))
    		go fmt.Println(err)
		}

		cpuUsado += valor 
		//go fmt.Println("valor ->" + s[i] + "<-" + strconv.Itoa(i))

	}


    fmt.Fprintf(w, "%f", cpuUsado)
}


func getRam(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("sh", "-c", "cat /proc/meminfo | grep -e MemTotal -e MemFree -e MemAvailable")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

    //go fmt.Println("Ram obtenida correctamente")
    output := string(out[:])
    //fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
    //MemTotal
	sMemTotal := strings.Fields(s[0])
	MemTotal := sMemTotal[1]
	//go fmt.Println(s[0])

	//MemAvailable
	//go fmt.Println(s[2])
	//MemFree
	sMemFree := strings.Fields(s[1])
	MemFree := sMemFree[1]
	//go fmt.Println(s[1])

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

    go fmt.Println("Ram obtenida correctamente")
    output := string(out[:])
    //fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
    //MemTotal
	sMemTotal := strings.Fields(s[0])
	MemTotal := sMemTotal[1]
	//go fmt.Println(s[0])

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
    
    //go fmt.Println("Ram obtenida correctamente")
    output := string(out[:])
    //fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
    //MemTotal
	sMemTotal := strings.Fields(s[0])
	MemTotal := sMemTotal[1]
	//MemAvailable
	//go fmt.Println(s[2])
	//MemFree
	sMemFree := strings.Fields(s[1])
	MemFree := sMemFree[1]
	//go fmt.Println(s[1])

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
    PID int
    Username string
	Command string
	State string
	Ram string
}
type Salida struct{
	Output []proc
}

type statistics struct{
	Total int
	Running int
	Suspend int
	Stop int
	Zombie int
}

func Statistics(w http.ResponseWriter, r *http.Request){
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		go fmt.Println(err)
	}
	contador := statistics{0,0,0,0,0}
	for _, file := range files {
		match, _ := regexp.MatchString("[0-9]+", file.Name())
		if match{
			//go fmt.Println(file.Name())
			pid := file.Name()

			cmd := exec.Command("sh", "-c", "cat /proc/"+pid+"/stat")
			out, err2 := cmd.CombinedOutput()
			if err2 != nil {
				go fmt.Println(err2)
			}
    		output := string(out[:])
			//fmt.Fprintf(w, output)
			
			s := strings.Split(output, "\n")
			
			for i := 0; i < len(s)-1; i++ {
				procs := strings.Fields(s[i])
				contador.Total++
				if strings.Contains(procs[2],"R") {
					contador.Running++
				}else if strings.Contains(procs[2],"S") {
					contador.Suspend++
				}else if strings.Contains(procs[2],"T") {
					contador.Stop++
				}else if strings.Contains(procs[2],"Z") {
					contador.Zombie++
				}
			}	
		}
	}
	go fmt.Println("Estadisticas: ",contador)
		js, err3 := json.Marshal(contador)
		if err3 != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
}




func getAllProcess(w http.ResponseWriter, r *http.Request){
	//memTotal
	cmd2 := exec.Command("sh", "-c", "cat /proc/meminfo | grep -e MemTotal")
	out2, err5 := cmd2.CombinedOutput()
	if err5 != nil {
		log.Fatal(err5)
	}

    go fmt.Println("Ram obtenida correctamente")
    output2 := string(out2[:])
    //fmt.Fprintf(w, output)

    s2 := strings.Split(output2, "\n")
    //MemTotal
	sMemTotal2 := strings.Fields(s2[0])
	MemTotal2 := sMemTotal2[1]
	//go fmt.Println(s[0])
	i2, err8 := strconv.Atoi(MemTotal2)
	if err8 != nil {
		log.Fatal(err8)
	}
	//PorcertajeRam
	//procesos
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		go fmt.Println(err)
	}
	arreglo := []proc{}
	for _, file := range files {
		match, _ := regexp.MatchString("[0-9]+", file.Name())
		if match{
			//go fmt.Println(file.Name())
			pid := file.Name()
				//go fmt.Println(procs)
				piid, err3 := strconv.Atoi(pid)
				if err3 != nil {
					go fmt.Println(err3)
				}
				//memoria y uid
				cmd3 := exec.Command("sh", "-c", "cat /proc/"+pid+"/status")
				out3, err9 := cmd3.CombinedOutput()
				if err9 != nil {
					go fmt.Println(err9)
				}
				output3 := string(out3[:])

		//fmt.Fprintf(w, output)
				s3 := strings.Split(output3, "\n")
				uids := strings.Fields(s3[8])
				uid := uids[1]
				MemUsed := strings.Fields(s3[17])

				States := strings.Fields(s3[2])
				state := States[1]
				Names := strings.Fields(s3[0])
				name := Names[1]
				if MemUsed[0] == "VmSize:"{
					memuses,err7:=  strconv.Atoi(MemUsed[1])					
				if err7 != nil {
					go fmt.Println(err7)
				}
				var ram1 float64 
				ram1 = float64(memuses) / float64(i2)
				porc_ram:= ram1
				pr := proc{piid, uid, name, state, strconv.FormatFloat(porc_ram, 'f', 2, 64)+"%"}
					arreglo = append(arreglo,pr)
				} else{
					pr := proc{piid, uid, name, state, "0%"}
					arreglo = append(arreglo,pr)
				}
				
				

				
		}
	}
	arreglo = BubbleSort(arreglo)
	salida := Salida{arreglo}
	js, err3 := json.Marshal(salida)
	if err3 != nil {
	  http.Error(w, err3.Error(), http.StatusInternalServerError)
	  return
	}
	go fmt.Println("Tabla de procesos generada")
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
go fmt.Println("error")
	}
    go fmt.Println(out)
	fmt.Fprintf(w, "OK")
}

type Proceso struct{
	Id string
	ParentId string
	Label string
	Items []Proceso
}
type Arbol struct{
	Arbol []Proceso
}

func BubbleSort(numbers []proc) []proc {
	//Start the loop in reverse order, so the loop will start with length
	//which is equal to the length of input array and then loop untill
	//reaches 1
	for i := len(numbers); i > 0; i-- {
	   //The inner loop will first iterate through the full length
	   //the next iteration will be through n-1
	   // the next will be through n-2 and so on
	   for j := 1; j < i; j++ {
		  if numbers[j-1].PID > numbers[j].PID {
			 intermediate := numbers[j]
			 numbers[j] = numbers[j-1]
			 numbers[j-1] = intermediate
						 }
			  }
	   }
	return numbers
  }


func getTreeProcess(w http.ResponseWriter, r *http.Request){
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		go fmt.Println(err)
	}
	arreglo := []Proceso{}
	for _, file := range files {
		match, _ := regexp.MatchString("[0-9]+", file.Name())
		if match{
			//go fmt.Println(file.Name())
			pid := file.Name()

			cmd := exec.Command("sh", "-c", "cat /proc/"+pid+"/stat")
			out, err2 := cmd.CombinedOutput()
			if err2 != nil {
			go fmt.Println(err2)
			}
    		output := string(out[:])
			//fmt.Fprintf(w, output)
			
			s := strings.Split(output, "\n")
			
			for i := 0; i < len(s)-1; i++ {
				procs := strings.Fields(s[i])
				emp := []Proceso{}
				pr := Proceso{pid,procs[3],procs[1],emp}
				arreglo = append(arreglo,pr);
			}	
		}
	}
	
		salida := Arbol{arreglo}
		js, err := json.Marshal(salida)
		if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)	  
}

func existe_pid(arreglo *[]Proceso, proceso *Proceso) bool{
	for j := 0; j < len(*arreglo); j++{
		if 	(*proceso).ParentId == (*arreglo)[j].Id {
			(*arreglo)[j].Items = append((*arreglo)[j].Items,*proceso)	

			return true		
		}else{
			a:= existe_pid(&(*arreglo)[j].Items, &(*proceso))
			if a {
				return true
			}
		}
	}

	return false
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	
	go fmt.Println("Server Running on port: 8080")
	go router.HandleFunc("/cpu", getCPU).Methods("GET")
	go router.HandleFunc("/ram", getRam).Methods("GET")
	go router.HandleFunc("/total", getTotalRam).Methods("GET")
	go router.HandleFunc("/actualram", getConsumeRam).Methods("GET")
	go router.HandleFunc("/statistics", Statistics).Methods("GET")
	go router.HandleFunc("/allprocess", getAllProcess).Methods("GET")
	go router.HandleFunc("/killprocess", killProcess).Methods("POST")
	go router.HandleFunc("/treeprocess", getTreeProcess).Methods("GET")
	// cors.Default() setup the middleware with default options being
    // all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	time.Sleep(time.Second)
    handler := cors.Default().Handler(router)
	 http.ListenAndServe(":8080", handler)
	 log.Fatal(http.ListenAndServe(":8080", router))
	 
}
