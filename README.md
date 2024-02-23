# gocalc 

Evaluating arbitrary C-like artithmetic/string expressions concurrently and provide rest api



## How to use
We can use j param to control the routine numbers
 
```json
POST http://localhost:8080/evaluate
Content-Type: application/json

{
    "exp": {
        "a":"b+c",
        "b":"23423",
        "c":"234",
        "d":"a-34/c",
        "e":"d+324*c",
        "G":"a+e",
        "F":"2048 - b",
        "x":"y-90",
        "y":"365"
    },
    "config": {
      "j":3
    }
}
```