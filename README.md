# go-web-server
A very basic golang webserver implemention

To run this application you have to install the mux Router by running the following command

```
go get github.com/gorilla/mux
```
To test the webservices you have to provide the values in a valid JSON format. 

## addTwoNumber
```
body: { "number1": 1.2, "number2": 2.2 }
-> result: 3
```
## squareNumber
```
body: { "number1": 4.0 }
-> result: 16
```
## divideNumber
```
body: { "number1": 10, "number2": 5 }
-> result: 2
```
## subtract
```
body: { "number1": 1, "number2": 2 }
-> result: -1
```
## modulo
```
body: { "number1": 7, "number2": 4 }
-> result: 3
```
## multiply
```
body: { "number1": 10, "number2": 4 }
-> result: 40
```
## factorial
```
body: { "number1": 7}
-> result: 5040
```
