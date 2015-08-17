package types

import (
	"encoding/json"
)

type Shape struct {
	rows    int
	columns int
}

func NewShape(rows, columns int) *Shape {
	s := &Shape{
		rows:    rows,
		columns: columns,
	}

	return s
}

func (s *Shape) Rows() int {
	return s.rows
}

func (s *Shape) Columns() int {
	return s.columns
}

func (s *Shape) MarshalJSON() ([]byte, error) {
	jsonObject := shapeJson{
		Rows:    s.rows,
		Columns: s.columns,
	}

	return json.Marshal(&jsonObject)
}

func (s *Shape) UnmarshalJSON(b []byte) error {
	jsonObject := &shapeJson{}

	if err := json.Unmarshal(b, jsonObject); err != nil {
		return err
	}

	s.rows = jsonObject.Rows
	s.columns = jsonObject.Columns

	return nil
}

type shapeJson struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}
