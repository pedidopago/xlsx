package xlsx

import "time"

type SFType string

const (
	SFString SFType = "inlineStr"
	SFNumber SFType = "n"
	SFDate   SFType = "d"
)

// SFCell is a cell used for streams
type SFCell struct {
	Value interface{}
	Type  SFType
}

func SFCellString(v string) *SFCell {
	return &SFCell{v, SFString}
}

func SFCellNumber(v string) *SFCell {
	return &SFCell{v, SFNumber}
}

func SFCellDate(v time.Time) *SFCell {
	return &SFCell{v, SFDate}
}
