![Asher_Duarte_Client_and_Server_Golang_code_script_database_7f9cb15f-a669-4508-a9a4-54293b156826 (1)](https://user-images.githubusercontent.com/105469529/213923352-4cd85230-85ea-4aa8-aeb0-e2a042f85d7f.png)
# Client-Server Golang
This code is a server written in Go language that consumes a JSON API to get the dollar-real exchange rate and stores it in a SQLite database. The server also has a simple endpoint that returns the current exchange rate.

## Dependencies
`"github.com/mattn/go-sqlite3"
"context"
"database/sql"
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
"time"`

## How it works
The code starts by opening a connection to a SQLite database using the `"database/sql"` package.
The main function calls the `"apiCotacao"` function, which makes a GET request to an API endpoint to get the exchange rate and returns it as a struct.
The `"insertCotacao"` function is called to store the exchange rate in the SQLite database.
The code creates a simple HTTP endpoint that returns the current exchange rate when accessed.
The ***server listens on port 8080***.

## Function Descriptions
`"apiCotacao"`: This function makes a GET request to an API endpoint to get the exchange rate and returns it as a struct.
`"insertCotacao"`: function stores the exchange rate in the SQLite database.
`"main"`:is the entry point of the program. It opens a connection to the database, calls the `"apiCotacao"` and `"insertCotacao"` functions, creates an HTTP endpoint, and starts the server.

### Note
The SQLite database and its table need to be set up before running the code.
The API endpoint used in the code is "[awesomeapi](https://economia.awesomeapi.com.br/json/last/USD-BRL)", which may not be available in the future.



