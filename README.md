# pipeline


## Dependencies

- `go install honnef.co/go/tools/cmd/staticcheck@2023.1.7`
- `go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2`
- `go get -u github.com/spf13/cobra@v1.8.0`


## Initial setup

- Create sample json file `touch data.json`
- Add these contents to `data.json`
    ```json
    {
        "data": {
            "temperature": 76.6
        }
    }
    ```
- Create pipeline config `touch config.yaml`
- Add these contents to `config.yaml`
    ```yaml
    actions:
    -   data: ""
        cmd: cat data.json
        repeat: 3
    -   data: ""
        cmd: ls
        repeat: 1
    ```
- Run this command `go run main.go run`
- Expect this output
    ```bash
    {
        "data": {
            "temperature": 76.6
        }
    }

    {
        "data": {
            "temperature": 76.6
        }
    }

    {
        "data": {
            "temperature": 76.6
        }
    }

    Dockerfile
    Makefile
    README.md
    cmd
    config.yaml
    data.json
    go.mod
    go.sum
    main.go
    ```