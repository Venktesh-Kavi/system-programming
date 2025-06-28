# Network Checks

For triaging a call from outside network use:
* nslookup - Used to check only for dns resolution
* Ping - Used to check for dns resolution + send ICMP packets to the address
* Traceroute - same as ping shows the complete resolution path to the destination.
* Windows (tracert)

### NetStat

* Identify listening ports in local using netstat -a -t (all sockets which are tcp)

### Nc (NetCat)
* Check connection to remote port using netcat
* nc -zv b-3.intcentralizedkafka01.03aq81.c4.kafka.ap-south-1.amazonaws.com 9094
* Z scan mode, V verbose

