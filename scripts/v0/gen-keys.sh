nodes=( v1 v2 v3 v4 v1.beta)

#Gen SSH Keys
for i in "${nodes[@]}"; do ssh root@$(dig "$i".sonr.ws +short) '< /dev/zero ssh-keygen -q -N ""'; done

mkdir -p keys
#pull them back
for i in "${nodes[@]}"; do scp root@$(dig "$i".sonr.ws +short):~/.ssh/id_rsa.pub keys/$i.pub; done