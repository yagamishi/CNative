### install
```
kubectl apply -f jaeger.yaml
kubectl edit configmap istio -n istio-system
set tracing.sampling=100
```
### deploy tracing
```
kubectl create ns httptracing
kubectl label ns httptracing istio-injection=enabled
kubectl -n httptracing apply -f httpserver.yaml
kubectl -n httptracing apply -f service0.yaml
kubectl -n httptracing apply -f service1.yaml
kubectl -n httptracing apply -f service2.yaml 
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=yagamishi Inc./CN=*' -keyout https.key -out https.crt
kubectl create -n istio-system secret tls https-credential --key=https.key --cert=https.crt
kubectl apply -f istio-specs.yaml -n httptracing
```
### check ingress ip
```
k get svc -nistio-system
istio-ingressgateway   LoadBalancer   $INGRESS_IP
```
### 
```
curl -k https://10.106.46.228/simple/hello
```

```
curl -k https://10.106.46.228/service0
```

