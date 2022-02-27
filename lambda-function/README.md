# ping-pong lambda function

## Steps to deploy
- create an iam role
```sh
aws iam create-role \
    --role-name ping-pong-function-role \
    --assume-role-policy-document file://trust-policy.json
```

- add an iam role policy
```sh
aws iam put-role-policy \
    --role-name ping-pong-function-role \
    --policy-name ping-pong-function-policy \
    --policy-document file://permissions.json
```

- deploy lambda for the first time
```sh
GOOS=linux go build main.go && \
zip deployment.zip main && \
aws lambda create-function \
    --region us-east-1 \
    --function-name ping-pong-function \
    --runtime go1.x \
    --handler main \
    --zip-file fileb://deployment.zip \
    --role arn:aws:iam::000000000000:role/ping-pong-function-role \
    --environment Variables="{VERSION=undefined, ENVIRONMENT=local}"
``` 

- update lambda function
```sh
GOOS=linux go build main.go && \
zip deployment.zip main && \
aws lambda update-function-code \
    --function-name ping-pong-function \
    --zip-file fileb://deployment.zip
```

- invoke lambda
```sh
aws lambda invoke \
    --function-name ping-pong-function \
    out.json
```

- cleanup
```sh
aws lambda delete-function \
    --function-name ping-pong-function
aws iam delete-role-policy \
    --role-name ping-pong-function-role \
    --policy-name ping-pong-function-policy
aws iam delete-role \
    --role-name ping-pong-function-role
```
