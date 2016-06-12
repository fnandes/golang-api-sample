package models

import "time"

type Transaction struct {
  TransactionKey        string        `json:transactionKey`
  CaptureDate           time.Time     `json:captureDate`
  TotalAmount           float32       `json:totalAmount`
  InstallmentsQuantity  int           `json:installmentsQuantity`
}

type TransactionList []Transaction
