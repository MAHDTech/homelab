# Nutanix Bootstrap

## Table of Contents

- [Nutanix Bootstrap](#nutanix-bootstrap)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Steps](#steps)

## Overview

Initial setup instructions for a new Nutanix Community Edition cluster.

## Steps

1. Determine the IP addressing for the cluster.

2. Install Nutanix Community Edition on each host.

3. SSH to the AHV host.

```bash
ssh root@<AHV_IP>
```

4. From AHV, SSH to the CVM which has the internal IP `192.168.5.2`.

```bash
ssh nutanix@192.168.5.2
```

5. Form the new cluster

```bash
CVM_IPS="10.10.100.22,10.10.100.23,10.10.100.24"

cluster --svm_ips $CVM_IPS create

cluster status
```

6. Configure the cluster.

```bash
CLUSTER_NAME="nutanix-cluster"
CLUSTER_VIP="10.10.100.20"
CLUSTER_DNS="prism-element.saltlabs.cloud"
DNS_SERVER_IP="10.10.100.254"
NTP_SERVER_IP="10.10.100.254"

ncli cluster edit-params new-name=$CLUSTER_NAME

ncli cluster add-to-name-servers servers=$DNS_SERVER_IP

ncli cluster set-external-ip-address external-ip-address=$CLUSTER_VIP

ncli cluster add-to-ntp-servers servers=$NTP_SERVER_IP
```

7. Once the cluster is formed, check the Hard Drives are correctly identified.

```bash
ssh ntnx-cvm-1

ncli disk ls

# Capture all the SSD disk IDs that are incorrectly identified as HDD.
SSD_IDS=(
  "00062342-a7a1-23db-5a27-1cfd087f170a::44"
  "00062342-a7a1-23db-5a27-1cfd087f170a::49"
  "00062342-a7a1-23db-5a27-1cfd087f170a::56"
  "00062342-a7a1-23db-5a27-1cfd087f170a::57"
  "00062342-a7a1-23db-5a27-1cfd087f170a::63"
  "00062342-a7a1-23db-5a27-1cfd087f170a::65"
)

for DISK_ID in ${SSD_IDS[@]}; do
  ncli disk update id=${DISK_ID} tier-name=SSD-SATA
done
```

8. Change the name of the CVMs. (you need to change the name from a different CVM)

```bash
ssh nutanix@<cvm_ip>

# Run this on each CVM
# NOTE: Takes a while to complete...
sudo /usr/local/nutanix/cluster/bin/change_cvm_hostname NTNX-AHV-1-CVM
sudo /usr/local/nutanix/cluster/bin/change_cvm_hostname NTNX-AHV-2-CVM
sudo /usr/local/nutanix/cluster/bin/change_cvm_hostname NTNX-AHV-3-CVM
sudo /usr/local/nutanix/cluster/bin/change_cvm_hostname NTNX-AHV-4-CVM

# Run this from a different CVM to the one you are changing.
change_cvm_display_name --cvm_ip=10.10.100.21 --cvm_name=NTNX-AHV-1-CVM
change_cvm_display_name --cvm_ip=10.10.100.22 --cvm_name=NTNX-AHV-2-CVM
change_cvm_display_name --cvm_ip=10.10.100.23 --cvm_name=NTNX-AHV-3-CVM
change_cvm_display_name --cvm_ip=10.10.100.24 --cvm_name=NTNX-AHV-4-CVM
```

9. Change the name of the AHV hosts.

- Option 1 from Prism Central
  - _Infrastructure_
  - _Hardware_
  - _Hosts_
  - Select the Host
  - Select _Rename_ from the _Actions_ dropdown menu.
  - Enter the new name.
  - Click _Save_

- Option 2 from CVM

```bash
ssh nutanix@<cvm_ip>

change_ahv_hostname --host_ip=10.10.100.11 --host_name=NTNX-AHV-1
change_ahv_hostname --host_ip=10.10.100.12 --host_name=NTNX-AHV-2
change_ahv_hostname --host_ip=10.10.100.13 --host_name=NTNX-AHV-3
change_ahv_hostname --host_ip=10.10.100.14 --host_name=NTNX-AHV-4
```
