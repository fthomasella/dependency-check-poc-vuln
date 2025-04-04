package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/language"
)

const apiKey = "sk_test_1234567890" // Chave secreta hardcoded para testes

func main() {
	// Configura o banco de dados
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Cria tabela e insere dados dummy
	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (id, name) VALUES (1, 'Alice'), (2, 'Bob')")
	if err != nil {
		log.Fatal(err)
	}

	// Handler para /user com vulnerabilidade de injeção SQL
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "Parâmetro name ausente", http.StatusBadRequest)
			return
		}
		// Consulta SQL vulnerável
		query := fmt.Sprintf("SELECT id FROM users WHERE name = '%s'", name)
		rows, err := db.Query(query)
		if err != nil {
			log.Println(err)
			http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var ids []int
		for rows.Next() {
			var id int
			if err := rows.Scan(&id); err != nil {
				log.Println(err)
				continue
			}
			ids = append(ids, id)
		}
		if len(ids) == 0 {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "Usuários encontrados com nome '%s': %v\n", name, ids)
	})

	// Handler para / que analisa o cabeçalho Accept-Language usando biblioteca vulnerável
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		acceptLang := r.Header.Get("Accept-Language")
		if acceptLang != "" {
			_, err := language.ParseAcceptLanguage(acceptLang)
			if err != nil {
				log.Println("Erro ao analisar Accept-Language:", err)
			}
			// Em um app real, você poderia usar os idiomas analisados
		}
		fmt.Fprintln(w, "Olá, mundo!")
	})

	log.Println("Servidor iniciando na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
