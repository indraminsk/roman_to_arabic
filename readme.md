1. run http server:

    a. if you have source files and don't want to build an executable file use next bash call in a folder with source code:
    
        `clear; go run *.go`
        
    b. if you have an executable file use next bash call in a folder with an executable file:
    
        `clear; ./rtoacm` – for mac os
        `clear; ./rtoacl` – for linux (arm64)
             
2. make request to http server

    `curl -X POST "127.0.0.1:9080" -H "Content-Type: application/json" -d '{"numeral": "M"}'`
    
    `127.0.0.1:9080` - if you run http server on different port then change `9080` on you value
    
    `{"numeral": "M"}` - if you want convert other roman numeral then change `M` for what do you want to convert 