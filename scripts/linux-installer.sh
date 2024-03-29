#!/bin/sh

FILE=LICENSE

if [ ! -f "$FILE" ]; then
    echo "Please run script from main directory"
    exit 1
fi


FILE=src/deploy/linux/src

if [ -f "$FILE" ]; then
    echo "$FILE exists."
    rm -Rf  tf-network-connector-linux
    rm tf-network-connector.deb
    mkdir -p tf-network-connector-linux/DEBIAN
    cp src/linux/control tf-network-connector-linux/DEBIAN
    chmod +x src/linux/postinst 
    cp src/linux/postinst tf-network-connector-linux/DEBIAN

    mkdir -p tf-network-connector-linux/usr/local/bin
    cp $FILE tf-network-connector-linux/usr/local/bin/tf-network-connector-linux
    mkdir -p tf-network-connector-linux/usr/share/icons/
    cp src/qml/icon.ico tf-network-connector-linux/usr/share/icons/tf.ico
    mkdir -p  tf-network-connector-linux/usr/share/tf
    cp src/linux/threefold-networkconnector.desktop tf-network-connector-linux/usr/share/tf
    dpkg-deb --build tf-network-connector-linux



else 
    echo "$FILE does not exist."
fi



