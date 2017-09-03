# FizzBuzz HTTP server

[![Build Status](https://api.travis-ci.org/msornay/lbc.svg)](http://travis-ci.org/msornay/lbc)

Serves FizzBuzz lists on /

## Example usage

* Install the package:
```
go get github.com/msornay/lbc
```

* Start the server:
```
lbc --addr :8070
```

* Query the server:
```
curl -H "Content-Type: application/json" \
     -d '{"string1":"fizz","string2":"buzz","int1":3,"int2":5,"limit":16}' \
     localhost:8070
```
