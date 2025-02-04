package hardware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/devices", devicesHandler)
	log.Println("Servidor iniciado na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func devicesHandler(w http.ResponseWriter, r *http.Request) {
	// Configuração CORS básica
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Validação simples do token
	if r.Header.Get("API-Key") != os.Getenv("API_KEY") {
		respondWithError(w, "Acesso não autorizado", http.StatusUnauthorized)
		return
	}

	// Conexão com o banco
	db, err := connectDB()
	if err != nil {
		respondWithError(w, "Erro no banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Rotas
	switch {
	case r.Method == "GET" && r.URL.Query().Get("user_id") != "":
		handleUserDevices(w, r, db)
	case r.Method == "GET" && r.URL.Query().Get("org_id") != "":
		handleOrgDevices(w, r, db)
	default:
		respondWithError(w, "Endpoint não encontrado", http.StatusNotFound)
	}
}

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	return sql.Open("mysql", dsn)
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
