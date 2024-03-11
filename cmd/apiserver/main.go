package main

import (
	"encoding/json"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
)

var (
	configPath string
	config     *Config
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	log.Println("START")
	flag.Parse()

	// Чтение конфиг файла
	config = NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Установка обработчика для POST запросов к эндпоинту /sessions
	http.HandleFunc(config.Route, handleSessions)
	log.Fatal(http.ListenAndServe(config.BindPort, nil))
}

func handleSessions(w http.ResponseWriter, r *http.Request) {
	log.Println("POST")
	// Разбор параметров POST-запроса
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		return
	}
	protection := r.Form.Get("protection")

	// Поиск значения в базе данных
	result := protection == config.SecretKey
	if !result {
		log.Println(" ")
		log.Println("protection", protection)
		log.Println("config.SecretKey", config.SecretKey)
		log.Println(" ")
	}

	// Отправка результата клиенту в формате JSON
	response := struct {
		Exists bool `json:"exists"`
	}{result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//_ "github.com/go-sql-driver/mysql"
// Подключение к базе данных MySQL
//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/corona_security?multiStatements=true")
//db, err := sql.Open("sqlite3", "corona_security.db")
//if err != nil {
//	log.Fatal(err)
//	return
//}
//defer db.Close()
