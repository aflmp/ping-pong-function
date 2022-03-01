AWS_ACCOUNT?=000000000000
AWS_REGION?=us-east-1
FUNCTION_NAME?=ping-pong-function
IAM_POLICY=${FUNCTION_NAME}-policy
IAM_ROLE=${FUNCTION_NAME}-role

.PHONY: create-iam-role
create-iam-role:
	@aws iam create-role \
        --role-name ${IAM_ROLE} \
        --assume-role-policy-document file://trust-policy.json > /dev/null
	@echo "- created iam role: "${IAM_ROLE}
	@sleep 5

.PHONY: add-iam-role-policy
add-iam-role-policy:
	@aws iam put-role-policy \
		--role-name ${IAM_ROLE} \
		--policy-name ${IAM_POLICY} \
		--policy-document file://permissions.json > /dev/null
	@echo "- added iam role policy: "${IAM_POLICY}
	@sleep 1

.PHONY: build
build: 
	@GOOS=linux go build main.go

.PHONY: zip
zip: 
	@zip deployment.zip main > /dev/null

.PHONY: create-lambda
create-lambda:
	@aws lambda create-function \
    	--region ${AWS_REGION} \
		--function-name ${FUNCTION_NAME} \
		--runtime go1.x \
		--handler main \
		--zip-file fileb://deployment.zip \
		--role arn:aws:iam::${AWS_ACCOUNT}:role/${IAM_ROLE} \
		--environment Variables="{VERSION=undefined, ENVIRONMENT=local}" > /dev/null
	@echo "- created lambda function: " ${FUNCTION_NAME}

.PHONY: clean
clean:
	@rm -rf main deployment.zip

.PHONY: deploy
deploy: clean create-iam-role add-iam-role-policy build zip create-lambda

.PHONY: delete-lambda
delete-lambda:
	@aws lambda delete-function \
    	--function-name ${FUNCTION_NAME}
	@echo "- deleted lambda function: " ${FUNCTION_NAME}

.PHONY: delete-iam-role-policy
delete-iam-role-policy:
	@aws iam delete-role-policy \
    	--role-name ${IAM_ROLE} \
    	--policy-name ${IAM_POLICY}
	@echo "- deleted iam role policy: "${IAM_POLICY}

.PHONY: delete-iam-role
delete-iam-role:
	@aws iam delete-role \
    	--role-name ${IAM_ROLE}
	@echo "- deleted iam role: "${IAM_ROLE}

.PHONY: destroy
destroy: delete-lambda delete-iam-role-policy delete-iam-role
