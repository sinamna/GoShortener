# GoShortener
what is GoShortener? 
GoShortener is a very simple url shortener implemented using GraphQL 
## How to use?
GoShortener stores links in mongodb listening on localhost:27017 and the GraphQL server will be running on localhost:8080 (or you can specifiy port using env.PORT)

when you run the server it will lead you to GraphQL Playground where you can send and receive requests

each Link model consists of 2 parts , a long link and a short link . when you create a new Link it will create short link using sha256 hashing algorithm and save it to database


### How to create a link
you can create a shortened link using below mutation 
```
mutation{
  createLink(input:{longLink:"Enter the link"}){
    longLink
    shortLink
  }
}
```

### How to get a link
to get all links
```
{
	links{
    shortLink
    longLink
  }
}
```

to get specific link using short form
```
{
  link(shortLink:"short form"){
    longLink
  }
}
```
or long form
```
{
  link(longLink:"long form"){
    shortLink
  }
}
```
