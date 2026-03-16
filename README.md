# Features
1. payloads.go file stored on the tool itself . [ OK ] [ Advanced Payload set ]
2. Advance Rate Limit Implemented . [ OK ]
3. Concurrent Request Sending .



# Current Updating status 
1. advance payloads in payloads.go 
2. try to compille payloads with the compiled tool or use payloads inside go file . [ OK ]
3. need to scan from requests.txt file as well . [ many problems inside this ] ..
4. add static url or something , that will not modified .
5. Live animated Logo and outfit while this tool start .
6. This tool should run until it finished  use all type of attack even it bypass , and also use combination attacks ..
7. concurrent request send .







# Bypasser Tool

## Features
- Path manipulation
- Header injection
- -- content-length variation
- -- user-agent rotation
- -- extra header with values
- Path capitalization techniques
- Randomized Capitalization Path 
- Method Manipulation
- Extension Manipulation [ from /sso to /sso.php , /sso.html ]
- Static path 
- Advance Rate limit feature implement .
- Request from request.txt file
- Proxy enabled 



## Structure
- `payloads/`: Editable payloads
- `internal/`: Core logic
  - `bypass/`: Combines techniques
  - `techniques/`: Encapsulates each bypass technique
  - `utils/`: HTTP client, logging, and rate-limiting

## Usage


```bash
echo "http://example.com" | bypasser -verbose

```bash
echo "http://example.com" | bypasser -proxy http://127.0.0.1:8080

```bash
bypasser -request request.txt

```bash
bypasser -request request.txt -proxy http://127.0.0.1:8080


