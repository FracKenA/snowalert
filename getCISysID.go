package main

import "fmt"

// GetCISysID looks up the SysID for a CI from the CI Name
func GetCISysID(incident *Incident) *Incident {

	var err error
	ciNameQuery := fmt.Sprintf("name=%s", incident.CIName)

	incident.CILookupResultDetail, incident.CIQueryDetails, err = SNOWConnect.GET("2", "name,sys_id", "", ciNameQuery, "").CmdbCi()
	if err != nil {
		incident.Err = append(incident.Err, fmt.Sprintf("GetCISysID - get ci sysid error: %s", err))
	} else if incident.CIQueryDetails.Results == 0 {
		incident.LogOutput = append(incident.LogOutput, fmt.Sprintf("GetCISysID - no results in cmdb_ci_server for %s\n", incident.CIName))
	} else if incident.CIQueryDetails.Results >= 2 {
		incident.LogOutput = append(incident.LogOutput, fmt.Sprintf("GetCISysID - too many results in cmdb_ci_server for %s\n", incident.CIName))
	} else if incident.CIQueryDetails.Results == 1 {
		incident.CISysID = incident.CILookupResultDetail.Result[0].SysID
		incident.CISysIDFound = true
		incident.LogOutput = append(incident.LogOutput, fmt.Sprintf("GetCISysID - Found HostName %s in cmdb_ci_server by hostName", incident.CIName))
	}
	if LogLevel == "DEBUG" {
		incident.LogOutput = append(incident.LogOutput, fmt.Sprintf("GetCISysID - query results: %T - query status: %v", incident.CILookupResultDetail, incident.CIQueryDetails))
	}
	return incident
}
