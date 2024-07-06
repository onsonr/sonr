---
title: New Testnet {{ date | date('dddd MMMM Do') }}
labels: enhancement
---
Server IP: {{ env.SERVER_IPV4_ADDR }}
SSH: `ssh -o StrictHostKeyChecking=no root@{{ env.SERVER_IPV4_ADDR }}`
Tag: {{ env.GITHUB_REF_NAME }} / Commit: {{ env.GITHUB_SHA }}
Local-Interchain API: http://{{ env.SERVER_IPV4_ADDR }}:{{ env.LOCALIC_PORT }}