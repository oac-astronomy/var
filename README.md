# var by Victor Gabriel Galvan
NBS Assignment Senior Network Specialist 2021


Hi everyone,

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
go build  var.go   
ls
Desktop  Documents  Downloads  Music  net.go  Pictures  Public  Templates  test  test.go  testing.git  var  var.go  Videos

#./var
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
LESS_TERMCAP_mb =>
LESS_TERMCAP_md =>
LESS_TERMCAP_me =>
LESS_TERMCAP_so =>
LESS_TERMCAP_se =>
LESS_TERMCAP_us =>
LESS_TERMCAP_ue =>
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

If you want to deploy Golang over the Ubuntu o Kali Linux, please follow the next command. 

sudo apt-get install -y golang

sudo nano ~/.bashrc
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

source .bashrc

At this point, you are able to run Go programs in your own environment.

 

 



  
 



