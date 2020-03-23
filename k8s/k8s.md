#### `K8S`学习笔记

#### 主要组件
* __高可用的集群副本数最好是》=3 的奇数个__
* `etcd` 可信赖的分布式键值存储服务，能够为整个分布式服务存储一些关键数据【参数信息】，协助集群存储服务。
* `apiserver`: 所有服务访问的统一入口。
* `controller Manager`:维持副本期望数目。
* `Scheduler`：负责介接收任务，选择合适的节点进行任务分配。
* `Kubelet`直接跟容器引擎交互实现容器声明周期管理。
* `Kube-proxy` 负责写入规则至`iptables` `ipvs` 实现服务映射访问。
* `core-dns` 可以为集群中的`SVC`创建一个域名、`IP`的对应关系解析。 
* `Dashborad` 给k8s结构体系提供一个`B/S`访问体系。
* `Ingress Controller` 官方只能实现四层代理 Ingress 可以实现七层代理。
* `FedEtation` 提供一个可以跨集群中心多k8s的统一管理系统。 
* `Prometheus`  提供k8s 集群的监控能力
* `ELk` 提供`k8s`集群日志统一分析接入平台。 

#### `pod`概念
* 自主式pod
* 控制器管理的pod