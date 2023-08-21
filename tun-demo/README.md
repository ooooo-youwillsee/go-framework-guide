```shell
# help command
ip tuntap help 
ip addr help
ip link help

# create a tun device
sudo ip tuntap add dev tun0 mode tun

# add an addr for tun0
sudo ip addr add 1.1.1.1/24 dev tun0

# up tun0
sudo ip link set dev tun0 up

# ping tun mode
ping -c 1.1.1.1

# execute go program
go run main.go
```