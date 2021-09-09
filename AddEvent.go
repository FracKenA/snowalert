package main

import (
	"encoding/json"
	"fmt"
)

// AddEvent will add an event to the custom event table.
func AddEvent(incident *Incident) *Incident {
	var EventTable TenMEvents

	EventTable.ShortDescription = incident.ShortDescription
	// buildDescription := fmt.Sprintf("%s\n\n\n%s", incident.Description, incident.TransactionID)
	// EventTable.WorkNotes = buildDescription
	EventTable.CmdbCi = incident.CISysID
	EventTable.Description = incident.Description
	EventTable.ShortDescription = incident.ShortDescription
	EventTable.Incident = incident.IncidentSysID
	EventTable.CorrelationID = incident.CorrelationID
	EventTable.CiName = incident.CIName

	payload, err := json.Marshal(EventTable)
	if err != nil {
		incident.Err = append(incident.Err, fmt.Sprintf("AddEvent - JSON Marshal Error: %s", err))
	}

	// resultDetail, err := SNOWConnect.PUT(payload, "", "number", "", "", incident.IncidentSysID).IncidentDetail()
	resultDetail, err := SNOWConnect.POST(payload, "", "", "", "", "").CustomDetail("u_10mevents")
	if err != nil {
		incident.Err = append(incident.Err, fmt.Sprintf("AddEvent - PUT.IncidentDetail: %s", err))
	} else {
		incident.EventResults = resultDetail
	}

	return incident
}
