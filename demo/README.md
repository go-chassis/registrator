# Demo Guide
1. start registries
```bash
cd source
sudo docker-compose up
```
```bash
cd target
sudo docker-compose up
```

2. build and run 
```bash
go install github.com/go-mesh/registrator
mv $GOPATH/bin/registrator ./
./registrator

```

source list is in 
http://127.0.0.1:30103/#!/sc/services/ 


target list is in 
http://127.0.0.1:30203/#!/sc/services/ 

you will see instance is same now