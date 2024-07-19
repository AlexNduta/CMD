# README

## Task details:

Write a simple GO HTTP REST API that will accept a shell command and return the output of that command.
Create an endpoint with the POST method. api/cmd POST
Accept the command via query param or JSON body.
Return the output of the command as a response.

Plus point:
Return error if the command is not found, with proper status code.


# guide

## step1 : run the app
- on the root folder, run the `main` file which is a generated binary from the `main.go` file
```sh
./main
```

## step2 : make a POST request
- This can be done using `curl`, `postman` or any API testing tool
URL: `localhost:8080/api/cmd`

```sh
# testing with the `ls -la` command

curl -X POST http://localhost:8080/api/cmd -H "Content-Type: application/json" -d '{"Command": "ls -l"}'
```

