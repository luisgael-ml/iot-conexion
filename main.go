package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func conectar(driver, dsn string) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("%s: %v", driver, err)
	}

	var version string
	db.QueryRow("SELECT version()").Scan(&version)
	fmt.Printf("✔ %s conectado: %s\n", driver, version)
}

func main() {
	// PostgreSQL: protocolo://usuario:contraseña@host:puerto/bd?parámetros
	conectar("pgx", "postgres://iot_user:iot_pass@localhost:5432/iot_analytics?sslmode=disable")

	// MySQL: usuario:contraseña@protocolo(host:puerto)/bd?parámetros
	conectar("mysql", "iot_user:iot_pass@tcp(localhost:3306)/iot_analytics?parseTime=true")
}