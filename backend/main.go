package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func main() {
  _, err := gorm.Open(
    "postgres",
    fmt.Sprintf(
      "host=%s port=%d user=%s password=%s sslmode=disable",
      "127.0.0.1", 5432, "postgres", "gql-example", "",
    ),
  )
  if err != nil {
    log.Fatalln(err)
  }

  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", welcome())

  err = e.Start(":3000")
  if err != nil {
    log.Fatalln(err)
  }
}

func welcome() echo.HandlerFunc {
  return func(c echo.Context) error {
    return c.String(http.StatusOK, "Welcome!")
  }
}
