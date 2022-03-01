# ping-pong lambda function

- to deploy
```sh
AWS_ACCOUNT=123456789123 make deploy
```

- to destroy
```sh
AWS_ACCOUNT=123456789123 make destroy
```

- invoke lambda
```sh
aws lambda invoke \
    --function-name ping-pong-function \
    out.json
```