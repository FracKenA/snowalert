package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// GenerateMD5 will generate an MD5 hash for the data provided.
func GenerateMD5(MD5Data string, incident *Incident) *Incident {
	data := []byte(MD5Data + MD5Seed)
	result := md5.Sum(data)
	incident.CorrelationID = hex.EncodeToString(result[:])

	return incident
}

// DelayProcessing will cause a random delay for a period of seconds.
func DelayProcessing(incident *Incident) *Incident {
	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn( 3+ 1)
	if LogLevel == "DEBUG" {
		incident.LogOutput = append(incident.LogOutput, fmt.Sprintf(fmt.Sprintf("Sleping: %d", delay)))
	}
	time.Sleep(time.Duration(delay) * time.Second)
	return incident
}

// TransactionToken creates a token to allow searching for transations.
func TransactionToken(incident *Incident, n int) *Incident {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		incident.Err = append(incident.Err, fmt.Sprintf("TransactionToken - unable to read %s", err))
	}
	incident.TransactionID = hex.EncodeToString(bytes)

	return incident
}

// CloseRequest in go by default using defer function does not allow error handling.
// this function is designed to mitigate that issue.
func CloseRequest(toClose *http.Request) {
	err := toClose.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
