package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

type User struct{
    Name  string `json:"name"`
}

type Response struct{
    Greeting string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    var user User
    json.NewDecoder(req.Body).Decode(&user)
    response := Response{}
    response.Greeting = "Hello "+ user.Name +"! "
    fmt.Println(user.Name)
    json.NewEncoder(rw).Encode(response)
}

func main() {
    mux := httprouter.New()
    mux.POST("/hello/", hello)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}