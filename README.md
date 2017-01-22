## A simple Go program to generate syslog messages for PaloAlto firewalls 
### Tested with PANOS 7.1.2: [Syslog Field Description](https://www.paloaltonetworks.com/documentation/61/pan-os/pan-os/reports-and-logging/syslog-field-descriptions)
##### It sends a X random syslog messages every Y seconds to a device of your designation.
##### Usage 
```
$syslog-generator -h'
  -count int
        Number of syslog messages to send (default 1)    
  -ip string
        Syslog server IP address (default "10.34.1.100")      
  -port string 
        Port (default "11666")       
  -protocol string  
        Protocol (default "tcp")        
  -sleep duration 
        Sleep time between syslog messages (default 1ns)        
  -v    Prints current version
```

URL field in threat logs is set to "G0s9J4jAU3" (per random.org) for easier identification of generated entries

## TODO:
1. Include Traffic logs
2. Provide option to generate random IP addresses and ports in syslog entries instead of "1.1.1.1","2.2.2.2" and "11111","22222"
3. Upload binaries for Linux and Windows
