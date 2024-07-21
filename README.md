<p align="center">
<a href="https://github.com/Dhruv9449">
	<img src="https://github.com/user-attachments/assets/436d2707-e79e-41d8-b1d0-fa2d2948307e" alt="Dhruv Shah" height=200/>
</a>
	<h2 align="center"> Tizori Backend 🔐</h2>
	<h4 align="center"> Backend API and admin panel for self-hosted credentials manager with role-based access control (RBAC).</h4>
</p>

<span align="center">
	
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54)
![Django](https://img.shields.io/badge/django-%23092E20.svg?style=for-the-badge&logo=django&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)

</span>

## Table of Contents
- [Key Features](#key-features)
- [Tech Stack and Dependencies](#tech-stack-and-dependencies)
	- [Deployment](#deployment)
	- [Web API](#web-api)
	- [Admin Panel](#admin-panel)
	- [Database](#database)
- [Setting-up and Installation](#setting-up-and-installation)
	- [Prerequisites](#prerequisites)
	- [Configure the environment variables](#configure-the-environment-variables)
- [Usage](#usage)
	- [Running the application](#running-the-application)
	- [Stopping the application](#stopping-the-application)
	- [Using management commands of the CLI application](#using-management-commands-of-the-cli-application)
- [Developer](#developer)


<br>

## Key Features
- [x] User Authenitication using Google Oauth2
- [x] Timetable detection from VTOP timetable page text
- [x] Add and remove friends using friend requests
- [x] User search to friend other users
- [x] Display friends' timetable
- [x] Show mutual friends
- [x] Friend suggestions based on mutual friends
- [x] CLI application to manage the web api

<br>
<br>

## Tech Stack and Dependencies
### Deployment
- [Docker](https://www.docker.com/)

### Web API
- **Language** - [Go](https://go.dev/)
- **Framework** - [Fiber](https://gofiber.io/)
- **ORM** - [Gorm](https://gorm.io/)
- **CLI** framework - [urfave/cli](https://cli.urfave.org/)
- **JWT** - [golang-jwt](https://golang-jwt.github.io/jwt/)
- **UUID** - [google/uuid](https://pkg.go.dev/github.com/google/uuid)
- **AES Encryption** - [crypto/aes](https://pkg.go.dev/crypto/aes)

### Admin Panel
- **Language** - [Python](https://www.python.org/)
- **Framework** - [Django](https://www.djangoproject.com/)
- **Admin Template** - [Jazzmin](https://django-jazzmin.readthedocs.io/)

### Database
- [PostgreSQL](https://www.postgresql.org/)  

<br>
<br>

## Setting-up and Installation  
### Prerequisites
- Download and install [Docker](https://docs.docker.com/get-docker/) and [Docker compose](https://docs.docker.com/compose/install/)
- Clone the repository using the following command -  
```bash
git clone https://github.com/Dhruv9449/tizori-backend.git
```

### Configure the environment variables
Configure the env files present in `/tizori/.env/`   
For local environment use `.local` and for production use `.production`    
<br> 
  
Following environment variables need to be configured -
- `FIBER_PORT` - Value of port used by the web api in the form `:<PORT>`, default value is `:3000`
- `DEBUG` - Set to `true` for local environment, `false` for production environment
- `POSTGRES_URL` - Set to `postgres://<POSTGRES_USER>:<POSTGRES_PASSWORD>@postgres:<POSTGRES_PORT>/<POSTGRES_DB>`
- `POSTGRES_USER` - Username for postgres database
- `POSTGRES_PASSWORD` - Password for postgres database
- `POSTGRES_DB` - Name of the postgres database
- `POSTGRES_HOST` - Hostname for postgres database, default value is `postgres`
- `POSTGRES_PORT` - Port for postgres database, default value is `5432`
- `JWT_SECRET` - JWT secret key that will be used to sign the tokens
- `AES_KEY` - AES key used for encryption and decryption of credentials

<br>
<br>


## Usage
### Running the application
Use the following command to run the application -  

#### MacOS and Linux
```zsh
./tizori.sh up
```

#### Windows
```cmd
tizori.cmd up
```

<br>

### Stopping the application
Use the following command to stop the application -

#### MacOS and Linux
```zsh
./tizori.sh down
```

#### Windows
```cmd
tizori.cmd down
```

<br>

### Using management commands of the CLI application
Use the following command to run the CLI application -

#### MacOS and Linux
```zsh
./tizori.sh cli <command>
```

#### Windows
```cmd
tizori.cmd cli <command>
```

<br>
<br>

## Developer

<table>
	<tr align="center">
		<td>
		Dhruv Shah
		<p align="center">
			<img src = "https://avatars.githubusercontent.com/u/88224695" width="150" height="150" alt="Dhruv Shah">
		</p>
			<p align="center">
				<a href = "https://github.com/Dhruv9449">
					<img src = "http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg" width="36" height = "36" alt="GitHub"/>
				</a>
				<a href = "https://www.linkedin.com/in/Dhruv9449" target="_blank">
					<img src = "http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg" width="36" height="36" alt="LinkedIn"/>
				</a>
				<a href = "mailto:dhruvshahrds@gmail.com" target="_blank">
					<img src = "https://www.iconninja.com/files/312/807/734/share-send-email-chat-circle-message-mail-icon.svg" width="36" height="36" 
					alt="Email"/>
				</a>
			</p>
		</td>
	</tr>
</table>

<p align="center">
	Made with ❤️ by <a href="https://github.com/Dhruv9449">Dhruv Shah</a>
</p>