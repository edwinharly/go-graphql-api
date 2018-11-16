# go-graphql-api

## Getting started

Fork this repository and clone it to your local machine, then run `govendor install` to fetch all the required external packages

Create a `.env` file on the root containing your postgre configurations, e.g.
```
DBHOST=postgre-host
DBPORT=postgre-port
DBNAME=your-db-name
DBUSER=your-username
DBPASS=your-password
```

Run `go build` to build the binary then run `go-graphql-api` to start the server
