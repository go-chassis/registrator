# registrator
a tool to register service instance from service center
to another service center


# Guide

1. build images
```bash
bash ./build_image.sh
```

2. prepare config reg.yaml
```yaml
source:
  address: http://127.0.0.1:30100
  fetchInterval: 60s #how long to fetch services and instances and register them to target registry,
                  #if your instances is not changing rapidly, you can set it to longer
  exclude: SERVICECENTER,CseConfigCenter,REGISTRATOR
target:
  address: http://127.0.0.1:30200
  heartbeatInterval: 30s # how long to send heart beat to target registry
auth:
  accessKey:
  secretKey:
  project:
```

3. Run, CSE_REGISTRY_ADDR must equal to address of source
```bash
docker run -e CSE_REGISTRY_ADDR=http://127.0.0.1:30100 -v /path/to/conf_folder:/etc/registrator/ -p 8080:8080 gomesh/registrator 
```
