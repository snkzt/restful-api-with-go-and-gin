# restful-api-with-go-and-gin

### WHAT
- Implement basics of writing a RESTful web service API with Go and Gin Web Framework.
- Following [this](https://go.dev/doc/tutorial/web-service-gin) official doc.

### HOW
1. Clone [this](https://github.com/snkzt/restful-api-with-go-and-gin) repo to your local
2. Run ```$ go.run .``` below in your terminal at the root directory of the repo is cloned
3. Open another terminal screen and run below and see what will happen 
   ```
   To see what albums are there
   $ curl http://localhost:8080/albums

   To find an album with the ID
   $ curl http://localhost:8080/albums/${id}

   To add your favourite album(s)
   $ curl http://localhost:8080/albums \
       --include\
       --header "Content-Type: application/json" \
       --request "POST" \
       --data '{"id": "${some-number}","title": "${your-favourite-album-title}","artist": "${your-favourite-artist}","price":${price-of the-album}}'
   ``` 
