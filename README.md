# github-repos-stats

# watching repositories status

This branch is for CNCF repos in China.
- [Mutli-Clusters related projects](https://github.com/pacoxu/github-repos-stats/tree/multi-clusters)


<!--START_SECTION:github_repos-->
## The CNCF repos
| ID |                                   REPO                                   | STARS | UPDATEDAT  | CREATEDAT  | FORKSCOUNT |                                                                                                                     DESCRIPTIONS                                                                                                                     |
|----|--------------------------------------------------------------------------|-------|------------|------------|------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  1 | [merbridge](https://github.com/merbridge/merbridge)                      |   361 | 2022-05-25 | 2022-01-12 |         40 | Use eBPF to speed up your Service Mesh like crossing an Einstein-Rosen Bridge.                                                                                                                                                                       |
|  2 | [clusterpedia](https://github.com/clusterpedia-io/clusterpedia)          |   325 | 2022-05-19 | 2021-10-08 |         46 | The Encyclopedia of Kubernetes clusters                                                                                                                                                                                                              |
|  3 | [piraeus](https://github.com/piraeusdatastore/piraeus)                   |   310 | 2022-05-21 | 2019-12-05 |         35 | High Available Datastore for Kubernetes                                                                                                                                                                                                              |
|  4 | [piraeus-operator](https://github.com/piraeusdatastore/piraeus-operator) |   191 | 2022-05-23 | 2019-07-25 |         31 | The Piraeus Operator manages LINSTOR clusters in Kubernetes.                                                                                                                                                                                         |
|  5 | [kubernetes-lts](https://github.com/klts-io/kubernetes-lts)              |   189 | 2022-05-04 | 2021-07-16 |         19 | Kubernetes LTS(long term support)                                                                                                                                                                                                                    |
|  6 | [hwameistor](https://github.com/hwameistor/hwameistor)                   |   159 | 2022-05-26 | 2022-01-19 |         15 | Local Storage is one of HwameiStor components. It will provision the local LVM volume.                                                                                                                                                               |
|  7 | [spiderpool](https://github.com/spidernet-io/spiderpool)                 |   113 | 2022-05-26 | 2022-03-07 |         10 | ipam for kubernetes  https://spidernet-io.github.io/spiderpool/                                                                                                                                                                                      |
|  8 | [public-image-mirror](https://github.com/DaoCloud/public-image-mirror)   |    78 | 2022-05-24 | 2021-09-09 |         27 | 很多镜像都在国外。比如 gcr 。国内下载很慢，需要加速。                                                                                                                                                                                                |
|  9 | [local-disk-manager](https://github.com/hwameistor/local-disk-manager)   |    27 | 2022-05-21 | 2022-01-19 |          6 | Local Disk Manager is one of HwameiStor components. It will manage all the local disks of the HwameiStor nodes, including provision local Disk volume, and disk health management.                                                                   |
| 10 | [ckube](https://github.com/DaoCloud/ckube)                               |    14 | 2022-05-19 | 2022-03-17 |          1 | Kubernetes APIServer 高性能代理组件，代理 APIServer 的 List 请求，其它类型的请求会直接反向代理到原生 APIServer。 CKube 还额外支持了分页、搜索和索引等功能。 并且，CKube 100% 兼容原生 kubectl 和 kube client sdk，只需要简单的配置即可实现全局替换。 |



## Skipped repos
<!--END_SECTION:github_repos-->

# Build

docker build . -t github-repo-stats:v0.1.0

following https://github.com/yihong0618/github-readme-stats style.
