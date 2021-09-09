package main

/*
// UpdateIncident will add a worknote to an existing incident.
func UpdateIncident(incident *Incident) *Incident {
	var IncidentTable snow.IncidentTable

	IncidentTable.ShortDescription = incident.ShortDescription
	buildDescription := fmt.Sprintf("%s\n\n\n%s", incident.Description, incident.TransactionID)
	IncidentTable.WorkNotes = buildDescription
	payload, err := json.Marshal(IncidentTable)
	if err != nil {
		incident.Err = append(incident.Err, fmt.Sprintf("UpdateIncident - JSON Marshal Error: %s", err))
	}

	resultDetail, err := SNOWConnect.PUT(payload, "", "number", "", "", incident.IncidentSysID).IncidentDetail()
	if err != nil {
		incident.Err = append(incident.Err, fmt.Sprintf("UpdateIncident - PUT.IncidentDetail: %s", err))
	} else {
		incident.IncidentID = resultDetail.Result.Number
	}

	return incident
}

 */
