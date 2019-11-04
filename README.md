# ibsdns
InternetBS DNS Updater
## Why?
My buddy Keenan wanted to update home.example.com to point to his IP at home,
however he has a dynamic IP provided to home and so he needed a way to keep it current.
He was insistent on using InternetBS as compared to Route53 as he believed it would be less expensive (free basically)

So I made this for him, though I did it in a way where it can be used by anyone with a InternetBS api key,
and a static host to ssh to and run the update tool from (internetBS limits their IP to a static IP you provide ahead of time)

## HowTo
On server at home (or where the IP will be dynamicly changing) you need to add `grabDynamicIp.sh` to a cronjob, like once every 5 minutes perhaps

```/bin/bash
*/5 * * * * /home/jmainguy/grabDynamicIp.sh jmainguy@remotehost.com > /home/jmainguy/dnsLog.txt 2>&1
```

On remotehost add ibsdns binary to `/usr/sbin/ibsdns` and edit `/opt/ibsdns/config.yaml` for your values.

Ensure passwordless ssh is setup from home to remote host, and that the user can read /opt/ibsdns/config.yaml
