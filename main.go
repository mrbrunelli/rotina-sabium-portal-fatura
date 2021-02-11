package main

import (
	"database/sql"
	json "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", execRoutine)
	http.HandleFunc("/favicon.ico", doNothing)
	log.Println("---> Servidor http iniciado!")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

type response struct {
	Message string `json:"message"`
}

func execRoutine(w http.ResponseWriter, r *http.Request) {
	log.Println("---> Iniciando rotina...")
	db, err := openConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var res response
	err = db.QueryRow("SELECT public.fn_atualiza_fatura() AS message").Scan(&res.Message)
	if err != nil {
		log.Fatal(err)
	}
	j, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(j))
	log.Println("---> Rotina executada com sucesso!")
	log.Printf("---> Mensagem: %s \n \n", res.Message)
}

func doNothing(w http.ResponseWriter, r *http.Request) {}

func getDotEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func openConnection() (*sql.DB, error) {
	DBHOST := getDotEnv("DB_HOST")
	DBNAME := getDotEnv("DB_NAME")
	DBUSER := getDotEnv("DB_USER")
	DBPASS := getDotEnv("DB_PASS")
	DBPORT := getDotEnv("DB_PORT")
	DBSSLMODE := getDotEnv("DB_SSLMODE")
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s sslmode=%s", DBHOST, DBNAME, DBUSER, DBPASS, DBPORT, DBSSLMODE)
	return sql.Open("postgres", connStr)
}
