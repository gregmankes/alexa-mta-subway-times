package models

type SubwayStopResults struct {
	Results []SubwayStopResult `json:"result"`
}

type SubwayStopResult struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
