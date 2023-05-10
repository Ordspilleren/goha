package goha

import "time"

type InputDatetime struct {
	HAEntity
}

func (i *InputDatetime) SetToNow() error {
	data := struct {
		Timestamp int `json:"timestamp,omitempty"`
	}{
		Timestamp: int(time.Now().Unix()),
	}
	i.integration.SendCommand(i, "set_datetime", data)
	return nil
}

func (i *InputDatetime) Timestamp() time.Time {
	time, _ := time.Parse("2006-01-02 15:04:05", i.State().State)
	return time
}
