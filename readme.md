# topx

Main purpose of this tool to provide more detailed information about 
system and processes. 

## Usage 
````bash 
$ ./topx 

 User		PID	    Name	  VMemory		 %MEM	  #Threads	Swap		CreateTime
 -----		-----	-----		-----		------		-----	---	    ------------------
 jicofo		1551	java		22.3 GB		0.298510	321		0		2020-09-25 15:00:14 +0000 UTC
 jvb		1378	java		10.2 GB		0.813751	75		0		2020-09-25 15:00:14 +0000 UTC
 root		1432	dockerd		4.0 GB		0.033659	51		0		2020-09-25 15:00:14 +0000 UTC
 root		1415	containerd	2.8 GB		0.016415	36		0		2020-09-25 15:00:14 +0000 UTC
 root		28780	ovs-vswitc	2.5 GB		0.108111	34		0		2020-09-25 15:05:12 +0000 UTC
 ubuntu		21138	go		    2.3 GB		0.007618	21		0		2020-10-01 21:55:03 +0000 UTC
 root		1464	libvirtd	1.8 GB		0.011852	16		0		2020-09-25 15:00:14 +0000 UTC
 ubuntu		21250	main		1.6 GB		0.004526	13		0		2020-10-01 21:55:03 +0000 UTC
 ubuntu		2906	ruby		1.1 GB		0.199768	6		0		2020-09-25 15:01:52 +0000 UTC
 ubuntu		2912	ruby		1.1 GB		0.197987	6		0		2020-09-25 15:01:52 +0000 UTC
````
By default, it prints top 10 processes which are using memory heavily with more information about process. 

`-top=20 ` flag could be used for listing top 20 processes which are using memory heavily. 

