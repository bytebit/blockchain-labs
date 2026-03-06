@echo off
echo Starting P2P Network Demo...
echo.

echo Starting Node 1 on 127.0.0.1:8001 (seed node)
start "Node 1" cmd /k "go run main.go -addr 127.0.0.1:8001"

timeout /t 3 /nobreak >nul

echo Starting Node 2 on 127.0.0.1:8002, connecting to seed 127.0.0.1:8001
start "Node 2" cmd /k "go run main.go -addr 127.0.0.1:8002 -seeds 127.0.0.1:8001"

timeout /t 2 /nobreak >nul

echo Starting Node 3 on 127.0.0.1:8003, connecting to seed 127.0.0.1:8001
start "Node 3" cmd /k "go run main.go -addr 127.0.0.1:8003 -seeds 127.0.0.1:8001"

timeout /t 2 /nobreak >nul

echo Starting Node 4 on 127.0.0.1:8004, connecting to seeds 127.0.0.1:8001,127.0.0.1:8002
start "Node 4" cmd /k "go run main.go -addr 127.0.0.1:8004 -seeds 127.0.0.1:8001,127.0.0.1:8002"

echo.
echo All nodes started! Each node will print discovered peers every 5 seconds.
echo Press any key to exit (this won't stop the nodes)
pause >nul
