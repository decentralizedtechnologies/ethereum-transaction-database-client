package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/jinzhu/gorm"
	"gitlab.com/canya-com/canwork-database-client"
)

// Transaction describes the data structure for an Ethereum tx request
type Transaction struct {
	Hash string `gorm:"column:hash" json:"hash"`
	From string `gorm:"column:from" json:"from"`
}

// Table : gets database table instance
func (model *Transaction) Table() *gorm.DB {
	return database.Client.Table(database.TableTransaction)
}

// GetRecordByHash : gets transaction query
func (model *Transaction) GetRecordByHash(tx *Transaction) *gorm.DB {
	return model.Table().Where("hash = ?", model.Hash).First(tx)
}

// New : creates new transaction record
func (model *Transaction) New() *gorm.DB {
	return model.Table().Create(model)
}

// RecordExists : check if a tx record exists
func (model *Transaction) RecordExists() bool {
	return !model.Table().Where("hash = ?", model.Hash).First(&model).RecordNotFound()
}

// IsValid : validates transaction with go-ethereum utils
func (model *Transaction) IsValid() bool {
	_, err := hexutil.Decode(model.Hash)
	return err == nil
}

// Length : returns tx length as of go-ethereum
func (model *Transaction) Length() int {
	return common.HashLength
}
