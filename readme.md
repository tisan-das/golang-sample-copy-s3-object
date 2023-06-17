curl --location 'localhost:8080/copy' \
--header 'Content-Type: application/json' \
--data '{
    "sourceBucket":"sample-asd-123",
    "objectKeys":["sample.txt","examTopics_2.png"],
    "destBucket":"sample-asd-234"
}'


{"msg":"Error occurred while copying: operation error S3: CopyObject, https response error StatusCode: 404, RequestID: 4RM8HCT1RSA4K68K, HostID: 780LyRMLcN+ThkmPC/t14L1ydFTY4wsqt+5TwnUMZ5nSO/WtTA28AAV3GQg9J6DWfPvj2zUkA9o=, api error NoSuchKey: The specified key does not exist."}



{
    "sourceBucket":"sample-asd-123",
    "objectKeys":["sample.txt","examTopics.png"],
    "destBucket":"sample-asd-234"
}

{"msg":"Copied successfully"}
