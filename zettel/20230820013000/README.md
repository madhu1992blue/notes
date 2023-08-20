# Networking Types suitability - Virtualbox

| Network Type      | VM -> Host | Host -> VM              | VM1 <-> VM2 (Same network) | VM -> Net/LAN | Net/LAN -> VM           |
|----------------   |------------|-------------------------|----------------------------|---------------|-------------------------|
| Internal Network  | No         | No                      | Yes                        | No            | No                      |
| Host Only         | Yes        | Yes                     | Yes                        | No            | No                      |
| NAT               | Yes        | Yes via Port forwarding | No                         | Yes           | Yes via Port forwarding |
| NAT Network       | Yes        | Yes via Port forwarding | Yes                        | Yes           | Yes via Port forwarding |
| Bridged Network   | Yes        | Yes                     | Yes                        | Yes           | Yes                     |

Reference: https://www.virtualbox.org/manual/ch06.html

Notes: 
* If Bridged network is selected for the VM, Guest VM acquires an IP from DHCP in the same network as the host machine.
* If NAT/NAT network is selected for the VM, Host can communicate to VM by enabling Port forwarding.  
  It consumes a port on Host machine. Since now a port on Host machine is opened, other nodes in Host network can connect to Guest VM via the forwarded port. 
  If you intend to prevent other nodes from accessing Guest VM via Host's port, set the Host IP address to 127.0.0.1 in Virtualbox forwarding rule. This allows Host -> VM and prevents Net/LAN -> VM.
  If you intend to allow other nodes to access Guest VM, then set the Host IP address to IP address of the host or 0.0.0.0 in Port forwarding settings.


Thoughts:
* Internal Network seems most secure but also more restrictive.
* Bridged Network - perceived as adding a new node in same network as the host. (Z20230820012800)
 
#virtualbox #network
