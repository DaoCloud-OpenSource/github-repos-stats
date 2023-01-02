# github-repos-stats

# watching repositories status

This branch is for CNCF repos in China.
- [Mutli-Clusters related projects](https://github.com/pacoxu/github-repos-stats/tree/multi-clusters)


<!--START_SECTION:github_repos-->
### The intersted repos
| ID |                                   REPO                                   | STARS | UPDATEDAT  | CREATEDAT  | FORKSCOUNT |                                                                                                    DESCRIPTIONS                                                                                                    |
|----|--------------------------------------------------------------------------|-------|------------|------------|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  1 | [octant](https://github.com/vmware-tanzu/octant)                         |  6178 | 2022-12-30 | 2019-06-19 |        438 | Highly extensible platform for developers to better understand the complexity of Kubernetes clusters.                                                                                                              |
|  2 | [Dragonfly](https://github.com/dragonflyoss/Dragonfly)                   |  6011 | 2023-01-02 | 2017-11-15 |        780 | Dragonfly is an intelligent P2P based image and file distribution system.                                                                                                                                          |
|  3 | [kubeedge](https://github.com/kubeedge/kubeedge)                         |  5536 | 2023-01-01 | 2018-09-28 |       1464 | Kubernetes Native Edge Computing Framework (project under CNCF)                                                                                                                                                    |
|  4 | [chaos-mesh](https://github.com/chaos-mesh/chaos-mesh)                   |  5413 | 2023-01-02 | 2019-09-04 |        678 | A Chaos Engineering Platform for Kubernetes.                                                                                                                                                                       |
|  5 | [kubevela](https://github.com/kubevela/kubevela)                         |  4546 | 2023-01-01 | 2020-07-03 |        642 | The Modern Application Platform.                                                                                                                                                                                   |
|  6 | [KubeOperator](https://github.com/KubeOperator/KubeOperator)             |  4532 | 2022-12-30 | 2018-11-19 |        824 | KubeOperator 是一个开源的轻量级 Kubernetes 发行版，专注于帮助企业规划、部署和运营生产级别的 K8s 集群。                                                                                                             |
|  7 | [sonobuoy](https://github.com/vmware-tanzu/sonobuoy)                     |  2662 | 2023-01-01 | 2017-07-26 |        324 | Sonobuoy is a diagnostic tool that makes it easier to understand the state of a Kubernetes cluster by running a set of Kubernetes conformance tests and other plugins in an accessible and non-destructive manner. |
|  8 | [kyuubi](https://github.com/apache/kyuubi)                               |  1396 | 2023-01-01 | 2017-12-18 |        510 | Apache Kyuubi is a distributed and multi-tenant gateway to provide serverless SQL on data warehouses and lakehouses.                                                                                               |
|  9 | [community-edition](https://github.com/vmware-tanzu/community-edition)   |  1347 | 2022-12-17 | 2020-10-13 |        299 | VMware Tanzu Community Edition is no longer an actively maintained project. Code is available for historical purposes only.                                                                                        |
| 10 | [clusternet](https://github.com/clusternet/clusternet)                   |  1078 | 2023-01-02 | 2021-06-07 |        243 | Managing your Kubernetes clusters (including public, private, edge, etc) as easily as visiting the Internet ⎈                                                                                                      |
| 11 | [Dragonfly2](https://github.com/dragonflyoss/Dragonfly2)                 |   941 | 2023-01-02 | 2020-11-04 |        135 | Dragonfly is an intelligent P2P based image and file distribution system, it also provides a variety of enterprise-level (efficiency, stability, safety, low-cost) product features.                               |
| 12 | [superedge](https://github.com/superedge/superedge)                      |   882 | 2023-01-01 | 2020-12-19 |        191 | An edge-native container management system for edge computing                                                                                                                                                      |
| 13 | [openpitrix](https://github.com/openpitrix/openpitrix)                   |   799 | 2023-01-01 | 2017-10-11 |        163 | Application Management Platform on Multi-Cloud Environment                                                                                                                                                         |
| 14 | [clusterpedia](https://github.com/clusterpedia-io/clusterpedia)          |   537 | 2022-12-27 | 2021-10-08 |         82 | The Encyclopedia of Kubernetes clusters                                                                                                                                                                            |
| 15 | [piraeus-operator](https://github.com/piraeusdatastore/piraeus-operator) |   239 | 2022-12-26 | 2019-07-25 |         37 | The Piraeus Operator manages LINSTOR clusters in Kubernetes.                                                                                                                                                       |
| 16 | [fake-kubelet](https://github.com/wzshiming/fake-kubelet)                |    59 | 2022-11-05 | 2020-07-20 |          7 | [Move to https://github.com/kubernetes-sigs/kwok] This is a fake kubelet. that can simulate any number of nodes and maintain pods on those nodes. It is useful for test control plane.                             |



#### Skipped repos
hwameistor/local-disk-manager
hwameistor/local-storage
karmada-io/karmada
openelb/openelb
wzshiming/fake-k8s
kubecube-io/KubeCube
cni-genie/CNI-Genie
kubediag/kubediag
fanux/sealos
carina-io/carina
merbridge/merbridge
klts-io/kubernetes-lts
kubesphere/kubesphere
goodrain/rainbond
piraeusdatastore/piraeus
volcano-sh/volcano
kubeovn/kube-ovn
bfenetworks/bfe
DaoCloud/ckube
ferry-proxy/ferry
OpenFunction/OpenFunction
FabEdge/fabedge
openyurtio/openyurt
radondb/radon
vmware-tanzu/velero
spidernet-io/spiderpool<!--END_SECTION:github_repos-->

# Build

docker build . -t github-repo-stats:v0.1.0

following https://github.com/yihong0618/github-readme-stats style.
