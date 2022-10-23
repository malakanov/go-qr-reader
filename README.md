
# QR reader api

QR Code Reader service by Golang and a scanning library zxing


## Features

- read from file
- read from url


## Supported image format types

JPEG, PNG, BMP, TIFF, WebP


## API Reference

#### Get from file

```http
  POST /api/file
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `file` | `binary` | **Required**. QR image file |

#### Get from url of file

```http
  POST /api/url
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `url`      | `string` | **Required**. QR image url |



## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`POST`

`HOST`


## Run Locally

Clone the project

```bash
  git clone https://github.com/malakanov/go-qr-reader
```

Go to the project directory

```bash
  cd go-qr-reader
```

Create .env with environment variables

```bash
  touch .env
```

Install dependencies

```bash
  go get ./...
```
Go to the cmd directory

```bash
  cd cmd
```

Start the server

```bash
  go run .
```

