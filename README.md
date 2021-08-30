# Simple HTTP directory fuzzer

## Introduction
A simple implementation of an HTTP directory fuzzer using goroutines.

## Usage
### Direct execution
```sh
go run functions.go dirfuzzer.go -u <URL> -w <WORDLIST> [-t <number of threads>]
```
Where:
-u Base url which will be tested. Example: `https://www.github.com`
-w Wordlist file where every line is a path. Example: `/api`
-t Number of concurrent goroutines (similar to threads in other languages)

### Compiling
```sh
go build -o binary/dir-fuzzer functions.go dirfuzzer.go
```

And can be **executed** as:
```sh
./binary/dir-fuzzer -u <URL> -w <WORDLIST> [-t <number of threads>]
```

## TODO
- Save results to a file