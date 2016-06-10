package main

import (
	"github.com/samalba/dockerclient"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)


func main() {
	router := mux.NewRouter()

	router.Methods("GET").Path("/container/{id}/portbinding").Queries("port", "{port:[0-9]+}").Queries("protocol", "{protocol:[a-z]+}").HandlerFunc(PortBinding)

	router.Methods("GET").Path("/").HandlerFunc(HealthCheck)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}


func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Print("Health check")
	w.WriteHeader(200)
}

func PortBinding(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	port := vars["port"]
	protocol := vars["protocol"]

	log.Print("Container: " + id + " port: " + port + " protocol: " + protocol)

	pb := ExposedPorts(id, port, protocol)
	//if pb != nil {
		json.NewEncoder(w).Encode(pb)
	//} else {
	//	w.WriteHeader(404)
	//}

}

func ExposedPorts(id string, port string, protocol string) []dockerclient.PortBinding {
	docker, _ := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	info, _ := docker.InspectContainer(id)
	if info == nil {
		return nil
	}

	ports :=  info.NetworkSettings.Ports

	return ports[port + "/" + protocol]
}
