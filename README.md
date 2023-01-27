# Csv_To_String_Conversor_Golang

This is a project that will receive a csv(only numbers accepted in the spreadsheet and needs to be a square matrix) as input in the body of the request and convert it for in many ways, for example:
1. Csv converted to an array
2. Csv converted to only one line string with all records
3. Show the sum of the numbers in the spreadsheet

The response will be show as a json response in the endpoint

## Prerequisites

```
Docker
```

```
Make
```

### Getting started

1. Start the server

```
make start
```

2. Run all tests with test coverage

```
make run-all-tests
```

3. Stop the server

```
make stop
```

## How to consume the project routes:

- Show converted csv as array

```
POST http://localhost:9898/echo
```

```
Body:
```

```
--form 'file=@"/spreadsheet.csv"'
```

```
Response: 
```

```
{
    [
        "1,2",
        "3,4"
    ]
}
```

- Show converted csv in only one line

```
POST http://localhost:9898/echo?type=flatten
```

```
Body:
```

```
--form 'file=@"/spreadsheet.csv"'
```

```
Response: 
```

```
{
    "1,2,3,4"
}
```

- Show the sum of the records from the csv

```
POST http://localhost:9898/echo?type=sum
```

```
Body:
```

```
--form 'file=@"/spreadsheet.csv"'
```

```
Response: 
```

```
{
    10
}
```
