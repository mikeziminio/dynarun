

## GRPC API

### Create model

```
grpcurl \
-plaintext \
-proto ./shared/proto/model/model.proto \
-d '{
  "name": "some new name",
  "repo_id": "some_repo",
  "filename": "some_file",
  "input_token_price": 10,
  "output_token_price": 10
}' \
127.0.0.1:8080 \
model.ModelService.CreateModel
```


### Update model

The `id` is required param.

```
grpcurl \
-plaintext \
-proto ./shared/proto/model/model.proto \
-d '{
  "id": "2b9ad1e4-d1a9-4c52-b4ae-2ca8c7d7b057",
  "name": "some new name 2"
}' \
127.0.0.1:8080 \
model.ModelService.UpdateModel
```


### Delete model

The `id` is required param.

```
grpcurl \
-plaintext \
-proto ./shared/proto/model/model.proto \
-d '{
  "id": "f29ec3bc-c099-42e9-aa3c-16e95df7a8ae"
}' \
127.0.0.1:8080 \
model.ModelService.DeleteModel
```


### List all models

```
grpcurl \
-plaintext \
-proto ./shared/proto/model/model.proto \
127.0.0.1:8080 \
model.ModelService.ListModel
```
