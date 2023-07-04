package models

import (
	"encoding"
	"encoding/json"
)

type WencaiStockData struct {
	Bull  int    `json:"bull"`
	Short string `json:"short"`
}

var _ encoding.BinaryMarshaler = new(WencaiStockData)
var _ encoding.BinaryUnmarshaler = new(WencaiStockData)

func (w *WencaiStockData) MarshalBinary() (data []byte, err error) {
	return json.Marshal(w)
}

func (w *WencaiStockData) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, w)

}

type WencaiHotStack struct {
	Market int    `json:"market"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Tag    struct {
		ConceptTag []string `json:"concept_tag"`
	} `json:"tag"`
	Order int `json:"order"`
}

type WencaiStackAnalyse struct {
	Result string `json:"result"`
}
