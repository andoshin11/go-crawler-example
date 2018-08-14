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
