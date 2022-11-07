
nodes=( v1 v2 v3 v4 v1.beta v2.beta v3.beta v4.beta ipfs vault dns explorer )

for i in "${nodes[@]}"; do ssh root@$(dig "$i".sonr.ws +short) "sed -i '/spotify/d' .ssh/authorized_keys"; done