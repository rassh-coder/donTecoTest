## Required conf
1 - Create .env file in root directory <br>
2 - Fill .env like .env.example
## Migration Info
1 - For up migration use - `go run ./cmd/migration/main.go -up` in terminal <br>
2 - For down migration use - `go run ./cmd/migration/main.go -down` in terminal <br>
## API Info
`POST /employee/get-list` <br>
For all list - don't fill body
```
{ 
    "limit": 1,
    "offset: 1,
}
```
`POST /employee/get-by-name`
```
{
    "name": "name"
}
```