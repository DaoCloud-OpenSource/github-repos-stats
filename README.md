# github-repos-stats

# watching repositories status

This branch is for CNCF repos in China.
- [Mutli-Clusters related projects](https://github.com/pacoxu/github-repos-stats/tree/multi-clusters)


<!--START_SECTION:github_repos-->
## The CNCF repos
| ID |                                   REPO                                   | STARS | UPDATEDAT  | CREATEDAT  | FORKSCOUNT |                                                                                                                                    DESCRIPTIONS                                                                                                                                     |
|----|--------------------------------------------------------------------------|-------|------------|------------|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  1 | [clusterpedia](https://github.com/clusterpedia-io/clusterpedia)          |   488 | 2022-10-18 | 2021-10-08 |         72 | The Encyclopedia of Kubernetes clusters                                                                                                                                                                                                                                             |
|  2 | [merbridge](https://github.com/merbridge/merbridge)                      |   467 | 2022-10-19 | 2022-01-12 |         54 | Use eBPF to speed up your Service Mesh like crossing an Einstein-Rosen Bridge.                                                                                                                                                                                                      |
|  3 | [hwameistor](https://github.com/hwameistor/hwameistor)                   |   342 | 2022-10-17 | 2022-03-07 |         24 | Hwameistor is an HA local storage system for cloud-native stateful workloads.                                                                                                                                                                                                       |
|  4 | [piraeus](https://github.com/piraeusdatastore/piraeus)                   |   332 | 2022-10-13 | 2019-12-05 |         39 | High Available Datastore for Kubernetes                                                                                                                                                                                                                                             |
|  5 | [spiderpool](https://github.com/spidernet-io/spiderpool)                 |   250 | 2022-09-29 | 2022-03-07 |         31 | kubernetes ipam                                                                                                                                                                                                                                                                     |
|  6 | [cloudtty](https://github.com/cloudtty/cloudtty)                         |   246 | 2022-10-11 | 2022-04-28 |         23 | A Friendly Kubernetes CloudShell (Web Terminal) !                                                                                                                                                                                                                                   |
|  7 | [piraeus-operator](https://github.com/piraeusdatastore/piraeus-operator) |   215 | 2022-10-18 | 2019-07-25 |         34 | The Piraeus Operator manages LINSTOR clusters in Kubernetes.                                                                                                                                                                                                                        |
|  8 | [kubernetes-lts](https://github.com/klts-io/kubernetes-lts)              |   197 | 2022-10-19 | 2021-07-16 |         17 | Kubernetes LTS(long term support)                                                                                                                                                                                                                                                   |
|  9 | [kubean](https://github.com/kubean-io/kubean)                            |   167 | 2022-10-20 | 2022-07-05 |         20 | Kubernetes lifecycle management operator based on kubespray.                                                                                                                                                                                                                        |
| 10 | [dao-2048](https://github.com/DaoCloud/dao-2048)                         |   167 | 2022-10-06 | 2015-06-03 |       2357 | 2048 is a number puzzle game.                                                                                                                                                                                                                                                       |
| 11 | [public-image-mirror](https://github.com/DaoCloud/public-image-mirror)   |   129 | 2022-10-18 | 2021-09-09 |         41 | 很多镜像都在国外。比如 gcr 。国内下载很慢，需要加速。                                                                                                                                                                                                                               |
| 12 | [daochain](https://github.com/DaoCloud/daochain)                         |   119 | 2021-11-12 | 2016-11-02 |         30 | Docker image verification system based on Ethereum                                                                                                                                                                                                                                  |
| 13 | [dao-style](https://github.com/DaoCloud/dao-style)                       |   116 | 2022-08-31 | 2017-03-14 |         13 | 🎉 A high quality component library built on Vue.js 2.0                                                                                                                                                                                                                             |
| 14 | [kwok](https://github.com/kubernetes-sigs/kwok)                          |    97 | 2022-10-20 | 2022-07-28 |          7 | The repository is a toolkit that enables setting up a cluster of thousands of Nodes in seconds. Under the scene, all Nodes are simulated to behave like real ones, so the overall approach employes a pretty low resource footprint that you can easily play around on your laptop. |
| 15 | [ckube](https://github.com/DaoCloud/ckube)                               |    17 | 2022-09-19 | 2022-03-17 |          6 | Kubernetes APIServer 高性能代理组件，代理 APIServer 的 List 请求，其它类型的请求会直接反向代理到原生 APIServer。 CKube 还额外支持了分页、搜索和索引等功能。 并且，CKube 100% 兼容原生 kubectl 和 kube client sdk，只需要简单的配置即可实现全局替换。                                |
| 16 | [ropee](https://github.com/DaoCloud/ropee)                               |     8 | 2022-03-27 | 2019-07-08 |          0 | A scalable prometheus remote storage adapter for splunk.                                                                                                                                                                                                                            |



## Skipped repos
hwameistor/local-disk-manager
hwameistor/local-storage<!--END_SECTION:github_repos-->

# Build

docker build . -t github-repo-stats:v0.1.0

following https://github.com/yihong0618/github-readme-stats style.
