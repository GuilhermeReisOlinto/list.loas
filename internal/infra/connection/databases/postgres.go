package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	host     string
	portStr  string
	user     string
	password string
	dbName   string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	host = os.Getenv("DB_HOST")
	portStr = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")

}

func Connection() (db *sql.DB) {

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Errorf("Erro ao converter porta para número: %v", err)
		return nil
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error in connect database", err)
		return nil // Retorna nulo para indicar uma falha na conexão
	}

	err = db.Ping()
	if err != nil {
		db.Close() // Fecha a conexão se ocorrer um erro no ping
		log.Fatal(err)
		fmt.Println("Error in connect database", err)
		return nil // Retorna nulo para indicar uma falha na conexão
	}

	fmt.Println("Successfully connected.")
	return db
}
