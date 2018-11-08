#!/bin/bash

DIR="$(cd -P "$(dirname "${BASH_SOURCE[0]}")" > /dev/null && pwd)"
if [ ! -z "$1" ]
then
    CMD="httpperfserv $1"
else
    CMD=httpperfserv
fi

sudo sh -c "cat > /lib/systemd/system/httpperfserv.service <<EOF
[Unit]
Description=Benchmark http server
After=network.target nss-lookup.target
Documentation=https://github.com/vsukhoml/httpperfserv

[Service]
Type=simple
ExecStart=${DIR}/${CMD}

[Install]
WantedBy=multi-user.target
EOF"

sudo systemctl daemon-reload
sudo systemctl enable httpperfserv.service
sudo systemctl start httpperfserv.service
