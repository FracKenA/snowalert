package main

import (
	"fmt"
)

func LookupIncident(incident *Incident) *Incident {

	q := fmt.Sprintf("active=true^correlation_id=%s", incident.CorrelationID)
	
	var err error
	incident.IncidentLookupResultDetail, incident.IncidentQueryDetails, err = SNOWConnect.GET("", "sys_id", "", q, "").Incident()
	if err != nil {
		incident.Err = append(incident.Err, fmt.Sprintf("LookupIncidentNoSysID - GET.Incident: %s ", err))
	}
	if incident.IncidentQueryDetails.Results >= 1 {
		if incident.IncidentQueryDetails.Results >= 2 {
			incident.LogOutput = append(incident.LogOutput, fmt.Sprintf("LookupIncidentNoSysID - Too Many Open Tickets for %s. Using first result.\n", incident.CISysID))
			incident.IncidentSysID = incident.IncidentLookupResultDetail.Result[0].SysID
			incident.IncidentExists = true


		} else {
			incident.IncidentSysID = incident.IncidentLookupResultDetail.Result[0].SysID
			incident.IncidentExists = true

		}
	} else {
		incident.IncidentSysID = ""
		incident.IncidentExists = false

	}

	if LogLevel == "DEBUG" {
		incident.LogOutput = append(incident.LogOutput, fmt.Sprintf("LookupIncident - Incident Sys ID  - %s", incident.IncidentID))
	}
	
	return incident
}
