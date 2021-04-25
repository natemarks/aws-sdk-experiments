# listS3Contents

This package builds an secetuatble that's run like this:
```bash

listS3Contents -b my_bucket_name
# prints out a list of the objects in the bucket
...

```

The exceutable is just a wrapper for the functions in pkg/s3.go.  These functions are sample code form AWS.
https://docs.aws.amazon.com/code-samples/latest/catalog/gov2-s3-ListObjects-ListObjectsv2.go.html