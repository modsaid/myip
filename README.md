# myip

It is common to need to know what is the ip of a given system as perceived by the outside world
whether it is for network debugging, vpn checking or other reasons

This simple configuration helps with that

## servier side

Any existing server can tell what is the externl ip address if you made a request to it

many servers out there provide that information with more detailed information about country, region and so
however, the simplified nginx configuration under `nginx` can service just the ip address plainly

this is available through http://ip.modsaid.com

feel free to use this from any server or anywhere, if it helps

```shell
curl ip.modsaid.com
```

## command wrapper

the go binary is just a wrapper to make it more intuitive and easy, of course you can extend it
to add any further info or to use another external ip provider (server side)

### Installation

```shell
cd cmd 
go build
sudo cp myip /usr/local/bin/
```


### Usage

```shell
myip
```
