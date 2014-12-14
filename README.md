# ReadMe

## Intro
Go program for autologin and login refresh for Securepoint Network Access Controler (NAC). [Like this one.](http://www.securepoint.cc/products-wifi-network-access-controller.html)

![Securepoint](http://www.securepoint.cc/images/securepoint-logo.jpg)

Picture probably see on your login page.

I am not associated with Securepoint, but they have a github profile: https://github.com/Securepoint

## Behavior of Securepoint NAC
The NAC has a web login page, which is served by the DNS server of the NAC. The Login page url is normally https://controller.mobile.lan/ and the page communicates with the the php skript: https://controller.mobile.lan/portal_api.php. There are 3 instructions:

* authenticate
* disconnect
* refresh

In my case the refresh instruction was turned off. Since it is a german appliance you will bes disconnected at least every 24h. In Germany you get disconnected every 24h without any reason, it's a cultural thing.

## Usage

The login data is provided via a config file in json. Just rename login.conf.default to login.conf and fill in you credentials. The password digest can be taken from the log file after the first connect instruction, the server answers in json some configuration details including the digest. After that the programm can be run via cron or windows task scheduler (Haven't tried compiling for windows yet). During the disconnect/connect cycle open connections are not interrupted. 

## ToDO
* Parse password digest from first response and fill in conf file automatically

## Known NAC Bugs
* There is a problem with Notebooks, when your ethernet card goes into energy saving mode your login session is lost
* Sometimes the DNS Server of the NAC crashes and your are unable to log in. If you have a valid session open you can still open ip connections. Sadly the use of other DNS Server will be blocked by the NAC (DNS caching is possible, but should be set up in advance)

Tested with: go version go1.2.1 linux/386