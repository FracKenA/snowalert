package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
)

// PostV1Incident is the first version of the handler for  a POST method request to the /api/v1/incident endpoint
// This function proceeds to call all related function, lookups and actions and at the end of the function the POST
// action is completed and logged.
func PostV1Incident(context echo.Context) error {
	// Create incident interface
	var incident Incident // Declared in models.go

	// Create custom error handling with deferred closure to unsure there are no unhandled errors.
	defer CloseRequest(context.Request()) // shortlines.go

	// Decode the inbound
	err := json.NewDecoder(context.Request().Body).Decode(&incident)
	if err != nil {
		incident.Err = append(incident.Err, fmt.Sprintf("Failed reading the request body %s", err))
	}

	// Generate custom transaction token for log correlation between alert client, event router and ServiceNow incident
	TransactionToken(&incident, 8) // shortlines.go

	// Query ServiceNow with CI Name to get the SysID
	GetCISysID(&incident)

	// Generate MD5 for correlation ID
	GenerateMD5(incident.CIName + incident.Source, &incident)
	if LogLevel == "DEBUG" {
		incident.LogOutput = append(incident.LogOutput, fmt.Sprintf("MatchLookup - correlationID generate - %s", incident.CorrelationID))
	}

	// Provide short delay, preventing fast firing alerts.
	DelayProcessing(&incident) // shortlines.go

	// Lookup if there is an existing active incident
	LookupIncident(&incident)

	// Decide what to do if there is an existing incident.
	if incident.IncidentExists != true {
		if incident.CISysIDFound == true {
			CreateIncident(&incident)
			AddEvent(&incident)
		} else {
			CreateIncidentNoMatch(&incident)
			AddEvent(&incident)
		}
	} else if incident.IncidentExists == true{
		// UpdateIncident(&incident)
		AddEvent(&incident)
	}

	var returnStatus int

	if incident.IncidentID == "" {
		returnStatus = 406
	} else {
		returnStatus = 201
	}
	return context.JSON(returnStatus, incident)
}
