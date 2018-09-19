package model

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/canya-com/canwork-database-client"
)

// TransactionLabel describes the data structure for an Ethereum tx request
type TransactionLabel struct {
	Hash string `gorm:"column:hash" json:"hash"`
	From string `gorm:"column:from" json:"from"`
}

// Table : gets database table instance
func (model *TransactionLabel) Table() *gorm.DB {
	return database.Client.Table(database.TableTransactionLabel)
}

// GetRecordByHash : gets transaction_label query
func (model *TransactionLabel) GetRecordByHash(tx *TransactionLabel) *gorm.DB {
	return model.Table().Where("hash = ?", model.Hash).First(tx)
}

// New : creates new transaction_label record
func (model *TransactionLabel) New() *gorm.DB {
	return model.Table().Create(model)
}

// RecordExists : check if a tx record exists
func (model *TransactionLabel) RecordExists() bool {
	return model.Table().Where("hash = ?", model.Hash).First(&TransactionLabel{}).RecordNotFound()
}
