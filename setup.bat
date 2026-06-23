@echo off
echo Installing Python dependencies...
cd py-fastapi
python -m venv venv
call venv\Scripts\activate
pip install -r requirements.txt
deactivate
cd ..

echo Installing TypeScript dependencies...
cd ts-express
npm install
cd ..

echo Installing Go dependencies...
cd go-gin
go mod tidy
cd ..

echo Setup complete!
pause