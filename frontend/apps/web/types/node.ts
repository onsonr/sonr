interface NodeInfo {
    protocol_version: {
        p2p: string;
        block: string;
        app: string;
    };
    id: string;
    listen_addr: string;
    network: string;
    version: string;
    channels: string;
    moniker: string;
    other: {
        tx_index: string;
        rpc_address: string;
    };
}

interface SyncInfo {
    latest_block_hash: string;
    latest_app_hash: string;
    latest_block_height: string;
    latest_block_time: string;
    earliest_block_hash: string;
    earliest_app_hash: string;
    earliest_block_height: string;
    earliest_block_time: string;
    catching_up: boolean;
}

interface ValidatorInfo {
    address: string;
    pub_key: {
        type: string;
        value: string;
    };
    voting_power: string;
}

interface SonrNodeResponse {
    jsonrpc: string;
    id: number;
    result: {
        node_info: NodeInfo;
        sync_info: SyncInfo;
        validator_info: ValidatorInfo;
    };
}

export type {
    NodeInfo,
    SyncInfo,
    ValidatorInfo,
    SonrNodeResponse,
}
