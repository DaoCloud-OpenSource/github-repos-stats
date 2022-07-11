# github-repos-stats

# watching repositories status

This branch is for CNCF repos in China.
- [Mutli-Clusters related projects](https://github.com/pacoxu/github-repos-stats/tree/multi-clusters)


<!--START_SECTION:github_repos-->
## The CNCF repos
| ID |                                   REPO                                   | STARS | UPDATEDAT  | CREATEDAT  | FORKSCOUNT |                                                                                                                     DESCRIPTIONS                                                                                                                     |
|----|--------------------------------------------------------------------------|-------|------------|------------|------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  1 | [merbridge](https://github.com/merbridge/merbridge)                      |   391 | 2022-07-11 | 2022-01-12 |         42 | Use eBPF to speed up your Service Mesh like crossing an Einstein-Rosen Bridge.                                                                                                                                                                       |
|  2 | [clusterpedia](https://github.com/clusterpedia-io/clusterpedia)          |   371 | 2022-07-09 | 2021-10-08 |         55 | The Encyclopedia of Kubernetes clusters                                                                                                                                                                                                              |
|  3 | [piraeus](https://github.com/piraeusdatastore/piraeus)                   |   316 | 2022-07-09 | 2019-12-05 |         36 | High Available Datastore for Kubernetes                                                                                                                                                                                                              |
|  4 | [spiderpool](https://github.com/spidernet-io/spiderpool)                 |   201 | 2022-07-05 | 2022-03-07 |         13 | kubernetes ipam                                                                                                                                                                                                                                      |
|  5 | [piraeus-operator](https://github.com/piraeusdatastore/piraeus-operator) |   201 | 2022-07-07 | 2019-07-25 |         32 | The Piraeus Operator manages LINSTOR clusters in Kubernetes.                                                                                                                                                                                         |
|  6 | [hwameistor](https://github.com/hwameistor/hwameistor)                   |   200 | 2022-07-06 | 2022-03-07 |          8 | HwameiStor system will be deployed by using Helm Charts, including Local Storage, Local Disk Manager, and Scheduler.                                                                                                                                 |
|  7 | [kubernetes-lts](https://github.com/klts-io/kubernetes-lts)              |   191 | 2022-06-29 | 2021-07-16 |         17 | Kubernetes LTS(long term support)                                                                                                                                                                                                                    |
|  8 | [cloudtty](https://github.com/cloudtty/cloudtty)                         |   172 | 2022-07-07 | 2022-04-28 |          8 | A Friendly Kubernetes CloudShell (Web Terminal) !                                                                                                                                                                                                    |
|  9 | [local-storage](https://github.com/hwameistor/local-storage)             |   165 | 2022-06-27 | 2022-01-19 |         16 | Local Storage is one of HwameiStor components. It will provision the local LVM volume.                                                                                                                                                               |
| 10 | [dao-2048](https://github.com/DaoCloud/dao-2048)                         |   162 | 2022-05-24 | 2015-06-03 |       2364 | 2048 is a number puzzle game.                                                                                                                                                                                                                        |
| 11 | [daochain](https://github.com/DaoCloud/daochain)                         |   119 | 2021-11-12 | 2016-11-02 |         30 | Docker image verification system based on Ethereum                                                                                                                                                                                                   |
| 12 | [dao-style](https://github.com/DaoCloud/dao-style)                       |   115 | 2022-07-10 | 2017-03-14 |         13 | 🎉 A high quality component library built on Vue.js 2.0                                                                                                                                                                                              |
| 13 | [public-image-mirror](https://github.com/DaoCloud/public-image-mirror)   |    96 | 2022-07-06 | 2021-09-09 |         29 | 很多镜像都在国外。比如 gcr 。国内下载很慢，需要加速。                                                                                                                                                                                                |
| 14 | [fake-kubelet](https://github.com/wzshiming/fake-kubelet)                |    48 | 2022-07-09 | 2020-07-20 |          7 | This is a fake kubelet. that can simulate any number of nodes and maintain pods on those nodes. It is useful for test control plane.                                                                                                                 |
| 15 | [local-disk-manager](https://github.com/hwameistor/local-disk-manager)   |    29 | 2022-06-22 | 2022-01-19 |          6 | Local Disk Manager is one of HwameiStor components. It will manage all the local disks of the HwameiStor nodes, including provision local Disk volume, and disk health management.                                                                   |
| 16 | [ckube](https://github.com/DaoCloud/ckube)                               |    15 | 2022-05-31 | 2022-03-17 |          2 | Kubernetes APIServer 高性能代理组件，代理 APIServer 的 List 请求，其它类型的请求会直接反向代理到原生 APIServer。 CKube 还额外支持了分页、搜索和索引等功能。 并且，CKube 100% 兼容原生 kubectl 和 kube client sdk，只需要简单的配置即可实现全局替换。 |
| 17 | [fake-k8s](https://github.com/wzshiming/fake-k8s)                        |    14 | 2022-07-04 | 2020-09-10 |          0 | fake-k8s is a tool for running Fake Kubernetes clusters, It can be used as an alternative to Kind in some scenarios where you don’t need to actually run the Pod                                                                                     |
| 18 | [ropee](https://github.com/DaoCloud/ropee)                               |     8 | 2022-03-27 | 2019-07-08 |          0 | A scalable prometheus remote storage adapter for splunk.                                                                                                                                                                                             |



## Skipped repos
<!--END_SECTION:github_repos-->

# Build

docker build . -t github-repo-stats:v0.1.0

following https://github.com/yihong0618/github-readme-stats style.
