@echo off

REM Tizori
REM Handy set of commands to run to get a new server up and running
if "%1" == "local" (
    shift
    set environment=production
    set file=.\tizori\production.yml
) else (
    set environment=local
    set file=.\tizori\local.yml
)
set command=%1

if "%command%" == "" (
    echo(
    echo      ███████████████╗                                                               
    echo     ███╔══════════███╗                                                              
    echo    ███╔╝           ███╗      ████████╗██╗███████╗ ██████╗ ██████╗ ██╗               
    echo    ███║            ███║      ╚══██╔══╝██║╚══███╔╝██╔═══██╗██╔══██╗██║               
    echo    ███║            ███║         ██║   ██║  ███╔╝ ██║   ██║██████╔╝██║               
    echo █████████████████████████╗      ██║   ██║ ███╔╝  ██║   ██║██╔══██╗██║               
    echo █████████████████████████║      ██║   ██║███████╗╚██████╔╝██║  ██║██║               
    echo ██████████╔═════█████████║      ╚═╝   ╚═╝╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝               
    echo ██████████║     █████████║                                                           
    echo ████████████╗ ███████████║      Developer: Dhruv Shah                               
    echo ████████████║ ███████████║      Github: https://github.com/Dhruv9449                
    echo ████████████║ ███████████║      Repository: https://github.com/Dhruv9449/tizori-cli 
    echo █████████████████████████║                                                           
    echo █████████████████████████║                                                           
    echo ╚════════════════════════╝                                                           
    echo. 
    echo Environment: %environment%
    echo.
    echo Usage: tizori [command]
    echo.
    echo Available commands:
    echo   up: Start the server
    echo   down: Stop the server
    echo   restart: Restart the server
    echo   cli: Run a command inside the container
    echo   logs: Show the logs of the container
    exit /b 1
)

REM Start server command
if "%command%" == "up" (
    echo Starting server
    docker compose -f "%file%" up -d --build
    exit /b 1
)

REM Stop server command
if "%command%" == "down" (
    echo Stopping server
    docker compose -f "%file%" down
    exit /b 1
)

REM Restart server command
if "%command%" == "restart" (
    echo Restarting server
    docker compose -f "%file%" down
    docker compose -f "%file%" up -d --build
    exit /b 1
)

REM Show logs command
if "%command%" == "logs" (
    echo Showing logs
    docker compose -f "%file%" logs -f
    exit /b 1
)

REM Management commands
if "%command%" == "cli" (
    shift
    docker compose -f "%file%" run --rm tizori-api ./bin/tizori %*
    exit /b 1
)