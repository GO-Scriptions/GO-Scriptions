This is how I created a local area network in GNS3:

1. Open GNS3 and create a new project.

2. Add a switch.

3. Add two virtual PC simulators. (VPCS)

4. Access the consoles of each VPCS.

5. set the ip of PC1 with "ip 10.0.0.1 255.255.255.0"

6. enter "save" to save

7. set the ip of PC2 with "ip 10.0.0.2 255.255.255.0"

8. enter "save" to save

9. Link the switch and two VPCS's together. 

10. The two VPCS's should be able to ping each other.

11. Now you can add a VirtualBox VM to the network. Open the Oracle VM VirtualBox Manager.

12. Choose the VM you want to add. Remove all network adapters. GNS3 will add its own. 

13. Start the VM. 

14. Enter "sudo nano /etc/netplan/01-netcfg.yam1"

15. Replace "dhcp4: yes" with "addresses: [10.0.0.3/24]"

16. Save and exit.

17. Enter "sudo netplan apply"

18. In GNS3, go to edit, then preferences, then VirtualBox VM's, and then click new.

19. Select the VM you want to add and then click finish. 

20. Click apply and okay. Now you should be able to add the VM and connect it to the network. 
