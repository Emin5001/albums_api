# Albums API 
This program is a simple API to retrieve data and information about a set of albums. 
It doesn't contain a database; the data is in the `main.go` file, in the `albums` slice.

> This program was based on the Go tutorial made by the Google team; it consisted of 3 different API endpoints, and I personally added other endpoints as well.

The goal of this project is to gain more experience in Go programming, and increase my knowledge in writing API's. I will be transitioning into an actual database with a wide variety of albums. 

# How to run
1. Clone this repository to your personal computer.
2. Run `go run .` on your terminal inside of this repository to start the server.
3. Make your API calls! An example of an API call to get the cheapest album would be 
`curl http://localhost:8080/albums/getCheapestAlbum`.

Have fun!