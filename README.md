# github-repos-stats

# watching repositories status

This branch is for CNCF repos in China.
- [Mutli-Clusters related projects](https://github.com/pacoxu/github-repos-stats/tree/multi-clusters)


<!--START_SECTION:github_repos-->
### The intersted repos
| ID |                                  REPO                                  | STARS | UPDATEDAT  | CREATEDAT  | FORKSCOUNT |                                                                                                                     DESCRIPTIONS                                                                                                                     |
|----|------------------------------------------------------------------------|-------|------------|------------|------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  1 | [Dragonfly](https://github.com/dragonflyoss/Dragonfly)                 |  6008 | 2022-12-30 | 2017-11-15 |        780 | Dragonfly is an intelligent P2P based image and file distribution system.                                                                                                                                                                            |
|  2 | [bfe](https://github.com/bfenetworks/bfe)                              |  5739 | 2022-12-30 | 2019-07-31 |        921 | A modern layer 7 load balancer from baidu                                                                                                                                                                                                            |
|  3 | [kubeedge](https://github.com/kubeedge/kubeedge)                       |  5533 | 2022-12-30 | 2018-09-28 |       1464 | Kubernetes Native Edge Computing Framework (project under CNCF)                                                                                                                                                                                      |
|  4 | [kubevela](https://github.com/kubevela/kubevela)                       |  4540 | 2022-12-30 | 2020-07-03 |        642 | The Modern Application Platform.                                                                                                                                                                                                                     |
|  5 | [rainbond](https://github.com/goodrain/rainbond)                       |  3706 | 2022-12-30 | 2017-11-05 |        674 | Cloud native multi cloud application management platform that make application management and delivery easier                                                                                                                                        |
|  6 | [volcano](https://github.com/volcano-sh/volcano)                       |  2760 | 2022-12-29 | 2019-03-14 |        644 | A Cloud Native Batch System (Project under CNCF)                                                                                                                                                                                                     |
|  7 | [radon](https://github.com/radondb/radon)                              |  1707 | 2022-12-29 | 2018-01-18 |        217 | RadonDB is an open source, cloud-native MySQL database for building global, scalable cloud services                                                                                                                                                  |
|  8 | [kyuubi](https://github.com/apache/kyuubi)                             |  1383 | 2022-12-30 | 2017-12-18 |        510 | Apache Kyuubi is a distributed and multi-tenant gateway to provide serverless SQL on data warehouses and lakehouses.                                                                                                                                 |
|  9 | [community-edition](https://github.com/vmware-tanzu/community-edition) |  1347 | 2022-12-17 | 2020-10-13 |        298 | VMware Tanzu Community Edition is no longer an actively maintained project. Code is available for historical purposes only.                                                                                                                          |
| 10 | [openelb](https://github.com/openelb/openelb)                          |  1211 | 2022-12-30 | 2019-02-01 |        159 | Load Balancer Implementation for Kubernetes in Bare-Metal, Edge, and Virtualization                                                                                                                                                                  |
| 11 | [clusterpedia](https://github.com/clusterpedia-io/clusterpedia)        |   537 | 2022-12-27 | 2021-10-08 |         82 | The Encyclopedia of Kubernetes clusters                                                                                                                                                                                                              |
| 12 | [merbridge](https://github.com/merbridge/merbridge)                    |   529 | 2022-12-29 | 2022-01-12 |         63 | Use eBPF to speed up your Service Mesh like crossing an Einstein-Rosen Bridge.                                                                                                                                                                       |
| 13 | [fabedge](https://github.com/FabEdge/fabedge)                          |   464 | 2022-12-12 | 2021-07-16 |         47 | Secure Edge Networking Solution Based On Kubernetes                                                                                                                                                                                                  |
| 14 | [spiderpool](https://github.com/spidernet-io/spiderpool)               |   269 | 2022-12-29 | 2022-03-07 |         35 | kubernetes ipam                                                                                                                                                                                                                                      |
| 15 | [ferry](https://github.com/ferryproxy/ferry)                           |    74 | 2022-12-16 | 2021-10-18 |          7 | Ferry is a Kubernetes multi-cluster communication component that eliminates communication differences between clusters as if they were in a single cluster, regardless of the network environment those clusters are in.                             |
| 16 | [ckube](https://github.com/DaoCloud/ckube)                             |    17 | 2022-09-19 | 2022-03-17 |          6 | Kubernetes APIServer 高性能代理组件，代理 APIServer 的 List 请求，其它类型的请求会直接反向代理到原生 APIServer。 CKube 还额外支持了分页、搜索和索引等功能。 并且，CKube 100% 兼容原生 kubectl 和 kube client sdk，只需要简单的配置即可实现全局替换。 |



#### Skipped repos
hwameistor/local-disk-manager
chaos-mesh/chaos-mesh
kubeovn/kube-ovn
vmware-tanzu/octant
vmware-tanzu/velero
wzshiming/fake-k8s
karmada-io/karmada
kubesphere/kubesphere
superedge/superedge
carina-io/carina
openyurtio/openyurt
klts-io/kubernetes-lts
piraeusdatastore/piraeus-operator
OpenFunction/OpenFunction
clusternet/clusternet
kubecube-io/KubeCube
openpitrix/openpitrix
wzshiming/fake-kubelet
KubeOperator/KubeOperator
cni-genie/CNI-Genie
vmware-tanzu/sonobuoy
piraeusdatastore/piraeus
dragonflyoss/Dragonfly2
kubediag/kubediag
fanux/sealos
hwameistor/local-storage<!--END_SECTION:github_repos-->

# Build

docker build . -t github-repo-stats:v0.1.0

following https://github.com/yihong0618/github-readme-stats style.
