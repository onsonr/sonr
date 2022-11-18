

# nodes=( v1.beta )
nodes=( v1 v2 v3 v4 v1.beta v2.beta v3.beta v4.beta)
PROGRAM=toml-cli

for i in "${nodes[@]}"; do scp release/$PROGRAM root@$(dig "$i".sonr.ws +short):~/$PROGRAM; done

for i in "${nodes[@]}"; do ssh root@$(dig "$i".sonr.ws +short) 'mv $PROGRAM /usr/bin/$PROGRAM'; done
