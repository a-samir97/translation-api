# Transcription Translation API

This project is a REST API built using GoLang and Gin framework that translates Arabic sentences in a transcription to English.

### Project Structure
```
├── handlers            # Contains the handler logic for the API
│   └── translate.go    # The main controller handling the transcription translation
├──routers              # contains router for every API (URL)
    └── translate.go    # Contains translation routers (/translate/) 
├── models              # Contains the data models used in the API
│   └── transcription.go # Defines the structure for the transcription input and output
├── services            # Contains the business logic for translating sentences
│   └── translator.go   # The translation service where Arabic sentences are translated to English
├── tests               # Contains unit tests for verifying functionality
│   └── translate_test.go # Test cases for the translation functionality
├── go.mod              # Go module definition file
├── go.sum              # Go dependencies lock file
├── main.go             # The entry point of the application; initializes the Gin router
├── README.md           # Documentation for how to run and use the project
```

### Screen Shot
![Request 1](/docs/request1.png)
![Request 2](/docs/request2.png)

### Installation

1. Clone the repository:
    ```
    git clone https://github.com/a-samir97/translation-api.git
    ```
   
2. Install dependencies:
    ```
    go mod tidy
    ```
    
3. Add env file:
    ```
    touch .env
    OPENAI_KEY_API=#### 
    ```
    
4. Run the application:
    ```
    go run main.go
    ```

5. Test the API:
    Use Postman or curl to send a POST request to `http://localhost:8080/translate` with a JSON body like:
    ```json
    [
      {"speaker": "John", "time": "00:00:04", "sentence": "Hello Maria."},
      {"speaker": "Maria", "time": "00:00:09", "sentence": "صباح الخير"}
    ]
    ```

6. To run the tests:
    ```
    go test ./tests
    ```
