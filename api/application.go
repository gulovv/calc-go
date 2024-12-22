package api

import (
    "bufio"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "strings"
    "github.com/timurgulov/calc_go/internal/calculation"
)

type Config struct {
    Addr string
}

func ConfigFromEnv() *Config {
    config := new(Config)
    config.Addr = os.Getenv("PORT")
    if config.Addr == "" {
        config.Addr = "8080"
    }
    return config
}

type Application struct {
    config *Config
}

func New() *Application {
    return &Application{
        config: ConfigFromEnv(),
    }
}

func (a *Application) RunCLI() error {
    for {
        log.Println("Input expression:")
        reader := bufio.NewReader(os.Stdin)
        text, err := reader.ReadString('\n')
        if err != nil {
            log.Println("Failed to read expression from console:", err)
        }
        text = strings.TrimSpace(text)
        if text == "exit" {
            log.Println("Application successfully closed")
            return nil
        }
        result, err := calculation.Calc(text)
        if err != nil {
            log.Println(text, "calculation failed with error:", err)
        } else {
            log.Println(text, "=", result)
        }
    }
}

type Request struct {
    Expression string `json:"expression"`
}

// Response структура для вывода результата
type Response struct {
    Result float64 `json:"result,omitempty"`
    Error  string `json:"error,omitempty"`
}

// CalcHandler обрабатывает HTTP-запросы на вычисление
func CalcHandler(w http.ResponseWriter, r *http.Request) {
    request := new(Request)
    defer r.Body.Close()

    // Декодирование входных данных
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        response := Response{Error: "Invalid JSON format: " + err.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }

    if request.Expression == "" {
        response := Response{Error: "Expression cannot be empty. Please provide a valid mathematical expression."}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusUnprocessableEntity)
        json.NewEncoder(w).Encode(response)
        return
    }

    result, err := calculation.Calc(request.Expression)
    if err != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)
        response := Response{Error: "Expression is not valid. Please check the syntax."}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
        return
    }

    response := Response{Result: result}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}


func (a *Application) RunServer() error {
    http.HandleFunc("/api/v1/calculate", CalcHandler)
    log.Printf("Starting server on :%s\n", a.config.Addr)
    return http.ListenAndServe(":"+a.config.Addr, nil)
}

