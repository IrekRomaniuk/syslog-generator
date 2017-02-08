## A simple Go program to generate syslog messages
### [PANOS Syslog Field Description](https://www.paloaltonetworks.com/documentation/61/pan-os/pan-os/reports-and-logging/syslog-field-descriptions)
##### It sends a X random syslog messages every Y seconds to a syslog server of your designation.
##### Usage 
```
$syslog-generator -h'
 -freq uint
       Frequency of syslog messages/sec (default 2)
 -ip string
       Syslog server IP address (default "10.34.1.100")
 -port string
       Port (default "11666")
 -protocol string
       Protocol (default "tcp")
 -v    Prints current version
  
  ```
##### Setup
  ```
  $ go get github.com/IrekRomaniuk/syslog-generator
  $ cd $GOPATH/bin
  ```
 ##### Example (from prev version, no count/sleep replaced by freq) 
  ```
  $ ./syslog-generator -ip="10.73.21.205" -port="514" -protocol="udp"
  $ sudo tail -f /var/log/messages
  Jan 22 10:10:14 January 22 10:10:14 dc01ap-p001mon 1,2017/22/01 10:10:14,001901000999,THREAT,file,1,2017/22/01 10:10:15,1.1.1.1,2.2.2.2,0.0.0.0,0.0.0.0,RULE fake,me,you,App test,vsys1,src,dst,ae1.100,ae2.200,LF-elk,2017/22/01 10:10:16,33891243,1,11111,22222,0,0,0x0,tcp,test,G0s9J4jAU3,This is test only,any,low,server-to-client,5210010,0x0,10.10.10.0-10.255.255.255,10.20.20.20-10.255.255.255,0,,,,,,,,,,,,,
  $ ./syslog-generator -ip="10.73.21.205" -port="514" -protocol="udp" -count=3 -sleep=1
  $ sudo tail -f /var/log/messages | grep G0s9J4jAU3
  Jan 22 10:24:07 January 22 10:24:07 dc01ap-p001mon 1,2017/22/01 10:24:07,001901000999,THREAT,file,1,2017/22/01 10:24:08,1.1.1.1,2.2.2.2,0.0.0.0,0.0.0.0,RULE fake,me,you,App test,vsys1,src,dst,ae1.100,ae2.200,LF-elk,2017/22/01 10:24:09,33891243,1,11111,22222,0,0,0x0,tcp,test,G0s9J4jAU3,This is test only,any,low,server-to-client,5210010,0x0,10.10.10.0-10.255.255.255,10.20.20.20-10.255.255.255,0,,,,,,,,,,,,,
  Jan 22 10:24:07 January 22 10:24:07 dc01ap-p001mon 1,2017/22/01 10:24:07,001901000999,THREAT,file,1,2017/22/01 10:24:08,1.1.1.1,2.2.2.2,0.0.0.0,0.0.0.0,RULE fake,me,you,App test,vsys1,src,dst,ae1.100,ae2.200,LF-elk,2017/22/01 10:24:09,33891243,1,11111,22222,0,0,0x0,tcp,test,G0s9J4jAU3,This is test only,any,low,server-to-client,5210010,0x0,10.10.10.0-10.255.255.255,10.20.20.20-10.255.255.255,0,,,,,,,,,,,,,
  Jan 22 10:24:07 January 22 10:24:07 dc01ap-p001mon 1,2017/22/01 10:24:07,001901000999,THREAT,file,1,2017/22/01 10:24:08,1.1.1.1,2.2.2.2,0.0.0.0,0.0.0.0,RULE fake,me,you,App test,vsys1,src,dst,ae1.100,ae2.200,LF-elk,2017/22/01 10:24:09,33891243,1,11111,22222,0,0,0x0,tcp,test,G0s9J4jAU3,This is test only,any,low,server-to-client,5210010,0x0,10.10.10.0-10.255.255.255,10.20.20.20-10.255.255.255,0,,,,,,,,,,,,,
```

Rule field is set to "G0s9J4jAU3" (per random.org) for easier identification of generated entries

## TODO:
1. Include Traffic logs
2. Provide option to generate random IP addresses and ports in syslog entries instead of "1.1.1.1","2.2.2.2" and "11111","22222"
3. Upload binaries for Linux and Windows
