package model

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/canya-com/canwork-database-client"
)

const table = "transaction"

// Transaction describes the data structure for an Ethereum tx request
type Transaction struct {
	Hash string `gorm:"column:hash" json:"hash"`
	From string `gorm:"column:from" json:"from"`
}

// GetRecordByHash : gets transaction query
func (query *Transaction) GetRecordByHash(tx *Transaction) *gorm.DB {
	return database.Client.Table(table).Where("hash = ?", query.Hash).First(tx)
}
