package routes

import (
  "github.com/gin-gonic/gin"
  "../models"
)

// Função de criar os movimentos, recebe o ponteiro do Context do gin
func CreateMovements(c *gin.Context) {
  // Declara a variável que irá receber a lista de transactions
  var transactions models.TransactionList

  // Se a função que faz Bind do json não retorna nada (sinal que não deu erro)
  if c.BindJSON(&transactions) == nil {
    // Declara a variável que vai receber os movimentos
    var movements models.MovementList

    // For Each: Para cada transaction dentro de transactions
    // Sintaxe: for índice, valor := range lista... O _ significa que estou ignorando o índice
    for _, transaction := range transactions {

      // For, número da parcela começando em um, enquanto menor ou igual ao total, acrescenta
      for installment := 1; installment <= transaction.InstallmentsQuantity; installment++ {
        // Declara uma variável de Movement, lembrando que não é instância!
        movement := models.Movement{
          TransactionKey:     transaction.TransactionKey,
          InstallmentNumber:  installment,
          Amount:             transaction.TotalAmount / float32(transaction.InstallmentsQuantity),
          DueDate:            transaction.CaptureDate.AddDate(0, 0, installment * 30),
        }

        // Adiciona o elemento na lista de movimentos
        // O método append retorna a própria lista
        movements = append(movements, movement)
      }
    }

    // Manda o Context (equivalente ao Response do .NET) retornar 200 com os movimentos serializados em JSON
    // Passamos a referência para a variável dos movimentos em vez dos valores, já que não são instâncias
    c.JSON(200, &movements)
  }
}
