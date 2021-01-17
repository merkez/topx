# topx

Main purpose of this tool to provide more detailed information about 
system and processes. 

## Usage 
````bash 
$ ./topx 

+-----------------+------+------------------+----------+---------------+------+-------------------------------+
|      USER       | PID  |       NAME       | VMEMORY  | MEMPERCENTAGE | SWAP |          CREATETIME           |
+-----------------+------+------------------+----------+---------------+------+-------------------------------+
| root            |  839 | snapd            | 756.5 MB |      4.188787 |    0 | 2021-01-17 14:59:35 +0000 UTC |
| root            | 1586 | amazon-ssm-agent | 666.0 MB |      2.626873 |    0 | 2021-01-17 14:59:41 +0000 UTC |
| root            |  636 | /usr/bin/python3 | 110.7 MB |     2.0852127 |    0 | 2021-01-17 14:59:28 +0000 UTC |
| root            |  264 | multipathd       | 286.9 MB |      1.794225 |    0 | 2021-01-17 14:59:20 +0000 UTC |
| root            |  443 | /usr/bin/python3 | 30.0 MB  |     1.7938257 |    0 | 2021-01-17 14:59:26 +0000 UTC |
| root            |  159 | systemd-journald | 52.9 MB  |     1.4513464 |    0 | 2021-01-17 14:59:19 +0000 UTC |
| root            |    1 | systemd          | 105.4 MB |     1.2825016 |    0 | 2021-01-17 14:59:17 +0000 UTC |
| systemd-resolve |  360 | systemd-resolved | 24.6 MB  |     1.2226276 |    0 | 2021-01-17 14:59:23 +0000 UTC |
| ubuntu          | 2024 | systemd          | 18.9 MB  |    0.96077853 |    0 | 2021-01-17 15:02:02 +0000 UTC |
| root            |  421 | accounts-daemon  | 246.8 MB |    0.92804736 |    0 | 2021-01-17 14:59:26 +0000 UTC |
+-----------------+------+------------------+----------+---------------+------+-------------------------------+

````
By default, it prints top 10 processes which are using memory heavily with more information about process. 

`-top=20 ` flag could be used for listing top 20 processes which are using memory heavily. 

