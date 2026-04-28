# Persistent Storage Solutions for k3s Home Lab Clusters (2025)

## Overview

For a bare-metal 3-node k3s home lab, choosing the right CSI (Container Storage Interface) driver is critical. The primary options that have gained traction in the k3s/lightweight Kubernetes ecosystem are:

- **Longhorn** (SUSE/Rancher)
- **OpenEBS** (CNCF, DataStax/MayaData)
- **Rook-Ceph** (CNCF)

Below is a detailed comparison followed by a recommendation.

---

## Background: k3s Storage Context

k3s ships with a built-in local-path-provisioner (LocalPath) that provides node-local storage. It works well for single-node or stateless workloads but offers no replication, no snapshots, and no cross-node access. For a proper 3-node home lab running stateful workloads (databases, monitoring stacks, media servers, etc.), a distributed storage solution is necessary.

---

## Persistent Storage Solutions Compared

### 1. Longhorn

**What it is:** A cloud-native distributed block storage system originally developed by Rancher (now SUSE). It is a CNCF Graduated project as of 2024.

**Maturity:**
- CNCF Graduated project (highest CNCF maturity tier)
- Actively maintained with frequent releases
- First-class citizen in the Rancher/k3s ecosystem — officially supported and recommended by SUSE/Rancher for k3s deployments
- Stable API and production-ready for small-to-medium workloads

**Architecture:**
- Runs as Kubernetes workloads (DaemonSets, Deployments)
- Each volume gets a dedicated Longhorn engine process and replicas distributed across nodes
- Uses iSCSI/NVMe-oF internally
- Provides: block storage (RWO), ReadWriteMany (RWX via NFS sharing layer), snapshots, backups to S3/NFS, volume live migration

**Resource Overhead:**
- Moderate. Each volume spawns lightweight engine + replica processes
- Minimum recommended: 2 CPUs and 4 GB RAM per node for comfortable operation
- Disk overhead: replica count (default 3) means 3x disk usage across nodes
- Memory footprint on idle: approximately 300–600 MB across the cluster for the control plane components

**Ease of Setup:**
- Very easy. Can be installed via:
  - `kubectl apply` with a single manifest
  - Helm chart (one-liner)
  - Rancher UI if using Rancher
- Prerequisites: `open-iscsi`, `nfs-common`, `bash`, `curl`, `findmnt`, `grep`, `awk` on host nodes — installable via a single prerequisite script
- Web UI included out of the box — accessible via browser for volume management, backup scheduling, and health monitoring
- k3s-specific documentation is thorough and officially provided

**Limitations:**
- Not suited for very high-throughput, latency-sensitive workloads (not a replacement for enterprise SAN)
- RWX support is layered (uses NFS share on top), slightly more complex than native RWX
- Significant per-volume process overhead at very high volume counts (100s of volumes)

---

### 2. OpenEBS

**What it is:** A CNCF Incubating project providing multiple storage engines under one umbrella. In 2024–2025, OpenEBS has converged around two primary engines: **Mayastor** (now called "Replicated PV Mayastor") and **LocalPV** variants.

**Maturity:**
- CNCF Incubating project
- Long history (originally MayaData, now community-driven)
- The older engines (cStor, Jiva) are in maintenance/deprecated mode
- Mayastor is the strategic engine going forward — it reached general availability and is production-ready
- LocalPV (hostpath, ZFS, LVM) variants are very stable and widely used

**Architecture (key engines):**

- **LocalPV Hostpath/LVM/ZFS**: Node-local storage with no replication. Simple and fast.
- **Mayastor (Replicated PV)**: High-performance replicated block storage built on NVMe-oF and SPDK (Storage Performance Development Kit). Runs in user-space for high throughput and low latency.

**Resource Overhead:**
- LocalPV: Extremely lightweight (near zero overhead — just a provisioner pod)
- Mayastor: Heavier than Longhorn — requires dedicated CPU cores (hugepages), significant RAM (1–2 GB per storage node for the IO engine), and NVMe or fast SSD backing. Requires kernel 5.15+ and specific kernel modules (io_uring, nvme_tcp)
- For a home lab with modest hardware, Mayastor can be overkill and demanding

**Ease of Setup:**
- LocalPV: Very easy, minimal prerequisites
- Mayastor: Moderately complex — requires configuring hugepages, loading specific kernel modules, and proper node labeling before installation. Helm-based installation, but not as streamlined as Longhorn for k3s
- No built-in web UI (relies on kubectl and Prometheus/Grafana for observability)
- k3s-specific documentation exists but is less polished than Longhorn's

**Limitations:**
- Fragmented ecosystem (multiple engines with different capabilities and maturity levels)
- Mayastor's hardware requirements (hugepages, fast NVMe) may not match typical home lab hardware
- Less "batteries included" experience compared to Longhorn
- Snapshots and backup support less mature than Longhorn for the replicated engine

---

### 3. Rook-Ceph

**What it is:** Rook is a CNCF Graduated storage orchestrator that deploys and manages Ceph — a battle-tested, enterprise-grade distributed storage system — inside Kubernetes. Ceph itself has been in production at large scale for over 15 years.

**Maturity:**
- Rook: CNCF Graduated
- Ceph: Extremely mature, production-proven at petabyte scale (used by major cloud providers and enterprises)
- The combination is very stable; Ceph Reef (18.x) and Squid (19.x) releases are current in 2025
- Feature-complete: supports block (RBD/RWO), filesystem (CephFS/RWX), and object storage (S3-compatible/RGW) from a single cluster

**Architecture:**
- Ceph requires dedicated OSD (Object Storage Daemon) processes per disk
- MON (Monitor) quorum nodes (minimum 3 for production)
- MGR, MDS (for CephFS), and optionally RGW (for object storage)
- Rook manages all of this as Kubernetes CRDs and operators

**Resource Overhead:**
- High. This is the main drawback for home labs.
- Minimum practical requirements: 3 nodes, each with at least 4 CPUs and 8–16 GB RAM, dedicated raw disks (not partitions by default)
- OSD processes alone consume 1–2 GB RAM each at rest
- MON processes: ~500 MB each × 3 = 1.5 GB just for monitors
- Full Ceph cluster on 3 nodes realistically consumes 6–12 GB RAM across the cluster for storage infrastructure alone
- Not suitable for low-resource home lab nodes (e.g., nodes with 4–8 GB RAM shared with workloads)

**Ease of Setup:**
- Complex. Rook-Ceph has one of the steeper learning curves in the Kubernetes storage ecosystem.
- Requires careful planning: dedicated block devices, network configuration, OSD device paths
- YAML configurations are verbose and require understanding of Ceph concepts (pools, placement groups, CRUSH maps)
- Helm chart available but initial setup still requires significant configuration
- No built-in simple web dashboard (Ceph Dashboard exists but requires additional setup)
- Troubleshooting Ceph issues requires Ceph-specific knowledge (ceph status, ceph osd tree, etc.)

**Strengths for home lab:**
- Once running, it is extremely capable and feature-rich
- Object storage (S3-compatible) is a major advantage if needed
- CephFS provides native RWX without the NFS intermediary that Longhorn uses

**Limitations:**
- Resource-heavy: often not practical on nodes with less than 8 GB RAM
- Operational complexity is high — recovery from node failures requires Ceph knowledge
- Overkill for most home lab use cases

---

## Summary Comparison Table

| Dimension               | Longhorn                        | OpenEBS (Mayastor)              | Rook-Ceph                         |
|-------------------------|----------------------------------|---------------------------------|-----------------------------------|
| CNCF Status             | Graduated                        | Incubating                      | Graduated (Rook) / Mature (Ceph)  |
| Maturity                | High (k3s ecosystem native)      | Medium-High (engine-dependent)  | Very High (enterprise-grade)      |
| Min RAM per node        | ~2–4 GB                          | ~4–8 GB (Mayastor)              | ~8–16 GB                          |
| Ease of Setup           | Easy (excellent k3s docs)        | Moderate                        | Hard                              |
| Block Storage (RWO)     | Yes                              | Yes                             | Yes (RBD)                         |
| Shared Storage (RWX)    | Yes (via NFS layer)              | No (Mayastor is RWO only)       | Yes (CephFS, native)              |
| Object Storage          | No (external backup only)        | No                              | Yes (S3-compatible)               |
| Snapshots/Backups       | Yes (built-in, S3/NFS target)    | Limited                         | Yes (RBD snapshots)               |
| Web UI                  | Yes (built-in)                   | No                              | Partial (Ceph Dashboard)          |
| k3s Documentation       | Excellent (first-party)          | Good                            | Community docs only               |
| Recommended for 3-node  | Yes                              | Conditional (depends on HW)     | Not ideal (resource constraints)  |

---

## Recommendation: Longhorn

For a 3-node bare-metal k3s home lab in 2025, **Longhorn is the recommended choice**.

### Reasons:

1. **Designed for k3s**: Longhorn is developed by the same organization (SUSE/Rancher) that builds k3s. It has first-class integration, tested compatibility, and documentation written specifically for k3s deployments. There are no compatibility surprises.

2. **Right-sized resource overhead**: On typical home lab hardware (nodes with 4–16 GB RAM), Longhorn's footprint is manageable. Rook-Ceph would consume the majority of available RAM leaving little for actual workloads. OpenEBS Mayastor has similar or higher overhead than Longhorn while offering less on home lab hardware.

3. **Ease of operation**: The built-in Longhorn UI makes it easy to monitor volume health, manage replicas, and schedule backups without needing deep storage expertise. For a home lab (where you are likely operating alone without a dedicated storage team), this observability matters greatly.

4. **Adequate feature set**: For home lab workloads (databases, monitoring stacks, media servers, home automation), Longhorn's RWO block storage and RWX (NFS-backed) cover essentially all use cases. Built-in snapshot and backup to S3/MinIO is a practical bonus.

5. **CNCF Graduated**: Longhorn has reached CNCF graduation, signaling production maturity, strong governance, and long-term viability.

6. **Resilience on 3 nodes**: With 3 replicas across 3 nodes, Longhorn can tolerate a single node failure while maintaining data availability — appropriate for a 3-node cluster.

### When to choose alternatives:

- **OpenEBS LocalPV** (not Mayastor): If you need maximum simplicity and your workloads are tolerant of node-local (non-replicated) storage, OpenEBS LocalPV with ZFS or LVM backend is an excellent lightweight option. It pairs well with application-level replication (e.g., databases with built-in replication).

- **Rook-Ceph**: If your nodes have 32+ GB RAM each, you need native S3-compatible object storage, and you want enterprise-grade Ceph for learning purposes or future scale-out. Not recommended for constrained home lab hardware.

- **OpenEBS Mayastor**: If your nodes have NVMe SSDs and you need maximum storage throughput (e.g., running high-performance databases), and you are comfortable with the additional setup complexity.

---

## Practical Setup Notes for Longhorn on k3s

1. Install prerequisites on all nodes:
   ```bash
   apt install open-iscsi nfs-common  # Debian/Ubuntu
   systemctl enable --now iscsid
   ```

2. Install Longhorn via Helm:
   ```bash
   helm repo add longhorn https://charts.longhorn.io
   helm repo update
   helm install longhorn longhorn/longhorn \
     --namespace longhorn-system \
     --create-namespace \
     --version 1.7.x
   ```

3. Access the UI via port-forward or ingress:
   ```bash
   kubectl port-forward svc/longhorn-frontend 8080:80 -n longhorn-system
   ```

4. Set Longhorn as the default StorageClass if desired:
   ```bash
   kubectl patch storageclass local-path \
     -p '{"metadata": {"annotations": {"storageclass.kubernetes.io/is-default-class": "false"}}}'
   ```

---

## Conclusion

For a 3-node bare-metal k3s home lab in 2025, Longhorn strikes the best balance of maturity, ease of use, resource efficiency, and feature completeness. It is the pragmatic choice that lets you focus on running workloads rather than managing storage infrastructure. Rook-Ceph, while more powerful, demands significantly more resources and operational knowledge than is justified for most home lab scenarios.
