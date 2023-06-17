package handler

import (
	"bytes"
	"copy-s3-object/dto"
	"copy-s3-object/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type S3Handler interface {
	GetHealth(http.ResponseWriter, *http.Request)
	CopyObjects(http.ResponseWriter, *http.Request)
}

type s3Handler struct {
	service service.S3Service
}

func NewS3Handler(s3Service service.S3Service) S3Handler {
	return &s3Handler{service: s3Service}
}

func (hndlr *s3Handler) GetHealth(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusAccepted)

	type healthStruct struct{ Status string }
	var responseBytes bytes.Buffer
	json.NewEncoder(&responseBytes).Encode(healthStruct{Status: "running"})
	res.Write(responseBytes.Bytes())
}

type responseStruct struct {
	Msg string `json:"msg"`
}

func (hndlr *s3Handler) CopyObjects(res http.ResponseWriter, req *http.Request) {

	// Parse request
	var requestStruct dto.CopyRequest
	err := json.NewDecoder(req.Body).Decode(&requestStruct)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while decoding request: %s", err)
		var responseBytes bytes.Buffer
		json.NewEncoder((&responseBytes)).Encode(responseStruct{Msg: msg})
		res.WriteHeader(http.StatusBadRequest)
		res.Write(responseBytes.Bytes())
		return
	}

	err = requestStruct.Validate()
	if err != nil {
		msg := fmt.Sprintf("Error occurred while validating request payload %+v: %s", requestStruct, err)
		var responseBytes bytes.Buffer
		json.NewEncoder(&responseBytes).Encode(responseStruct{Msg: msg})
		res.WriteHeader(http.StatusBadRequest)
		res.Write(responseBytes.Bytes())
		return
	}

	// get response from s3 client
	log.Printf("req payload: %+v\n", requestStruct)
	err = hndlr.service.CopyAllObjects(requestStruct)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while copying: %s", err)
		var responseBytes bytes.Buffer
		json.NewEncoder(&responseBytes).Encode(responseStruct{Msg: msg})
		res.WriteHeader(http.StatusBadRequest)
		res.Write(responseBytes.Bytes())
		return
	}

	msg := "Copied successfully"
	var responseBytes bytes.Buffer
	json.NewEncoder(&responseBytes).Encode(responseStruct{Msg: msg})
	res.WriteHeader(http.StatusOK)
	res.Write(responseBytes.Bytes())
}
