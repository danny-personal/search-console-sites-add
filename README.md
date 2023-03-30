# search-console-sites-add
This is the function that adds a site to Search Console.

## Create a project.
```shell
go mod init github.com/danny-personal/search-console-sites-add
go mod tidy
```

## Install.
```shell
sudo chmod -R 777 /go/
go get golang.org/x/oauth2/google
go get google.golang.org/api/webmasters/v3
```

## Environment.
```shell
export KEY_JSON=`cat credentials.json`
```

## Check the application.
```shell
vscode ➜ /workspaces/search-console-sites-add (main) $ curl -XPOST -d '{"store":"562"}' -H "Content-Type: application/json" localhost:8080 -i
HTTP/1.1 200 OK
Date: Thu, 30 Mar 2023 09:51:20 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

OKvscode ➜ /workspaces/search-console-sites-add (main) $
```
