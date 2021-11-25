## cm.yaml:  创建配置文件.   
## pvc-log.yaml: RWX,使用的是rook-cephfs。作用：存储log的pvc.   
## pvc-www.yaml： RWX,使用的是rook-cephfs。作用：存储代码的pvc.   
   
## service-cluster.yaml和service-nodeport.yaml：暴露接口.   
## deployment.yaml：.   
###	initContainers：检查service有没有启动   
###	configMap：读取cm里面的配置文件   
###	resources：资源限制   
###	readinessProbe：探活   
###	volumeMounts：日志和代码挂盘   

## daemonset.yaml:   
###	使用fluentd_elasticsearch对日志进行处理   
###	同时限制挂载的盘符是只读