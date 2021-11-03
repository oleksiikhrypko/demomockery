package models

type DataRec struct {
	ID  int    `db:"id" json:"id"`
	Val string `db:"val" json:"val"`
}
