# document-service

### Build the Docker Image

```bash
docker build -t document-service .
```


### Run the container
```bash
docker run -p 8080:8080 --rm document-service
```

### Testing the API Endpoints

#### Creating a document

```bash
curl -X POST http://localhost:8080/document/create \
  -H "Content-Type: application/json" \
  -d '{
        "id": "1",
        "name": "Test Document",
        "description": "This is a test"
      }'
```

#### Getting all documents

```bash
curl http://localhost:8080/documents
```


#### Get document by ID

```bash
curl "http://localhost:8080/document/get?id=1"
```


#### Delete a document

```bash
curl -X DELETE "http://localhost:8080/document/delete?id=1"
```

#### Search documents

```bash
curl "http://localhost:8080/document/search?q=test"
```

