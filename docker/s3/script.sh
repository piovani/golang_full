#!/bin/bash

echo Create bucket
awslocal s3api create-bucket --bucket $AWS_BUCKET --region $AWS_DEFAULT_REGION

echo Add ACL
awslocal s3api put-bucket-acl --bucket $AWS_BUCKET --acl public-read-write

echo ADD Bucket Policy
awslocal s3api put-bucket-policy --policy file://bucket-policy.json --bucket my-bucket 

echo List buckets
awslocal s3api list-buckets
