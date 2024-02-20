# Form Buffer
Partial Form Submission handler.

This microservice API is designed to handle partial form submissions, capturing and storing user forms. It submits only the most recent entry to the specified endpoint, as defined by the URL environment variable.
Requests are to be sent as a POST to the "/" URL pathway

## How to setup
### 1. Install Golang v1.22
Install Golang through the following link https://go.dev/doc/install

### 2. Clone the Repository
```
git clone https://github.com/eFlink/form-buffer.git
```

### 3. Download Dependencies
```
go mod tidy
```

### 4. Build the Project
```
go build
```

### 5. Run the Project
```
go run .
```