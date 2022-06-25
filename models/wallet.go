package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	WalletID      int    `json:"wallet_id"`
	NoPhone       string `json:"no_phone"`
	Amount        int    `json:"amount"`
	Amount_locked int    `json:"amount_locked"`
	IsLocked      int    `json:"amount_locked"`
	UserId        int    `json:"user_id"`
}
