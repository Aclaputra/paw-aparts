@echo off
:: Force the script to run from its own directory
cd /d "%~dp0"

echo Setting environment variables...
set GOOS=js
set GOARCH=wasm

echo Building game.wasm inside the web folder...
go build -o ./web/game.wasm main.go

if %ERRORLEVEL% EQU 0 (
    echo Build successful! 
    echo Starting Python server pointing to the 'web' folder...
    echo Open your browser to http://localhost:8080
    
    :: The -d flag tells Python to treat the 'web' folder as the root directory
    python -m http.server 8080 -d ./web
) else (
    echo Build failed! Check your Go code.
    pause
)