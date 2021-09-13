1. download an executable file.

2. run http server by bash call
    
    `clear; ./rtoacm` (for mac os, by default will be use port 9080)
    
    `clear; ./rtoacm -p=8080` (for mac os, will be use port 8080, it means that you can use any available port to run server)
    
    `clear; ./rtoacl` – for linux (arm64, by default will be use port 9080)
    
    `clear; ./rtoacl -p=8080` – for linux (arm64, will be use port 8080, it means that you can use any available port to run server)
             
3. make request to http server by bash call:

    `curl -X POST "127.0.0.1:9080" -H "Content-Type: application/json" -d '{"numeral": "M"}'`
    
    `127.0.0.1:9080` - if you run http server on different port then change `9080` on you value
    
    `{"numeral": "M"}` - if you want convert other roman numeral then change `M` for what do you want to convert 