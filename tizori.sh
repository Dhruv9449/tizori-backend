#!/bin/sh

# Tizori
# Handy set of commands to run to get a new server up and running
if [ "$1" = "local" ]; then
    shift # Discard the first argument
    environment="production"
    file="./tizori/production.yml"
else
    environment="local"
    file="./tizori/local.yml"
fi
command=$1

if [ -z "$command" ]; then
    echo 
    echo "     ███████████████╗                                                               "                
    echo "    ███╔══════════███╗                                                              "     
    echo "   ███╔╝           ███╗      ████████╗██╗███████╗ ██████╗ ██████╗ ██╗               "      
    echo "   ███║            ███║      ╚══██╔══╝██║╚══███╔╝██╔═══██╗██╔══██╗██║               "                       
    echo "   ███║            ███║         ██║   ██║  ███╔╝ ██║   ██║██████╔╝██║               "                            
    echo "█████████████████████████╗      ██║   ██║ ███╔╝  ██║   ██║██╔══██╗██║               "                          
    echo "█████████████████████████║      ██║   ██║███████╗╚██████╔╝██║  ██║██║               "                     
    echo "██████████╔═════█████████║      ╚═╝   ╚═╝╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝               "                  
    echo "██████████║     █████████║                                                          "                 
    echo "████████████╗ ███████████║      Developer: Dhruv Shah                               "
    echo "████████████║ ███████████║      Github: https://github.com/Dhruv9449                "
    echo "████████████║ ███████████║      Repository: https://github.com/Dhruv9449/tizori-cli "        
    echo "█████████████████████████║                                                          "      
    echo "█████████████████████████║                                                          "    
    echo "╚════════════════════════╝                                                          "  
    echo        
    echo "Environment: $environment"
    echo
    echo "Usage: tizori [command]"
    echo
    echo "Available commands:"
    echo "  up: Start the server"
    echo "  down: Stop the server"
    echo "  restart: Restart the server"
    echo "  cli: Run a command inside the container"
    echo "  seed: Seed the database"
    echo "  logs: Show the logs of the container"
    exit 1
fi

# Start server command
if [ "$command" = "up" ]; then
    echo "Starting server"
    docker compose -f "$file" up -d --build
    exit 1
fi

# Stop server command
if [ "$command" = "down" ]; then
    echo "Stopping server"
    docker compose -f "$file" down
    exit 1
fi

# Restart server command
if [ "$command" = "restart" ]; then
    echo "Restarting server"
    docker compose -f "$file" down
    docker compose -f "$file" up -d --build
    exit 1
fi

# Show logs command
if [ "$command" = "logs" ]; then
    echo "Showing logs"
    docker compose -f "$file" logs -f
    exit 1
fi

# Management commands
if [ "$command" = "cli" ]; then
    shift # Discard the first argument
    docker compose -f "$file" run --rm tizori-api ./bin/tizori "$@"
    exit 1
fi

if [ "$command" = "seed" ]; then
    echo "Seeding database"
    docker compose -f "$file" run --rm tizori-api ./bin/tizori seed
    exit 1
fi