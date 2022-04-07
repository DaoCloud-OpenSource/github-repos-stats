# github-repos-stats

# watching repositories status

This branch is for CNCF repos in China.
- [Mutli-Clusters related projects](https://github.com/pacoxu/github-repos-stats/tree/multi-clusters)


<!--START_SECTION:github_repos-->
## The CNCF repos
| ID |                                   REPO                                   | STARS | UPDATEDAT  | CREATEDAT  | FORKSCOUNT |                                                                                                                     DESCRIPTIONS                                                                                                                     |
|----|--------------------------------------------------------------------------|-------|------------|------------|------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  1 | [merbridge](https://github.com/merbridge/merbridge)                      |   305 | 2022-04-06 | 2022-01-12 |         31 | Use eBPF to speed up your Service Mesh like crossing an Einstein-Rosen Bridge.                                                                                                                                                                       |
|  2 | [piraeus](https://github.com/piraeusdatastore/piraeus)                   |   304 | 2022-03-31 | 2019-12-05 |         34 | High Available Datastore for Kubernetes                                                                                                                                                                                                              |
|  3 | [clusterpedia](https://github.com/clusterpedia-io/clusterpedia)          |   226 | 2022-03-31 | 2021-10-08 |         33 | The Encyclopedia of Kubernetes clusters                                                                                                                                                                                                              |
|  4 | [kubernetes-lts](https://github.com/klts-io/kubernetes-lts)              |   185 | 2022-03-31 | 2021-07-16 |         18 | Kubernetes LTS(long term support)                                                                                                                                                                                                                    |
|  5 | [piraeus-operator](https://github.com/piraeusdatastore/piraeus-operator) |   177 | 2022-03-31 | 2019-07-25 |         27 | The Piraeus Operator manages LINSTOR clusters in Kubernetes.                                                                                                                                                                                         |
|  6 | [local-storage](https://github.com/hwameistor/local-storage)             |   157 | 2022-03-31 | 2022-01-19 |         14 | Local Storage is one of HwameiStor components. It will provision the local LVM volume.                                                                                                                                                               |
|  7 | [public-image-mirror](https://github.com/DaoCloud/public-image-mirror)   |    57 | 2022-04-06 | 2021-09-09 |         24 | 很多镜像都在国外。比如 gcr 。国内下载很慢，需要加速。                                                                                                                                                                                                |
|  8 | [spiderpool](https://github.com/spidernet-io/spiderpool)                 |    29 | 2022-04-04 | 2022-03-07 |          7 | ipam for kubernetes  https://spidernet-io.github.io/spiderpool/                                                                                                                                                                                      |
|  9 | [local-disk-manager](https://github.com/hwameistor/local-disk-manager)   |    24 | 2022-04-01 | 2022-01-19 |          4 | Local Disk Manager is one of HwameiStor components. It will manage all the local disks of the HwameiStor nodes, including provision local Disk volume, and disk health management.                                                                   |
| 10 | [ckube](https://github.com/DaoCloud/ckube)                               |    13 | 2022-04-04 | 2022-03-17 |          0 | Kubernetes APIServer 高性能代理组件，代理 APIServer 的 List 请求，其它类型的请求会直接反向代理到原生 APIServer。 CKube 还额外支持了分页、搜索和索引等功能。 并且，CKube 100% 兼容原生 kubectl 和 kube client sdk，只需要简单的配置即可实现全局替换。 |



## Skipped repos
<!--END_SECTION:github_repos-->

# Build

docker build . -t github-repo-stats:v0.1.0

following https://github.com/yihong0618/github-readme-stats style.
