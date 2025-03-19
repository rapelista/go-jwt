# go-jwt

A Go implementation for parsing and validating JSON Web Tokens (JWT). This tool allows you to decode and inspect JWT tokens easily.

## Features

- Decode JWT tokens to view their header, payload, and signature.
- Lightweight and easy to use.

## Prerequisites

Ensure you have the following installed on your system:

- [Go](https://golang.org/dl/) (version 1.16 or later)

## Build

To build the project, run the following command in the project directory:

```bash
go build
```

This will generate an executable file named `go-jwt`.

## Usage

Once built, you can use the tool to parse a JWT token by running:

```bash
./go-jwt [JWT_TOKEN]
```

Replace `[JWT_TOKEN]` with the actual JWT string you want to parse.

### Example

```bash
./go-jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

The output will display the decoded header, payload, and signature of the token.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the project.
