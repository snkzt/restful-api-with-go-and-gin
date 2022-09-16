package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

type album struct {
  ID     string `json:"id"`
  Title  string `json:"title"`
  Artist string `json:"artist"`
  Price  float64 `json:"price"`
}

var albums = []album{
  {ID: "1", Title: "Blue Moon", Artist: "Ela Fitzgerald", Price: 56.99},
  {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
  {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Assign the handler function to an endpoint path.
func main() {
  // getAlbumshandles requests to the /albums endpoint path
  // With Gin, you can associate a handler with an HTTP method-and-path.
  // In this way you can separately route requests route requests sent to a single path
  router := gin.Default()
  router.GET("/albums", getAlbums)
  router.POST("/albums", postAlbums)

  router.Run("localhost:8080")  
}


// getAlbums resopnds with the list of all albums as JSON
func getAlbums(ctx *gin.Context) {
  // This Context of Gin is different from the one in Go
  // It carries request details, validates and serialises JSON...
  // IndentedJSON serialise the struct albums into JSON and add it to the response.
  // The http status code is what you want to send to the client.
  // Also Contect.JSON is availabe to send simpler JSON but intended form is easier to 
  // debugging and so on. The size difference between the two is usually small.
    ctx.IndentedJSON(http.StatusOK, albums)
}

// postAlbums add albums data from JSON received in the request body.
func postAlbums(c *gin.Context) {
  var newAlbum album
  
  // BindJSON bind the received JSON to newAlbum
  err := c.BindJSON(&newAlbum)
  if err != nil {
    return
  } 

  // append newAlbum (album struct) initialised from the JSON to the albums slice
  albums = append(albums, newAlbum)
  // Send response of Created == 201 with JSON representing the newly added alubm newAlbum
  c.IndentedJSON(http.StatusCreated, newAlbum)
}


