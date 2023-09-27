This is README.md

 ## How to run

    1. create & fill the file .env
    2. make sure to download all dependencies with the command "go mod vendor"
    3. open the terminal and run with the command "go run main.go" then open the browser


    To shorten a URL
    http://localhost:{port}/shorten?url=<URL>&expiry=<RFC3339 Expiry Date>
    example: http://localhost:8080/shorten?url=https://exampleurl.com&expiry=2023-10-01T00:00:00Z (this will redirect to the original URL)

    Access the shortener URL
    http://localhost:{port}/:shortURL
    example: if the shortened URL is qwerty, access it at http://localhost:8080/qwerty

    Get the click count for a shortener URL
    http://localhost:{port}/:shortURL/clicks
    example: http://localhost:8080/qwerty/clicks

    To sort the shortener URLS by click count in asc or desc
    http://localhost:{port}/sort
    example: http://localhost:8080/sort?sortType=asc (this is will sorted by ascending)
