@echo off
echo Starting all servers...

start "Python FastAPI" cmd /k "cd py-fastapi && call venv\Scripts\activate && uvicorn main:app --reload"
start "TypeScript Express" cmd /k "cd ts-express && npm start"
start "Go Gin" cmd /k "cd go-gin && go run main.go"

echo All servers launching in separate windows!
pause