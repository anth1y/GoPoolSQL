# gophyderm
A PostgreSQL load balancer / clustering tool written in Go.

Pronounced "Pak-Kee-Go"

Future capabilites: replication, failover... more to come stay tuned

This will have ssl connections by default.

Will have command line  flags

Notes and ideas will be kept in a google doc: https://docs.google.com/document/d/1Q0zh_GgvjsBKpsakTp-P44nwAvfYEe4XAfwBn6WtPiA/edit

## TODO: ##
 
###*** Learn how postgres does clustering,*** ###

*  Have the app detect the master
* 
*  Send the master "writes"
* 
*  Send the slaves "reads"
* 
*  How to do recovery through the filesystem.