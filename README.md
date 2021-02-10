# Lists Environment Variables (Variables.go)  
# Lists information about network interfaces (Network.go)
# by Victor Gabriel Galvan - info@itborder.ca

NBS Assignment Senior Network Specialist 2021

# Lists Environment Variables (Variables.go)

Hi Team,

There is my NBS Assignment so I would like to explain what I did.

Firstly, I created a program that lists environment variables the name on the pool is var and it is written in GO language. 
I tested it on my Kaly Desktop machine and it is working. 

Secondly, I moved the file var.go to my GitHub path to commit with GitHub Desktop. After checking all details about the path and setup, I commit to main. Besides, the GitHub repository on the web was updated. 

Finally, I wrote a brief explanation of how I did this process.

Technical Info

Compiling your GO Code.

It can be compiled on 
or 
In Linux environment. Remember to use the right privileges such as Root user or Sudo command.
In my case, I used Root privileged.

# go build  var.go   

# ls
Desktop  Documents  Downloads  Music  net.go  Pictures  Public  Templates  test  test.go  testing.git  var  var.go  Videos

# ./var

PATH => /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

SHELL => /usr/bin/zsh

TERM => xterm

LANG => en_US.UTF-8

LS_COLORS => rs

MAIL => /var/mail/root

LOGNAME => root

USER => root

HOME => /root

SUDO_COMMAND => /usr/bin/zsh

SUDO_USER => kali

SUDO_UID => 1000

SUDO_GID => 1000

SHLVL => 1

PWD => /home/kali

OLDPWD => /home/kali/testing.git

_ => /home/kali/./var

CUSTOM => 500

CUSTOM=> 500

GOROOT=>

The output must be changed because of the system. 

The same executable code on play.golang.com show me:

PATH => /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

HOSTNAME => 8ca34333c129

CUSTOM => 500

CUSTOM=> 500

GOROOT=> 

If you want to deploy Golang over the Ubuntu o Kali Linux, please follow the next commands. 

# sudo apt-get install -y golang

# sudo nano ~/.bashrc

# export GOROOT=/usr/local/go

# export GOPATH=$HOME/go

# export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# source .bashrc

At this point, you are able to run go programs in your own environment.

In addition, you can use git command on Linux. It is included by default. 


# Lists information about network interfaces (Network.go)

My IP address

# ifconfig

	eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500

    inet 192.168.2.216  netmask 255.255.255.0  broadcast 192.168.2.255

    inet6 fe80::20c:29ff:fe3d:e7e0  prefixlen 64  scopeid 0x20<link>

    inet6 fd74:c308:a7ae:0:c2ad:db5c:f52e:c9e8  prefixlen 64  scopeid 0x0<global>

    inet6 fd74:c308:a7ae:0:20c:29ff:fe3d:e7e0  prefixlen 64  scopeid 0x0<global>

    inet6 fd74:c308:a7ae::820  prefixlen 128  scopeid 0x0<global>

    ether 00:0c:29:3d:e7:e0  txqueuelen 1000  (Ethernet)

    RX packets 50913  bytes 66871505 (63.7 MiB)

    RX errors 0  dropped 0  overruns 0  frame 0

    TX packets 11498  bytes 1363097 (1.2 MiB)

    TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

	lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536

    inet 127.0.0.1  netmask 255.0.0.0

    inet6 ::1  prefixlen 128  scopeid 0x10<host>

    loop  txqueuelen 1000  (Local Loopback)

    RX packets 18  bytes 834 (834.0 B)

    RX errors 0  dropped 0  overruns 0  frame 0

    TX packets 18  bytes 834 (834.0 B)

    TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

# ip a

	1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000

    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00

    inet 127.0.0.1/8 scope host lo

	valid_lft forever preferred_lft forever

    inet6 ::1/128 scope host

	valid_lft forever preferred_lft forever

	2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    
	link/ether 00:0c:29:3d:e7:e0 brd ff:ff:ff:ff:ff:ff

    inet 192.168.2.216/24 brd 192.168.2.255 scope global dynamic noprefixroute eth0

	valid_lft 33760sec preferred_lft 33760sec

    inet6 fd74:c308:a7ae::820/128 scope global noprefixroute

	valid_lft forever preferred_lft forever

    inet6 fd74:c308:a7ae:0:c2ad:db5c:f52e:c9e8/64 scope global temporary dynamic

	valid_lft 595359sec preferred_lft 76577sec

    inet6 fd74:c308:a7ae:0:20c:29ff:fe3d:e7e0/64 scope global mngtmpaddr noprefixroute

	valid_lft forever preferred_lft forever

    inet6 fe80::20c:29ff:fe3d:e7e0/64 scope link noprefixroute

	valid_lft forever preferred_lft forever

┌──(root💀kali)-[/home/kali]

└─# ./network

# 00:0c:29:3d:e7:e0

2021/02/10 12:49:30 lo 127.0.0.1/8

2021/02/10 12:49:30 lo ::1/128

# 2021/02/10 12:49:30 eth0 192.168.2.216/24

2021/02/10 12:49:30 eth0 fd74:c308:a7ae::820/128

2021/02/10 12:49:30 eth0 fd74:c308:a7ae:0:c2ad:db5c:f52e:c9e8/64

2021/02/10 12:49:30 eth0 fd74:c308:a7ae:0:20c:29ff:fe3d:e7e0/64

2021/02/10 12:49:30 eth0 fe80::20c:29ff:fe3d:e7e0/64

My ip is 192.168.2.216 the mask is 255.255.255.0 or /24 and the Mac address is 00:0c:29:3d:e7:e0

The code is not elegant but it is working. 
		
		
		
		
 

 

