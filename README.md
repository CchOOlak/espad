# Espad

This is a test project for [Espad](http://www.espad-co.com/).

GO Version: `1.20`


## Environment

You should edit `docker-compose.env`:

| Name                |                                   Description                                    | Default |
| :------------------ | :------------------------------------------------------------------------------: | :------ |
| ESPAD_ENV           | choice environment that you run the code (choices are `local`, `dev` and `prod`) | `local` |
| ESPAD_LOG_LEVEL     |             Specific log level (`debug`, `warning`, `error`, `info`)             | `info`  |
| ESPAD_LOG_TIMESTAMP |                          show timestamp in logs or not                           | `true`  |

## Run
First build:

    docker-compose build
then run:

    docker-compose up -d
you can check the logs like this:

    docker-compose logs -f app

## API

### Create new shortLink
- path:   `/add`
- method: `POST`
- request:
    ```jsonc
    {
        "username": <string>,
        "url": <urlString>,
    }
    ```
- response:
    ```jsonc
    {
        "original": <string>,
        "shorten": <string>,
        "username": <string>,
    }
    ```

### Access shortLink
- path:   `/c/:path`
- method: `GET`
- response:
    ```jsonc
    {
        "original": <string>,
        "shorten": <string>,
        "username": <string>,
    }
    ```

### Redirect by shortLink
- path:   `/r/:path`
- method: `GET`
