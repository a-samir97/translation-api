# Transcription Translation API

This project is a REST API built using GoLang and Gin framework that translates Arabic sentences in a transcription to English.

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
