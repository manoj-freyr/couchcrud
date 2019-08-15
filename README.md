# couchcrud
A small crud application using CouchBase and Golang

Uses couchbase's gocb.v1 SDK .
Before running the example we need to
1. Install CouchBase server.
2. Initialize the cluster running Couchbase . One can use cli or REST API or Web UI. With cli,
     couchbase-cli cluster-init -c 127.0.0.1 --cluster-username SuperRootUser --cluster-password dummypwd \
     --services data --cluster-ramsize 2048
     
3. Create a bucket with some name:
      couchbase-cli bucket-create -c 127.0.0.1:8091 --username SuperRootUser  --password dummypwd \
      --bucket travel-data --bucket-type couchbase --bucket-ramsize 1024
      
4. After connect to the cluster, we must authenticate using the uname and pwd as set.

