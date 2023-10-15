echo Create bucket
aws --endpoint-url=http://127.0.0.1:4566 s3api create-bucket --bucket mybucket

echo List buckets
aws --endpoint-url=http://127.0.0.1:4566 s3api list-buckets

echo Add ACL
aws --endpoint-url=http://127.0.0.1:4566 s3api put-bucket-acl --bucket mybucket --acl public-read