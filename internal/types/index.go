package types

import (
	"encoding/json"
)

type Index struct {
	row    int
	column int
}

func NewIndex(row, column int) *Index {
	i := &Index{
		row:    row,
		column: column,
	}

	return i
}

func (i *Index) Row() int {
	return i.row
}

func (i *Index) Column() int {
	return i.column
}

func (i *Index) MarshalJSON() ([]byte, error) {
	jsonObject := indexJson{
		Row:    i.row,
		Column: i.column,
	}

	return json.Marshal(&jsonObject)
}

func (i *Index) UnmarshalJSON(b []byte) error {
	jsonObject := &indexJson{}

	if err := json.Unmarshal(b, jsonObject); err != nil {
		return err
	}

	i.row = jsonObject.Row
	i.column = jsonObject.Column

	return nil
}

type indexJson struct {
	Row    int `json:"rows"`
	Column int `json:"columns"`
}
