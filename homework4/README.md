###步骤
```
*  安装helm，prometheus,grafana
    1. helm pull grafana/loki-stack
    2. tar -xvf loki-stack-2.4.1.tgz
    3. 修改目录下(包含子目录)`rbac.authorization.k8s.io/v1beta1` --》 `rbac.authorization.k8s.io/v1`
    4. helm upgrade --install loki ./loki-stack --set grafana.enabled=true,prometheus.enabled=true,prometheus.alertmanager.persistentVolume.enabled=false,prometheus.server.persistentVolume.enabled=false
* 修改service（loki-prometheus-server和loki-grafana）变成NodePort
* 编写代码
* k create -f deployment.yaml
* 查看prometheus,grafana
```
### prometheus访问地址
```
http://114.116.238.152:31547/
```
### grafana访问地址
```
user: admin
password: BAhGSZawXIUjHXsGWjCFfxIIUcayFPr3UmNdO0TK
```
