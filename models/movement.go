package models

import "time"

// Define a estrutura de dados
type Movement struct {
  TransactionKey      string      `json:transactionKey`
  InstallmentNumber   int         `json:"installmentNumber"`
  Amount              float32     `json:"amount"`
  DueDate             time.Time   `json:"dueDate"`
}

// As estruturas de listas também precisam ser declaradas
type MovementList []Movement
