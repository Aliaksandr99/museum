package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bukhavtsov/museum/back-end/pkg/data"

	"github.com/gorilla/mux"
)

type ArtifactData interface {
	ReadAll() ([]*data.ArtifactMaster, error)
	Add(artifact *data.ArtifactMaster) (int64, error)
}

type artifactAPI struct {
	data ArtifactData
}

func ServerArtifactResource(r *mux.Router, data ArtifactData) {
	api := &artifactAPI{data: data}
	r.HandleFunc("/artifacts", api.getArtifacts).Methods("GET")
	r.HandleFunc("/artifacts", api.createArtifact).Methods("POST")
}

func (api artifactAPI) getArtifacts(writer http.ResponseWriter, request *http.Request) {
	cards, err := api.data.ReadAll()
	if err != nil {
		log.Println("artifacts haven't been read")
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	err = json.NewEncoder(writer).Encode(cards)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (api artifactAPI) createArtifact(writer http.ResponseWriter, request *http.Request) {
	artifact := new(data.ArtifactMaster)
	err := json.NewDecoder(request.Body).Decode(&artifact)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if artifact == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	printArtifact(artifact) // for temp manual testing
	artifactId, err := api.data.Add(artifact)
	if err != nil {
		log.Println("artifact hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Location", fmt.Sprintf("/artifact/%d", artifactId))
	writer.WriteHeader(http.StatusCreated)
}

func printArtifact(artifact *data.ArtifactMaster) {
	fmt.Println("Artifact")
	fmt.Println("artifact_id:", artifact.ID)
	fmt.Println("artifact_object_group:", artifact.ObjectGroup)
	fmt.Println("artifact_elements:", artifact.Elements)
	fmt.Println("artifact_preservation:", artifact.Preservation)
	fmt.Println("artifact_measurement:", artifact.ArtifactMeasurement)
	fmt.Println("artifact_excavationDate:", artifact.ExcavationDate)
	fmt.Println("artifact_transferredBy:", artifact.TransferredBy)
	fmt.Println("artifact_artifactStyle:", artifact.ArtifactStyle)
	fmt.Println("artifact_creator:", artifact.Creator)
}
