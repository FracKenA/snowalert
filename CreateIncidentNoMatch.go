package main

import (
	"encoding/json"
	"fmt"
	snow "github.com/FracKenA/gosnow"
	"strconv"
)

// CreateIncidentNoMatch creates in incident in ServiceNow when it doesn't have a matching CI
func CreateIncidentNoMatch(incident *Incident) *Incident {
	var incidentTable snow.IncidentTable

	incidentTable.CmdbCi = incident.CIName
	incidentTable.CorrelationID = incident.CorrelationID
	buildDescription := fmt.Sprintf("%s\n\n\n%s", incident.Description, incident.TransactionID)
	incidentTable.Description = buildDescription
	incidentTable.ShortDescription = incident.ShortDescription

	if incident.Category != "" {
		incidentTable.Category = incident.Category
	} else {
		incidentTable.Category = SNOWCategory
	}

	if incident.Impact != 0 {
		incidentTable.Impact = strconv.Itoa(incident.Impact)
	} else {
		incidentTable.Impact = strconv.Itoa(SNOWImpact)
	}

	if incident.Urgency != 0 {
		incidentTable.Urgency = strconv.Itoa(incident.Urgency)
	} else {
		incidentTable.Urgency = strconv.Itoa(SNOWUrgency)
	}

	if incident.AssignmentGroup != "" {
		incidentTable.AssignmentGroup = incident.AssignmentGroup
	} else {
		incidentTable.AssignmentGroup = SNOWAssignmentGroup
	}

	pl, err := json.Marshal(incidentTable)
	if err != nil {
		incident.Err = append(incident.Err, fmt.Sprintf("CreateIncident - json.Marshal %s", err))
	}
	results, err := SNOWConnect.POST(pl, "", "number,sys_id", "", "", "").IncidentDetail()
	if err != nil {
		incident.Err = append(incident.Err, fmt.Sprintf("CreateIncident - POST.CaseDetail %s", err))
	} else {
		incident.IncidentID = results.Result.Number
		incident.IncidentSysID = results.Result.SysID
	}

	return incident
}
