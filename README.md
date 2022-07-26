This is an HTTP Server which serves a hello-world response


1. Build:
   1. build code with go
   `go build -o hello main.go`
   2. move it to a docker image(alpine)
   
   

2. Test:
   1. run and check for response. It should give a response with `Hello world in the response`
         ` curl http://localhost:8080/`
3. Deploy:
   1. Push the container to registry
      1. docker tag
      2. docker push
   2. create a deployment and a service
      1. manifest-files
      
      `kubectl create -f deployment.yaml`
   
      `kubectl create -f service.yaml`
      2. Don't forget to change the image version in the file
      3. test whether it is live