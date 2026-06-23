# REST API Comparison

A project comparing the same REST API implemented in three different languages and frameworks: Python (FastAPI), TypeScript (Express), and Go (Gin). All three expose identical endpoints using in-memory storage, allowing for a direct comparison of syntax, code structure, typing systems, error handling, and performance.

## Project Structure

FINAL4102/

├── py-fastapi/        # Python implementation (port 8000)

├── ts-express/        # TypeScript implementation (port 3000)

├── go-gin/            # Go implementation (port 8080)

├── postman/           # Postman collection for testing

├── README.md          # Readme file

├── setup.bat          # Installs all dependencies

└── start.bat          # Launches all three servers

## Getting Started

Run the following scripts from the project root.

Install all dependencies:
setup.bat

Start all three servers:
start.bat

## Ports

| Language | Framework | Port |
|----------|-----------|------|
| Python   | FastAPI   | 8000 |
| TypeScript | Express | 3000 |
| Go       | Gin       | 8080 |

## Testing

Import the Postman collection from the `postman/` folder into Postman. It contains five requests for each language (Get All, Get by ID, Create, Update, Delete) pointed at the correct ports, with response time logging built into each request.

## Dependencies

Each implementation can also be run individually:

**Python**
cd py-fastapi

pip install -r requirements.txt

uvicorn main:app --reload

**TypeScript**
cd ts-express

npm install

npm start

**Go**
cd go-gin

go run main.go