# Comments API

## Configuration

You can add the variables in an `.env` file (copy the .env.exemple), or overwrite them in environment variables.


| Key                                      | Required | Type    | Default | Secret | Description                                                                   |
|------------------------------------------|----------|---------|---------|--------|-------------------------------------------------------------------------------|
| COMMENTS_DATABASE_ADDRESSS               |          | string  | empty   |        | DSN of the comments database                                                  |
| ACCESS_KEY                               |        X | string  |         |        | Key to access the API from an other server application                        |
| PORT_NUMBER                              |          | integer | 8080    |        | Port number to access to this service                                         |
| LOG_LEVEL                                |          | string  | INFO    |        | Level of log (i.e. TRACE, DEBUG, INFO, WARNING, ERROR, FATAL)                 |
| LOG_FILENAME                             |          | string  | empty   |        | Path the log file (otherwise, the logs are redirected to the standard output) |

## Run the service

### Run from command line

To run the service:
```
go run cmd/comments-api/main.go run
```
