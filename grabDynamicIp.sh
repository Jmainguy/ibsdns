#!/bin/bash
CONNECTION=$1
if [[ $CONNECTION == "" ]]; then
    echo "Usage ./grabDynamicIp.sh user@remoteHost"
    echo "You must pass connection string for remote box with a static IP where ibsdns is installed"
    echo "Please setup ssh keys so you can add this to a cronjob"
    exit 1
fi
NEWIP=$(curl -s https://ip.jmainguy.com | awk '{print $4}')
OLDIP=$(cat ~/lastIP.txt)
if [[ $NEWIP != $OLDIP ]]; then
    echo $NEWIP > ~/lastIP.txt
    ssh $CONNECTION "/usr/sbin/ibsdns --value \"$NEWIP\""
else
    echo "No update needed"
fi
