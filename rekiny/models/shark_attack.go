package models

type SharkAttack struct {
	Date                 string `json:"date"`
	Year                 string `json:"year"`
	Type                 string `json:"type"`
	Country              string `json:"country"`
	Area                 string `json:"area"`
	Location             string `json:"location"`
	Activity             string `json:"activity"`
	Name                 string `json:"name"`
	Sex                  string `json:"sex"`
	Age                  string `json:"age"`
	Injury               string `json:"injury"`
	FatalYN              string `json:"fatal_y_n"`
	Time                 string `json:"time"`
	Species              string `json:"species"`
	InvestigatorOrSource string `json:"investigator_or_source"`
	Pdf                  string `json:"pdf"`
	HrefFormula          string `json:"href_formula"`
	Href                 string `json:"href"`
	CaseNumber           string `json:"case_number"`
	CaseNumber0          string `json:"case_number0"`
	OriginalOrder        string `json:"original_order"`
}

type SharkAttacks []SharkAttack
