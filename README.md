# github-repos-stats

# watching repositories status

This branch is for CNCF repos in China.
- [Mutli-Clusters related projects](https://github.com/pacoxu/github-repos-stats/tree/multi-clusters)


<!--START_SECTION:github_repos-->
## The CNCF repos
| ID |                                   REPO                                   | STARS | UPDATEDAT  | CREATEDAT  | FORKSCOUNT |                                                                                                                     DESCRIPTIONS                                                                                                                     |
|----|--------------------------------------------------------------------------|-------|------------|------------|------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  1 | [piraeus](https://github.com/piraeusdatastore/piraeus)                   |   303 | 2022-03-23 | 2019-12-05 |         34 | High Available Datastore for Kubernetes                                                                                                                                                                                                              |
|  2 | [merbridge](https://github.com/merbridge/merbridge)                      |   294 | 2022-03-30 | 2022-01-12 |         30 | Use eBPF to speed up your Service Mesh like crossing an Einstein-Rosen Bridge.                                                                                                                                                                       |
|  3 | [clusterpedia](https://github.com/clusterpedia-io/clusterpedia)          |   223 | 2022-03-30 | 2021-10-08 |         33 | The Encyclopedia of Kubernetes clusters                                                                                                                                                                                                              |
|  4 | [kubernetes-lts](https://github.com/klts-io/kubernetes-lts)              |   184 | 2022-03-22 | 2021-07-16 |         18 | Kubernetes LTS(long term support)                                                                                                                                                                                                                    |
|  5 | [piraeus-operator](https://github.com/piraeusdatastore/piraeus-operator) |   176 | 2022-03-24 | 2019-07-25 |         26 | The Piraeus Operator manages LINSTOR clusters in Kubernetes.                                                                                                                                                                                         |
|  6 | [dao-2048](https://github.com/DaoCloud/dao-2048)                         |   161 | 2022-03-18 | 2015-06-03 |       2371 | 2048 is a number puzzle game.                                                                                                                                                                                                                        |
|  7 | [daochain](https://github.com/DaoCloud/daochain)                         |   119 | 2021-11-12 | 2016-11-02 |         31 | Docker image verification system based on Ethereum                                                                                                                                                                                                   |
|  8 | [dao-style](https://github.com/DaoCloud/dao-style)                       |   116 | 2022-03-03 | 2017-03-14 |         13 | 🎉 A high quality component library built on Vue.js 2.0                                                                                                                                                                                              |
|  9 | [local-storage](https://github.com/hwameistor/local-storage)             |    91 | 2022-03-21 | 2022-01-19 |         13 | Local Storage is one of HwameiStor components. It will provision the local LVM volume.                                                                                                                                                               |
| 10 | [public-image-mirror](https://github.com/DaoCloud/public-image-mirror)   |    56 | 2022-03-24 | 2021-09-09 |         24 | 很多镜像都在国外。比如 gcr 。国内下载很慢，需要加速。                                                                                                                                                                                                |
| 11 | [spiderpool](https://github.com/spidernet-io/spiderpool)                 |    29 | 2022-03-29 | 2022-03-07 |          8 | ipam for kubernetes                                                                                                                                                                                                                                  |
| 12 | [local-disk-manager](https://github.com/hwameistor/local-disk-manager)   |    23 | 2022-03-18 | 2022-01-19 |          4 | Local Disk Manager is one of HwameiStor components. It will manage all the local disks of the HwameiStor nodes, including provision local Disk volume, and disk health management.                                                                   |
| 13 | [ckube](https://github.com/DaoCloud/ckube)                               |    11 | 2022-03-22 | 2022-03-17 |          0 | Kubernetes APIServer 高性能代理组件，代理 APIServer 的 List 请求，其它类型的请求会直接反向代理到原生 APIServer。 CKube 还额外支持了分页、搜索和索引等功能。 并且，CKube 100% 兼容原生 kubectl 和 kube client sdk，只需要简单的配置即可实现全局替换。 |
| 14 | [fake-kubelet](https://github.com/wzshiming/fake-kubelet)                |    11 | 2022-03-29 | 2020-07-20 |          4 | This is a fake kubelet. The pod on this node will always be in the ready state, but no process will be started.                                                                                                                                      |
| 15 | [ferry](https://github.com/ferry-proxy/ferry)                            |    10 | 2022-03-04 | 2021-10-18 |          1 | Ferry is a multi-cluster communication component of Kubernetes that supports mapping services from one cluster to another.                                                                                                                           |
| 16 | [fake-k8s](https://github.com/wzshiming/fake-k8s)                        |     5 | 2022-03-22 | 2020-09-10 |          0 | Run the fake k8s with docker-compose, It can be used as an alternative to Kind in some scenarios where you don’t need to actually run the Pod                                                                                                        |



## Skipped repos
<!--END_SECTION:github_repos-->

# Build

docker build . -t github-repo-stats:v0.1.0

following https://github.com/yihong0618/github-readme-stats style.
