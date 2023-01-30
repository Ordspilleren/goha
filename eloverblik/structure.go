package eloverblik

import "time"

type EloverblikResponse struct {
	Result []Result `json:"result,omitempty"`
}
type SenderMarketParticipantMRID struct {
	CodingScheme interface{} `json:"codingScheme,omitempty"`
	Name         interface{} `json:"name,omitempty"`
}
type PeriodTimeInterval struct {
	Start time.Time `json:"start,omitempty"`
	End   time.Time `json:"end,omitempty"`
}
type MRID struct {
	CodingScheme string `json:"codingScheme,omitempty"`
	Name         string `json:"name,omitempty"`
}
type MarketEvaluationPoint struct {
	MRID MRID `json:"mRID,omitempty"`
}
type TimeInterval struct {
	Start time.Time `json:"start,omitempty"`
	End   time.Time `json:"end,omitempty"`
}
type Point struct {
	Position            string `json:"position,omitempty"`
	OutQuantityQuantity string `json:"out_Quantity.quantity,omitempty"`
	OutQuantityQuality  string `json:"out_Quantity.quality,omitempty"`
}
type Period struct {
	Resolution   string       `json:"resolution,omitempty"`
	TimeInterval TimeInterval `json:"timeInterval,omitempty"`
	Point        []Point      `json:"Point,omitempty"`
}
type TimeSeries struct {
	MRID                  string                `json:"mRID,omitempty"`
	BusinessType          string                `json:"businessType,omitempty"`
	CurveType             string                `json:"curveType,omitempty"`
	MeasurementUnitName   string                `json:"measurement_Unit.name,omitempty"`
	MarketEvaluationPoint MarketEvaluationPoint `json:"MarketEvaluationPoint,omitempty"`
	Period                []Period              `json:"Period,omitempty"`
}
type MyEnergyDataMarketDocument struct {
	MRID                        string                      `json:"mRID,omitempty"`
	CreatedDateTime             time.Time                   `json:"createdDateTime,omitempty"`
	SenderMarketParticipantName string                      `json:"sender_MarketParticipant.name,omitempty"`
	SenderMarketParticipantMRID SenderMarketParticipantMRID `json:"sender_MarketParticipant.mRID,omitempty"`
	PeriodTimeInterval          PeriodTimeInterval          `json:"period.timeInterval,omitempty"`
	TimeSeries                  []TimeSeries                `json:"TimeSeries,omitempty"`
}
type Result struct {
	MyEnergyDataMarketDocument MyEnergyDataMarketDocument `json:"MyEnergyData_MarketDocument,omitempty"`
	Success                    bool                       `json:"success,omitempty"`
	ErrorCode                  int                        `json:"errorCode,omitempty"`
	ErrorText                  string                     `json:"errorText,omitempty"`
	ID                         string                     `json:"id,omitempty"`
	StackTrace                 interface{}                `json:"stackTrace,omitempty"`
}

type TokenResult struct {
	Result string `json:"result,omitempty"`
}
