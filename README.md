# Installation

## Clone
```
git clone git@github.com:Fajar3108/simple-blog-api-go.git
```


## Setup

```
go mod download
wire gen simple-blog-api-golang
```

## Database

copy env file

```
cp .env.example .env
```

Setup your env file
example:
```
DB_USER=root
DB_PASSWORD=
DB_HOST=localhost
DB_PORT=3306
DB_NAME=simple_blog_api_go
```

create your database, make sure the name was exactly the same with your env setting

## Start Project

```
go run cmd/main.go
```

