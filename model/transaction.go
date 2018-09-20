package model

import (
	"database/sql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/jinzhu/gorm"
	"gitlab.com/canya-com/canwork-database-client"
)

var (
	// DefaultNetwork : mainnet
	DefaultNetwork = "mainnet"
)

const (
	statusPending = "pending"
	statusFailed  = "failed"
	statusTimeout = "timeout"
	statusSuccess = "success"
)

// StatusPending : get tx status pending
func (model *Transaction) StatusPending() string { return statusPending }

// StatusFailed : get tx status failed
func (model *Transaction) StatusFailed() string { return statusFailed }

// StatusTimeout : get tx status timeout
func (model *Transaction) StatusTimeout() string { return statusTimeout }

// StatusSuccess : get tx status success
func (model *Transaction) StatusSuccess() string { return statusSuccess }

// Transaction describes the data structure for an Ethereum tx request
type Transaction struct {
	Hash             string         `gorm:"column:hash" json:"hash,omitempty"`
	From             string         `gorm:"column:from" json:"from,omitempty"`
	Status           sql.NullString `gorm:"column:status" json:"status,omitempty"`
	IsWebhookCalled  int8           `gorm:"column:is_webhook_called" json:"isWebhookCalled,omitempty"`
	Network          string         `gorm:"column:network" json:"network,omitempty"`
	CreatedAt        int64          `gorm:"column:created_at" json:"createdAt,omitempty"`
	CompletedAt      int64          `gorm:"column:completed_at" json:"completedAt,omitempty"`
	Timeout          int64          `gorm:"column:timeout" json:"timeout,omitempty"`
	WebhookOnSuccess string         `gorm:"column:webhook_on_success" json:"webhookOnSuccess,omitempty"`
	WebhookOnTimeout string         `gorm:"column:webhook_on_timeout" json:"webhookOnTimeout,omitempty"`
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
