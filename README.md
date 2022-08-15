# github-repos-stats

# watching repositories status

This branch is for CNCF repos in China.
- [Mutli-Clusters related projects](https://github.com/pacoxu/github-repos-stats/tree/multi-clusters)


<!--START_SECTION:github_repos-->
## The CNCF repos
| ID |                                   REPO                                   | STARS | UPDATEDAT  | CREATEDAT  | FORKSCOUNT |                                                                                                                             DESCRIPTIONS                                                                                                                             |
|----|--------------------------------------------------------------------------|-------|------------|------------|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  1 | [clusterpedia](https://github.com/clusterpedia-io/clusterpedia)          |   438 | 2022-08-15 | 2021-10-08 |         61 | The Encyclopedia of Kubernetes clusters                                                                                                                                                                                                                              |
|  2 | [merbridge](https://github.com/merbridge/merbridge)                      |   421 | 2022-08-15 | 2022-01-12 |         51 | Use eBPF to speed up your Service Mesh like crossing an Einstein-Rosen Bridge.                                                                                                                                                                                       |
|  3 | [piraeus](https://github.com/piraeusdatastore/piraeus)                   |   325 | 2022-08-14 | 2019-12-05 |         37 | High Available Datastore for Kubernetes                                                                                                                                                                                                                              |
|  4 | [hwameistor](https://github.com/hwameistor/hwameistor)                   |   261 | 2022-08-15 | 2022-03-07 |         17 | Hwameistor is an HA local storage system for cloud-native stateful workloads.                                                                                                                                                                                        |
|  5 | [spiderpool](https://github.com/spidernet-io/spiderpool)                 |   238 | 2022-08-13 | 2022-03-07 |         21 | kubernetes ipam                                                                                                                                                                                                                                                      |
|  6 | [piraeus-operator](https://github.com/piraeusdatastore/piraeus-operator) |   206 | 2022-08-01 | 2019-07-25 |         32 | The Piraeus Operator manages LINSTOR clusters in Kubernetes.                                                                                                                                                                                                         |
|  7 | [cloudtty](https://github.com/cloudtty/cloudtty)                         |   199 | 2022-08-13 | 2022-04-28 |         18 | A Friendly Kubernetes CloudShell (Web Terminal) !                                                                                                                                                                                                                    |
|  8 | [kubernetes-lts](https://github.com/klts-io/kubernetes-lts)              |   194 | 2022-08-08 | 2021-07-16 |         17 | Kubernetes LTS(long term support)                                                                                                                                                                                                                                    |
|  9 | [local-storage](https://github.com/hwameistor/local-storage)             |   165 | 2022-08-06 | 2022-01-19 |         17 | [DEPRECATED] The code has been move to https://github.com/hwameistor/hwameistor.  Local Storage is one of HwameiStor components. It will provision the local LVM volume.                                                                                             |
| 10 | [dao-2048](https://github.com/DaoCloud/dao-2048)                         |   163 | 2022-07-21 | 2015-06-03 |       2363 | 2048 is a number puzzle game.                                                                                                                                                                                                                                        |
| 11 | [daochain](https://github.com/DaoCloud/daochain)                         |   119 | 2021-11-12 | 2016-11-02 |         30 | Docker image verification system based on Ethereum                                                                                                                                                                                                                   |
| 12 | [dao-style](https://github.com/DaoCloud/dao-style)                       |   115 | 2022-07-10 | 2017-03-14 |         13 | 🎉 A high quality component library built on Vue.js 2.0                                                                                                                                                                                                              |
| 13 | [public-image-mirror](https://github.com/DaoCloud/public-image-mirror)   |   104 | 2022-08-06 | 2021-09-09 |         35 | 很多镜像都在国外。比如 gcr 。国内下载很慢，需要加速。                                                                                                                                                                                                                |
| 14 | [kwok](https://github.com/kubernetes-sigs/kwok)                          |    44 | 2022-08-15 | 2022-07-28 |          4 | Simulate thousands of fake kubelets, on a laptop with minimum resource footprint.                                                                                                                                                                                    |
| 15 | [local-disk-manager](https://github.com/hwameistor/local-disk-manager)   |    28 | 2022-08-06 | 2022-01-19 |          8 | [DEPRECATED] The code has been move to https://github.com/hwameistor/hwameistor . Local Disk Manager is one of HwameiStor components. It will manage all the local disks of the HwameiStor nodes, including provision local Disk volume, and disk health management. |
| 16 | [ckube](https://github.com/DaoCloud/ckube)                               |    15 | 2022-05-31 | 2022-03-17 |          3 | Kubernetes APIServer 高性能代理组件，代理 APIServer 的 List 请求，其它类型的请求会直接反向代理到原生 APIServer。 CKube 还额外支持了分页、搜索和索引等功能。 并且，CKube 100% 兼容原生 kubectl 和 kube client sdk，只需要简单的配置即可实现全局替换。                 |
| 17 | [ropee](https://github.com/DaoCloud/ropee)                               |     8 | 2022-03-27 | 2019-07-08 |          0 | A scalable prometheus remote storage adapter for splunk.                                                                                                                                                                                                             |



## Skipped repos
<!--END_SECTION:github_repos-->

# Build

docker build . -t github-repo-stats:v0.1.0

following https://github.com/yihong0618/github-readme-stats style.
