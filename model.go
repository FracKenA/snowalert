package main

import snow "github.com/FracKenA/gosnow"

type Incident struct {
	AssignmentGroup            string                    `json:"assignment_group,omitempty"`
	Caller                     string                    `json:"caller,omitempty"`
	Category                   string                    `json:"category,omitempty"`
	CILookupResultDetail       snow.CmdbCiResultsArray   `json:"CI_Result_Details,omitempty"`
	CIName                     string                    `json:"ci_name,omitempty"`
	CIQueryDetails             snow.Result               `json:"CI_Query_Results,omitempty"`
	CISysID                    string                    `json:"ci_sys_id,omitempty"`
	CISysIDFound               bool                      `json:"ci_sys_id_found"`
	CorrelationID              string                    `json:"correlation_id,omitempty"`
	Description                string                    `json:"description,omitempty"`
	Err                        []string                  `json:"errors,omitempty"`
	Impact                     int                       `json:"impact,omitempty"`
	IncidentID                 string                    `json:"number,omitempty"`
	IncidentLookupResultDetail snow.IncidentResultsArray `json:"Incident Result_Details,omitempty"`
	IncidentQueryDetails       snow.Result               `json:"Incident_Query_Results,omitempty"`
	IncidentSysID              string                    `json:"event_id,omitempty"`
	IncidentExists             bool                      `json:"json_exists,omitempty"`
	LogOutput                  []string                  `json:"logs,omitempty"`
	ShortDescription           string                    `json:"short_description,omitempty"`
	State                      int                       `json:"existing_state,omitempty"`
	Source                     string                    `json:"source_alert,omitempty"`
	TransactionID              string                    `json:"transaction_id,omitempty"`
	Urgency                    int                       `json:"urgency,omitempty"`
	EventResults               map[string]interface{}    `json:"event_results,omitempty"`
}
