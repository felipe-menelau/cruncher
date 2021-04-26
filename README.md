# Cruncher
Basic number cruncher, receives text as input and outputs text

## Usage

Import the lib and call `Crunch` on your code

```
var s string
result := cruncher.Crunch(s)
fmt.Printf("%s", result)
```
Make sure the string is in the format follow the examples in the `examples` folder.
You can also build `main.go` and pass a correctly formatted file as a CLI argument.
For example: 
```
go build main.go
```
Then
```
./main examples/example_input1.txt
```

## Tests
To run all unit tests
```
go test ./...
```
