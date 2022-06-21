# github-repos-stats

# watching repositories status

This branch is for CNCF repos in China.
- [Mutli-Clusters related projects](https://github.com/pacoxu/github-repos-stats/tree/multi-clusters)


<!--START_SECTION:github_repos-->
## The CNCF repos
| ID |                                   REPO                                   | STARS | UPDATEDAT  | CREATEDAT  | FORKSCOUNT |                                                                                                                     DESCRIPTIONS                                                                                                                     |
|----|--------------------------------------------------------------------------|-------|------------|------------|------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  1 | [merbridge](https://github.com/merbridge/merbridge)                      |   379 | 2022-06-20 | 2022-01-12 |         40 | Use eBPF to speed up your Service Mesh like crossing an Einstein-Rosen Bridge.                                                                                                                                                                       |
|  2 | [clusterpedia](https://github.com/clusterpedia-io/clusterpedia)          |   350 | 2022-06-19 | 2021-10-08 |         54 | The Encyclopedia of Kubernetes clusters                                                                                                                                                                                                              |
|  3 | [piraeus](https://github.com/piraeusdatastore/piraeus)                   |   312 | 2022-06-20 | 2019-12-05 |         35 | High Available Datastore for Kubernetes                                                                                                                                                                                                              |
|  4 | [piraeus-operator](https://github.com/piraeusdatastore/piraeus-operator) |   198 | 2022-06-16 | 2019-07-25 |         32 | The Piraeus Operator manages LINSTOR clusters in Kubernetes.                                                                                                                                                                                         |
|  5 | [kubernetes-lts](https://github.com/klts-io/kubernetes-lts)              |   189 | 2022-05-04 | 2021-07-16 |         18 | Kubernetes LTS(long term support)                                                                                                                                                                                                                    |
|  6 | [hwameistor](https://github.com/hwameistor/hwameistor)                   |   165 | 2022-06-21 | 2022-03-07 |          5 | HwameiStor system will be deployed by using Helm Charts, including Local Storage, Local Disk Manager, and Scheduler.                                                                                                                                 |
|  7 | [local-storage](https://github.com/hwameistor/local-storage)             |   162 | 2022-06-18 | 2022-01-19 |         16 | Local Storage is one of HwameiStor components. It will provision the local LVM volume.                                                                                                                                                               |
|  8 | [spiderpool](https://github.com/spidernet-io/spiderpool)                 |   152 | 2022-06-13 | 2022-03-07 |         12 | ipam for kubernetes  https://spidernet-io.github.io/spiderpool/                                                                                                                                                                                      |
|  9 | [cloudtty](https://github.com/cloudtty/cloudtty)                         |   114 | 2022-06-21 | 2022-04-28 |          6 | A Friendly Kubernetes CloudShell (Web Terminal) !                                                                                                                                                                                                    |
| 10 | [public-image-mirror](https://github.com/DaoCloud/public-image-mirror)   |    89 | 2022-06-16 | 2021-09-09 |         28 | 很多镜像都在国外。比如 gcr 。国内下载很慢，需要加速。                                                                                                                                                                                                |
| 11 | [fake-kubelet](https://github.com/wzshiming/fake-kubelet)                |    38 | 2022-06-21 | 2020-07-20 |          7 | This is a fake kubelet. that can simulate any number of nodes and maintain pods on those nodes. It is useful for test control plane.                                                                                                                 |
| 12 | [local-disk-manager](https://github.com/hwameistor/local-disk-manager)   |    28 | 2022-06-02 | 2022-01-19 |          6 | Local Disk Manager is one of HwameiStor components. It will manage all the local disks of the HwameiStor nodes, including provision local Disk volume, and disk health management.                                                                   |
| 13 | [ckube](https://github.com/DaoCloud/ckube)                               |    15 | 2022-05-31 | 2022-03-17 |          1 | Kubernetes APIServer 高性能代理组件，代理 APIServer 的 List 请求，其它类型的请求会直接反向代理到原生 APIServer。 CKube 还额外支持了分页、搜索和索引等功能。 并且，CKube 100% 兼容原生 kubectl 和 kube client sdk，只需要简单的配置即可实现全局替换。 |
| 14 | [fake-k8s](https://github.com/wzshiming/fake-k8s)                        |    10 | 2022-06-14 | 2020-09-10 |          0 | fake-k8s is a tool for running Fake Kubernetes clusters, It can be used as an alternative to Kind in some scenarios where you don’t need to actually run the Pod                                                                                     |



## Skipped repos
<!--END_SECTION:github_repos-->

# Build

docker build . -t github-repo-stats:v0.1.0

following https://github.com/yihong0618/github-readme-stats style.
