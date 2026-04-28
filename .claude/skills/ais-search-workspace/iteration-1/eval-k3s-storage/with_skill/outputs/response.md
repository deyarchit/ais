# k3s Bare Metal Persistent Storage: CSI Driver Comparison (2025)

> The default k3s `local-path` provisioner ties PVs to a specific node — unsuitable for multi-node HA. A distributed CSI driver is essential.

## Comparison: Longhorn vs OpenEBS vs Rook-Ceph

| | **Longhorn** | **OpenEBS** | **Rook-Ceph** |
|---|---|---|---|
| **Ease of deployment** | Easiest — GUI included, helm one-liner | Moderate — disable unused engines for home lab | Most complex — steep learning curve |
| **Resource overhead** | Low–medium. Reports of increased idle CPU in versions post-1.4 (~0.1–0.3 cores/node) | Varies by engine. Mayastor is heavy; Jiva/cStor lighter (~0.05–0.5 cores/node) | Highest (~0.5–1.5 cores/node). Heavy CRUSH map recalculation |
| **Performance** | Good for most workloads; can bottleneck on spinning disk with large volumes | Mayastor (NVMe-over-TCP): excellent. Jiva/cStor: moderate | Excellent at scale with SSDs; slower for high-IOPS DB workloads on HDDs |
| **HA / replication** | Block replication, auto-rebuild, snapshots + backups | cStor/Jiva replicated; Mayastor less so | Highly resilient erasure coding — needs 3–4 nodes for true HA |
| **Access modes** | RWO + RWX (block) | Primarily RWO; cStor offers RWX | RWO + RWX (CephFS) + object (S3-compat) |
| **Best for** | Beginners, general workloads, resource-constrained nodes | NVMe workloads or mixed storage profiles | Advanced labs, enterprise simulation, object storage needs |

## Recommendation for a 3-Node Home Lab

**Use Longhorn.** For a 3-node k3s bare metal cluster in 2025:

- Easiest install: `helm install longhorn longhorn/longhorn -n longhorn-system`
- Web UI for volume/snapshot/backup management out of the box
- Native k3s integration, well-documented
- Appropriate resource footprint for mini PCs or similar hardware
- Built-in S3-compatible backup to external storage

**Upgrade path:** If you add NVMe drives and need higher IOPS, migrate to OpenEBS with Mayastor. Only move to Rook-Ceph if you need S3-compatible object storage or CephFS RWX volumes cluster-wide.

## Quick Install (Longhorn on k3s)

```bash
# Prerequisites: open-iscsi on all nodes
sudo apt install open-iscsi

# Install via Helm
helm repo add longhorn https://charts.longhorn.io
helm repo update
helm install longhorn longhorn/longhorn \
  --namespace longhorn-system \
  --create-namespace \
  --set defaultSettings.defaultReplicaCount=2
```

Set `defaultReplicaCount=2` for a 3-node cluster (tolerates 1 node failure).

*(Sources retrieved via `ais` CLI — see sources.txt)*
