## uses pydig v 0.4.0
## Run from within scripts

## Requires toml-cli [https://github.com/MinseokOh/toml-cli] in the release folder
## Requires the release folder to exist
## Maybe needs a genesis.json in the tmep folder
import os
import subprocess
import pathlib
from tracemalloc import stop
import pydig
from os.path import exists

TYPE='dev'

PORT = 26656
NODE_ENDPOINTS = [
    'v1.sonr.ws',
    'v2.sonr.ws',
    'v3.sonr.ws',
    'v4.sonr.ws',
]

TEMP_DIR="./temp"
BIN_DIR="../release"
SCRIPT_DIR="."

def send_command(node, cmd):
    print(f"ssh root@{node} '{cmd}'")
    resp, err =  subprocess.Popen(f"ssh root@{node} '{cmd}'", shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE).communicate()
    resp, err = resp.decode('utf-8').strip(), err.decode('utf-8').strip()
    if err:
        print(err)
        return False
    return resp

def download_file(node, file_remote, file_local):
    resp, err = subprocess.Popen(f'scp root@{node}:{file_remote} {file_local}', shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE).communicate()
    resp,err =  resp.decode('utf-8').strip(), err.decode('utf-8').strip()
    if err:
        print(err)
        return False
    return resp

def upload_file(node, file_remote, file_local):
    resp, err = subprocess.Popen(f'scp {file_local} root@{node}:{file_remote}', shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE).communicate()
    resp,err =  resp.decode('utf-8').strip(), err.decode('utf-8').strip()
    if err:
        print(err)
        return False
    return resp

def download_genesis(node):
    return download_file(node, '~/.sonr/config/genesis.json', f'{TEMP_DIR}/genesis.json')

def upload_genesis(node):
    return upload_file(node, f'~/.sonr/config/genesis.json', f'{TEMP_DIR}/genesis.json')

def download_ssh_key(node):
    return download_file(node, '~/.ssh/id_rsa.pub', f'{TEMP_DIR}/{node}_id_rsa.pub')

def add_key_to_remote_server(node, key_location):
    print(f"Adding key {key_location} to {node}")
    print(f"cat {key_location} | ssh root@{node} 'mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys'")
    resp, err =  subprocess.Popen(f"cat {key_location} | ssh root@{node} 'mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys'", shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE).communicate()
    resp, err = resp.decode('utf-8').strip(), err.decode('utf-8').strip()
    if err:
        print(err)
        return False
    return resp

def get_node_endpoint(node):
    return pydig.query(node, 'A')[0]

def get_node_id(node):
    # Python 3
    return send_command(node, 'sonrd tendermint show-node-id')

# Upload persistent persistent peers to the other nodes
def upload_persistent_peers(node_list):
    for i, node in enumerate(node_list):
        print("Persistent peers for {}".format(node))
        others = node_list[:i] + node_list[i + 1:]
        print(','.join(others))
        # Check commands
        persistent_peers = send_command(node.split(':')[0].split('@')[1], f'toml-cli get /root/.sonr/config/config.toml p2p.persistent_peers')
        if persistent_peers:
            current_peers = persistent_peers.split(',')
            print(current_peers)
            for peer in current_peers:
                others.append(peer)
        print(node)
        send_command(node.split(':')[0].split('@')[1], f'toml-cli set /root/.sonr/config/config.toml p2p.persistent_peers {",".join(others)}')

def upload_toml_cli(node):
    return upload_file(node, '/usr/bin/toml-cli', f'{BIN_DIR}/toml-cli')

def upload_libwasmvm(node):
    #TODO check arch of recieving machine and upload appropriate version
    return upload_file(node, '/usr/lib/libwasmvm.x86_64.so', f'{BIN_DIR}/libwasmvm.x86_64.so')

def stop_sonrd(node):
    return send_command(node, 'systemctl stop sonrd')

def restart_sonrd(node):
    return send_command(node, 'systemctl restart sonrd')

def enable_sonrd(node):
    return send_command(node, 'systemctl enable sonrd')

def disable_sonrd(node):
    return send_command(node, 'systemctl disable sonrd')

def reload_daemon(node):
    return send_command(node, 'systemctl daemon-reload')

def upload_sonrd_service(node):
    print(f"Uploading sonrd service to {node}")
    if(TYPE=='beta'):
        return upload_file(node, '/etc/systemd/system/sonrd.service', f'{SCRIPT_DIR}/beta.sonrd.service')
    else:
        return upload_file(node, '/etc/systemd/system/sonrd.service', f'{SCRIPT_DIR}/sonrd.service')

def upload_setup_chain_dev_script(node):
    return upload_file(node, f'~/setup_chain_dev.sh', f'{SCRIPT_DIR}/setup_chain_dev.sh')

def init_node(node):
    print(f"Initializing node {node}")
    send_command(node, f'sonrd init {node}')

def build_sonrd():
    # Build sonrd
    print("Building sonrd")
    resp, err = subprocess.Popen(f'ignite chain build -t linux:amd64 -o {BIN_DIR} --release', shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE).communicate()
    print(resp, err)
    # Unpack sonrd
    print("Unpacking sonrd")
    resp, err = subprocess.Popen(f'tar -xzvf {BIN_DIR}/sonr_linux_amd64.tar.gz -C {BIN_DIR}', shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE).communicate()
    print(resp, err)


# Uploads, unpacks, and stores the binary sonrd in the remote node
def deploy_sonrd(node):
    # Upload sonrd
    print("Uploading sonrd")
    upload_file(node, f'sonrd', f'{BIN_DIR}/sonrd')
    # Move sonrd to /usr/bin
    print("Moving sonrd to /usr/bin")
    send_command(node, f'mv ~/sonrd /usr/bin/sonrd')

def enable_api(node):
    return send_command(node, 'toml-cli set /root/.sonr/config/app.toml api.enable true')

if __name__ == "__main__":

    # Get the ssh keys from all the nodes, generate them if there isn't one
    for node in NODE_ENDPOINTS:
        print("Getting ssh keys for {}".format(node))
        download_ssh_key(node)
        if not exists(f'{TEMP_DIR}/{node}_id_rsa.pub'):
            # Generate the ssh key if it doesn't exist
            print("Generating ssh key for {}".format(node))
            send_command(node, 'yes "" | ssh-keygen -N ""')
            # Download the ssh key
            download_ssh_key(node)


    # upload the ssh keys to all the other servers
    for i, node in enumerate(NODE_ENDPOINTS):
        for j, node_other in enumerate(NODE_ENDPOINTS):
            if i==j:
                continue
            print(f"Adding ssh key for {TEMP_DIR}/{node}_id_rsa.pub to {node_other}")
            add_key_to_remote_server(node_other, f'{TEMP_DIR}/{node}_id_rsa.pub')

    # Build sonrd
    build_sonrd()

    # Stop all the nodes
    for node in NODE_ENDPOINTS:
        stop_sonrd(node)

    # send data to all the nodes, secondary nodes will get more data than the primary
    for i, node in enumerate(NODE_ENDPOINTS):

        # upload the wasm lib
        upload_libwasmvm(node)
        
        # upload the sonrd binary
        deploy_sonrd(node)

        # Upload the toml-cli binary
        upload_toml_cli(node)


        # make the systemd log directories
        send_command(node, 'mkdir -p /var/log/sonrd/')

        # Upload the sonrd service
        upload_sonrd_service(node)

        # Reload the daemon
        reload_daemon(node)
        
        # Kill the sonrd folder
        send_command(node, 'rm -rf /root/.sonr')

        # Run certain commands on the primary node only
        if i == 0:
            # upload the setup chain dev script
            upload_setup_chain_dev_script(node)

            #mark setup chain dev as executable
            send_command(node, 'chmod +x setup_chain_dev.sh')

            # init the primary node
            send_command(node, './setup_chain_dev.sh')

            # Get the genesis file from the first node
            genesis = download_genesis(NODE_ENDPOINTS[0])
        else:
            # init the node
            init_node(node)

            # delete the genesis file
            send_command(node, 'rm -rf /root/.sonr/config/genesis.json')

            # Upload the genesis file
            upload_genesis(node)
        
        # Enable the api
        enable_api(node)





    # Set the persistent peers in the config file
    nodes = []
    for node in NODE_ENDPOINTS:
        nodes.append(f"{get_node_id(node)}@{get_node_endpoint(node)}:{str(PORT)}")
    upload_persistent_peers(nodes)

    # enable systemctl for sonrd on all the nodes
    for node in NODE_ENDPOINTS:
        enable_sonrd(node)

    # start the nodes
    for node in NODE_ENDPOINTS:
        restart_sonrd(node)