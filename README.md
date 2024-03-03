# pipeline


## Dependencies

- `go install honnef.co/go/tools/cmd/staticcheck@2023.1.7`
- `go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2`
- `go get -u github.com/spf13/cobra@v1.8.0`
- `go install github.com/goreleaser/goreleaser@v1.24.0`


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


## Release Process

1. Create an inital `.goreleaser.yaml` with `goreleaser init`
1. See what goreleaser provides as options `goreleaser release -h`
1. Test out goreleaser to create archives in `dist` folder `goreleaser release --snapshot --clean`
1. Setup github credentials `export GITHUB_TOKEN="YOUR_GH_TOKEN"`
1. Create your first tag `git tag -a v0.1.0 -m "First release" && git push origin v0.1.0`
1. Create a release `goreleaser release --clean`