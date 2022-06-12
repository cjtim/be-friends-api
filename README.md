## API

#### Prerequisite

- Docker
- Golang 1.17
- Make

#### Local run

* Start
    ```bash
    export LINE_CLIENT_ID=
    export LINE_SECRET_ID=
    export DATABASE_URL=postgresql://postgres:postgres@localhost:5432/be-friends?sslmode=disable

    make run
    ```
* Cleanup
    ```bash
    make clean
    ```
