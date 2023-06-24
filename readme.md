
##### Example 01: One of the object not exists

```curl
curl --location 'localhost:8080/copy' \
--header 'Content-Type: application/json' \
--data '{
    "sourceBucket":"sample-asd-123",
    "objectKeys":["sample.txt","examTopics_2.png"],
    "destLocation": "/prefix11/prefix2/prefix3/",
    "destBucket":"sample-asd-234"
}'
```

```sh
{"msg":"Error occurred while copying: operation error S3: CopyObject, https response error StatusCode: 404, RequestID: HV02NEHWB5FGNAEE, HostID: q8ymJt1VndKZ0za2+7YSGXh0S6o4FtHhp6onJBe8QYSKsFIKhqokQkNW6Uo+TuaAr/8I2i1+94M=, api error NoSuchKey: The specified key does not exist."}
```



##### Example 02: All object exists without any prefix path to be nested in

```curl
curl --location 'localhost:8080/copy' \
--header 'Content-Type: application/json' \
--data '{
    "sourceBucket":"sample-asd-123",
    "objectKeys":["sample.txt","examTopics.png"],
    "destLocation": "",
    "destBucket":"sample-asd-234"
}'
```

```sh
{"msg":"Copied successfully"}


PS C:\Users\Tisan\Documents\workspace\Go Tutorials> aws s3 ls sample-asd-234 --recursive
2023-06-24 19:00:54     251535 examTopics.png
2023-06-24 19:00:54       1915 sample.txt
PS C:\Users\Tisan\Documents\workspace\Go Tutorials> 
```



##### Example 03: All object exists and to be nested on prefix path

```curl
curl --location 'localhost:8080/copy' \
--header 'Content-Type: application/json' \
--data '{
    "sourceBucket":"sample-asd-123",
    "objectKeys":["sample.txt","examTopics.png"],
    "destLocation": "prefix11/prefix2/prefix3",
    "destBucket":"sample-asd-234"
}'
```

```sh
{"msg":"Copied successfully"}


PS C:\Users\Tisan\Documents\workspace\Go Tutorials> aws s3 ls sample-asd-234 --recursive
2023-06-24 19:03:03     251535 prefix11/prefix2/prefix3/examTopics.png
2023-06-24 19:03:03       1915 prefix11/prefix2/prefix3/sample.txt
PS C:\Users\Tisan\Documents\workspace\Go Tutorials> 
```


##### Example 04: All object exists and to be nested on prefix path with extra slash

```curl
curl --location 'localhost:8080/copy' \
--header 'Content-Type: application/json' \
--data '{
    "sourceBucket":"sample-asd-123",
    "objectKeys":["sample.txt","examTopics.png"],
    "destLocation": "/prefix11/prefix2/prefix3/",
    "destBucket":"sample-asd-234"
}'
```

```sh
{"msg":"Copied successfully"}


PS C:\Users\Tisan\Documents\workspace\Go Tutorials> aws s3 ls sample-asd-234 --recursive
2023-06-24 19:17:20     251535 /prefix11/prefix2/prefix3//examTopics.png
2023-06-24 19:17:20       1915 /prefix11/prefix2/prefix3//sample.txt
PS C:\Users\Tisan\Documents\workspace\Go Tutorials> 
```



##### Example 05: Health check

```sh
curl -X GET localhost:8080/health

{"Status":"running"}
```
