# URL Shortener

The goal of this exercise is to create an [http.Handler](https://golang.org/pkg/net/http/#Handler) that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for `/dogs` to `https://www.somesite.com/a-story-about-dogs` we would look for any incoming web requests with the path `/dogs` and redirect them.

## How to get this working on your own 
Simply clone the repo and start with go run main/main.go.

You may have to do go get for the yaml packages. 

Once application is started on 8080
just hit 
localhost:8080/big for example and this will redirect you to youtube.com 

## Implementation 
In my implementation both JSON and YAML format inputs are accepted. How it works is that the YAML handler is first called, if the input is of type YAML the byte data will be unmarshalled and mapped into URL and Path. The handler then redirects the url given the url. If the input is not of type YAML the handler then checks if it is of type JSON with a similiar process as above. Given that both formats fail you will be redirected to a page that says hello world

### This project was taken from the gopher excercises url shortener.
