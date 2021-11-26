> cm.yaml:  创建配置文

> pvc-log.yaml: RWX,使用的是rook-cephfs。作用：存储log的pvc   
> pvc-www.yaml： RWX,使用的是rook-cephfs。作用：存储代码的pvc
   
> service-cluster.yaml和service-nodeport.yaml：暴露接口  

> deployment.yaml: 
>>	initContainers：检查service有没有启动 
>>	configMap：读取cm里面的配置文件   
>>	resources：资源限制   
>>	readinessProbe：探活   
>>	volumeMounts：日志和代码挂盘   

> daemonset.yaml:使用fluentd_elasticsearch对日志进行处理   同时限制挂载的盘符是只读

> ingress.yaml:设置ing

> img：执行结果图
>> POD相关：pod.png,deployment.png,daemonset.png
>> 挂盘：rook.png,pvc.png
>> 网络相关：svc.png，ing-svc.png,ing.png
>> 执行结果http,https： curl.png

> 访问路径：通过外部ip(node)+ingress端口(192.168.31.199:31046)-->nginx:80-->httpservice-clusterip:8888-->最后负载均衡的pod:8887