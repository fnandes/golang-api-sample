package main

import (
  "github.com/gin-gonic/gin"
  "./routes"
)

// Main do programa, equivalente ao public void Main(string[] args)
func main() {
  // Inicia o router do gin
  r := gin.Default()

  // Cria o agrupamento das rotas
  // Tudo aqui dentro será /api/v1/xxxxxxxx
  v1 := r.Group("api/v1")
  {
    // Define a rota e passa a função que irá processá-la
    v1.POST("/createmovements", routes.CreateMovements)
  }

  // Roda a Api na porta definida
  r.Run(":8080")
}
