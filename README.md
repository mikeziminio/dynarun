
```

grpcurl \
-plaintext \
-proto ./shared/proto/model/model.proto \
-d '{
  name: "some name",
  repo_id: "some_repo",
  filename: "some_file",
  input_token_price: 0,
  output_token_price: 0,
}' \
127.0.0.1:8080 \
model.ModelService.CreateModel

```
