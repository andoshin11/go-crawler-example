This is a crawler example with Go.

# Setup
1. Install deps

```
$ dep ensure
```

2. Add auth info

We need some auth info (including tokens) to communicate with Firestore.

Download `secret.json` from GCP IAM and save it on the project root directory.

# Run

```
$ go run main.go
```

# Crawl

```
$ curl http://localhost:8080/crawl/items
```

# Using Docker

You can also run app using Docker.

**!warning!**

In this way, you don't need to install deps by yourself, yet still need to download the `secret.json`.

```
$ docker build -t go-crawler-example:1.0 .
$ docker run -it -p 8080:8080 go-crawler-example:1.0
```
