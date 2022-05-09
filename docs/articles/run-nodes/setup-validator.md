---
title: Setup Validator
id: setup-validator
displayed_sidebar: runSidebar
---
#Setup Validator
## Adding a Validator Node

After creating your first node (which will be the proposer) you will need to hit the genesis endpoint of the node.

## Call to Initial Node

### Get genesis file from proposer node

```curl
  GET http://0.0.0.0:26657/genesis
```

**Make sure to replace with the proper proposer IP**

Return Will Look As Follows

```json
{
	"jsonrpc": "2.0",
	"id": -1,
	"result": {
		"genesis": {
			"genesis_time": "2022-04-04T17:57:31.550285Z",
			"chain_id": "sonr",
			"initial_height": "1",
			"consensus_params": {
				"block": {
					"max_bytes": "22020096",
					"max_gas": "-1",
					"time_iota_ms": "1000"
				},
				"evidence": {
					"max_age_num_blocks": "100000",
					"max_age_duration": "172800000000000",
					"max_bytes": "1048576"
				},
				"validator": {
					"pub_key_types": [
						"ed25519"
					]
				},
				"version": {}
			},
			"app_hash": "",
			"app_state": {
				"auth": {
					"params": {
						"max_memo_characters": "256",
						"tx_sig_limit": "7",
						"tx_size_cost_per_byte": "10",
						"sig_verify_cost_ed25519": "590",
						"sig_verify_cost_secp256k1": "1000"
					},
					"accounts": [
						{
							"@type": "/cosmos.auth.v1beta1.BaseAccount",
							"address": "snr19hpavu3wlggjmcmp8wj4n6vehc8vtuc93yk643",
							"pub_key": null,
							"account_number": "0",
							"sequence": "0"
						},
						{
							"@type": "/cosmos.auth.v1beta1.BaseAccount",
							"address": "snr1mtq5lfpfwgk782jwkue52psmwhqrwf93pjk2lz",
							"pub_key": null,
							"account_number": "0",
							"sequence": "0"
						}
					]
				},
				"bank": {
					"params": {
						"send_enabled": [],
						"default_send_enabled": true
					},
					"balances": [
						{
							"address": "snr19hpavu3wlggjmcmp8wj4n6vehc8vtuc93yk643",
							"coins": [
								{
									"denom": "snr",
									"amount": "20000"
								},
								{
									"denom": "stake",
									"amount": "200000000"
								}
							]
						},
						{
							"address": "snr1mtq5lfpfwgk782jwkue52psmwhqrwf93pjk2lz",
							"coins": [
								{
									"denom": "snr",
									"amount": "10000"
								},
								{
									"denom": "stake",
									"amount": "100000000"
								}
							]
						}
					],
					"supply": [],
					"denom_metadata": []
				},
				"blob": {
					"params": {},
					"thereIsList": []
				},
				"bucket": {
					"params": {},
					"whichIsList": []
				},
				"capability": {
					"index": "1",
					"owners": []
				},
				"channel": {
					"howIsList": [],
					"params": {}
				},
				"crisis": {
					"constant_fee": {
						"amount": "1000",
						"denom": "stake"
					}
				},
				"distribution": {
					"delegator_starting_infos": [],
					"delegator_withdraw_infos": [],
					"fee_pool": {
						"community_pool": []
					},
					"outstanding_rewards": [],
					"params": {
						"base_proposer_reward": "0.010000000000000000",
						"bonus_proposer_reward": "0.040000000000000000",
						"community_tax": "0.020000000000000000",
						"withdraw_addr_enabled": true
					},
					"previous_proposer": "",
					"validator_accumulated_commissions": [],
					"validator_current_rewards": [],
					"validator_historical_rewards": [],
					"validator_slash_events": []
				},
				"evidence": {
					"evidence": []
				},
				"feegrant": {
					"allowances": []
				},
				"genutil": {
					"gen_txs": [
						{
							"body": {
								"messages": [
									{
										"@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
										"description": {
											"moniker": "mynode",
											"identity": "",
											"website": "",
											"security_contact": "",
											"details": ""
										},
										"commission": {
											"rate": "0.100000000000000000",
											"max_rate": "0.200000000000000000",
											"max_change_rate": "0.010000000000000000"
										},
										"min_self_delegation": "1",
										"delegator_address": "snr19hpavu3wlggjmcmp8wj4n6vehc8vtuc93yk643",
										"validator_address": "snrvaloper19hpavu3wlggjmcmp8wj4n6vehc8vtuc979fvg8",
										"pubkey": {
											"@type": "/cosmos.crypto.ed25519.PubKey",
											"key": "jtNoxMnorE0Z7h2zet5C2je4Q/uglxXt+Te8sra68BY="
										},
										"value": {
											"denom": "stake",
											"amount": "100000000"
										}
									}
								],
								"memo": "ff2063aec84c4f15eff6c9c673523d33fd7e2160@192.168.0.167:26656",
								"timeout_height": "0",
								"extension_options": [],
								"non_critical_extension_options": []
							},
							"auth_info": {
								"signer_infos": [
									{
										"public_key": {
											"@type": "/cosmos.crypto.secp256k1.PubKey",
											"key": "Ar9URUi9cKAbvoyms5AkJuOM+Hb10kCN8waeERgIUEpG"
										},
										"mode_info": {
											"single": {
												"mode": "SIGN_MODE_DIRECT"
											}
										},
										"sequence": "0"
									}
								],
								"fee": {
									"amount": [],
									"gas_limit": "200000",
									"payer": "",
									"granter": ""
								}
							},
							"signatures": [
								"LQNQURhHPfDDR75YRr0kmszBRTRAG6NNWNpbzEucUIg0nc6wpNOFJ/1aEWZgbJCoYQYiuXYZaNuSORIXRCfgKg=="
							]
						}
					]
				},
				"gov": {
					"deposit_params": {
						"max_deposit_period": "172800s",
						"min_deposit": [
							{
								"amount": "10000000",
								"denom": "stake"
							}
						]
					},
					"deposits": [],
					"proposals": [],
					"starting_proposal_id": "1",
					"tally_params": {
						"quorum": "0.334000000000000000",
						"threshold": "0.500000000000000000",
						"veto_threshold": "0.334000000000000000"
					},
					"votes": [],
					"voting_params": {
						"voting_period": "172800s"
					}
				},
				"ibc": {
					"channel_genesis": {
						"ack_sequences": [],
						"acknowledgements": [],
						"channels": [],
						"commitments": [],
						"next_channel_sequence": "0",
						"receipts": [],
						"recv_sequences": [],
						"send_sequences": []
					},
					"client_genesis": {
						"clients": [],
						"clients_consensus": [],
						"clients_metadata": [],
						"create_localhost": false,
						"next_client_sequence": "0",
						"params": {
							"allowed_clients": [
								"06-solomachine",
								"07-tendermint"
							]
						}
					},
					"connection_genesis": {
						"client_connection_paths": [],
						"connections": [],
						"next_connection_sequence": "0",
						"params": {
							"max_expected_time_per_block": "30000000000"
						}
					}
				},
				"mint": {
					"minter": {
						"annual_provisions": "0.000000000000000000",
						"inflation": "0.130000000000000000"
					},
					"params": {
						"blocks_per_year": "6311520",
						"goal_bonded": "0.670000000000000000",
						"inflation_max": "0.200000000000000000",
						"inflation_min": "0.070000000000000000",
						"inflation_rate_change": "0.130000000000000000",
						"mint_denom": "stake"
					}
				},
				"object": {
					"params": {},
					"whatIsList": []
				},
				"params": null,
				"registry": {
					"params": {},
					"whoIsList": []
				},
				"slashing": {
					"missed_blocks": [],
					"params": {
						"downtime_jail_duration": "600s",
						"min_signed_per_window": "0.500000000000000000",
						"signed_blocks_window": "100",
						"slash_fraction_double_sign": "0.050000000000000000",
						"slash_fraction_downtime": "0.010000000000000000"
					},
					"signing_infos": []
				},
				"staking": {
					"delegations": [],
					"exported": false,
					"last_total_power": "0",
					"last_validator_powers": [],
					"params": {
						"bond_denom": "stake",
						"historical_entries": 10000,
						"max_entries": 7,
						"max_validators": 100,
						"unbonding_time": "1814400s"
					},
					"redelegations": [],
					"unbonding_delegations": [],
					"validators": []
				},
				"transfer": {
					"denom_traces": [],
					"params": {
						"receive_enabled": true,
						"send_enabled": true
					},
					"port_id": "transfer"
				},
				"upgrade": {},
				"vault": {
					"params": {}
				},
				"vesting": {}
			}
		}
	}
}
```

Copy the following from the return

```json
{
			"genesis_time": "2022-04-04T17:57:31.550285Z",
			"chain_id": "sonr",
			"initial_height": "1",
			"consensus_params": {
				"block": {
					"max_bytes": "22020096",
					"max_gas": "-1",
					"time_iota_ms": "1000"
				},
				"evidence": {
					"max_age_num_blocks": "100000",
					"max_age_duration": "172800000000000",
					"max_bytes": "1048576"
				},
				"validator": {
					"pub_key_types": [
						"ed25519"
					]
				},
				"version": {}
			},
			"app_hash": "",
			"app_state": {
				"auth": {
					"params": {
						"max_memo_characters": "256",
						"tx_sig_limit": "7",
						"tx_size_cost_per_byte": "10",
						"sig_verify_cost_ed25519": "590",
						"sig_verify_cost_secp256k1": "1000"
					},
					"accounts": [
						{
							"@type": "/cosmos.auth.v1beta1.BaseAccount",
							"address": "snr19hpavu3wlggjmcmp8wj4n6vehc8vtuc93yk643",
							"pub_key": null,
							"account_number": "0",
							"sequence": "0"
						},
						{
							"@type": "/cosmos.auth.v1beta1.BaseAccount",
							"address": "snr1mtq5lfpfwgk782jwkue52psmwhqrwf93pjk2lz",
							"pub_key": null,
							"account_number": "0",
							"sequence": "0"
						}
					]
				},
				"bank": {
					"params": {
						"send_enabled": [],
						"default_send_enabled": true
					},
					"balances": [
						{
							"address": "snr19hpavu3wlggjmcmp8wj4n6vehc8vtuc93yk643",
							"coins": [
								{
									"denom": "snr",
									"amount": "20000"
								},
								{
									"denom": "stake",
									"amount": "200000000"
								}
							]
						},
						{
							"address": "snr1mtq5lfpfwgk782jwkue52psmwhqrwf93pjk2lz",
							"coins": [
								{
									"denom": "snr",
									"amount": "10000"
								},
								{
									"denom": "stake",
									"amount": "100000000"
								}
							]
						}
					],
					"supply": [],
					"denom_metadata": []
				},
				"blob": {
					"params": {},
					"thereIsList": []
				},
				"bucket": {
					"params": {},
					"whichIsList": []
				},
				"capability": {
					"index": "1",
					"owners": []
				},
				"channel": {
					"howIsList": [],
					"params": {}
				},
				"crisis": {
					"constant_fee": {
						"amount": "1000",
						"denom": "stake"
					}
				},
				"distribution": {
					"delegator_starting_infos": [],
					"delegator_withdraw_infos": [],
					"fee_pool": {
						"community_pool": []
					},
					"outstanding_rewards": [],
					"params": {
						"base_proposer_reward": "0.010000000000000000",
						"bonus_proposer_reward": "0.040000000000000000",
						"community_tax": "0.020000000000000000",
						"withdraw_addr_enabled": true
					},
					"previous_proposer": "",
					"validator_accumulated_commissions": [],
					"validator_current_rewards": [],
					"validator_historical_rewards": [],
					"validator_slash_events": []
				},
				"evidence": {
					"evidence": []
				},
				"feegrant": {
					"allowances": []
				},
				"genutil": {
					"gen_txs": [
						{
							"body": {
								"messages": [
									{
										"@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
										"description": {
											"moniker": "mynode",
											"identity": "",
											"website": "",
											"security_contact": "",
											"details": ""
										},
										"commission": {
											"rate": "0.100000000000000000",
											"max_rate": "0.200000000000000000",
											"max_change_rate": "0.010000000000000000"
										},
										"min_self_delegation": "1",
										"delegator_address": "snr19hpavu3wlggjmcmp8wj4n6vehc8vtuc93yk643",
										"validator_address": "snrvaloper19hpavu3wlggjmcmp8wj4n6vehc8vtuc979fvg8",
										"pubkey": {
											"@type": "/cosmos.crypto.ed25519.PubKey",
											"key": "jtNoxMnorE0Z7h2zet5C2je4Q/uglxXt+Te8sra68BY="
										},
										"value": {
											"denom": "stake",
											"amount": "100000000"
										}
									}
								],
								"memo": "ff2063aec84c4f15eff6c9c673523d33fd7e2160@192.168.0.167:26656",
								"timeout_height": "0",
								"extension_options": [],
								"non_critical_extension_options": []
							},
							"auth_info": {
								"signer_infos": [
									{
										"public_key": {
											"@type": "/cosmos.crypto.secp256k1.PubKey",
											"key": "Ar9URUi9cKAbvoyms5AkJuOM+Hb10kCN8waeERgIUEpG"
										},
										"mode_info": {
											"single": {
												"mode": "SIGN_MODE_DIRECT"
											}
										},
										"sequence": "0"
									}
								],
								"fee": {
									"amount": [],
									"gas_limit": "200000",
									"payer": "",
									"granter": ""
								}
							},
							"signatures": [
								"LQNQURhHPfDDR75YRr0kmszBRTRAG6NNWNpbzEucUIg0nc6wpNOFJ/1aEWZgbJCoYQYiuXYZaNuSORIXRCfgKg=="
							]
						}
					]
				},
				"gov": {
					"deposit_params": {
						"max_deposit_period": "172800s",
						"min_deposit": [
							{
								"amount": "10000000",
								"denom": "stake"
							}
						]
					},
					"deposits": [],
					"proposals": [],
					"starting_proposal_id": "1",
					"tally_params": {
						"quorum": "0.334000000000000000",
						"threshold": "0.500000000000000000",
						"veto_threshold": "0.334000000000000000"
					},
					"votes": [],
					"voting_params": {
						"voting_period": "172800s"
					}
				},
				"ibc": {
					"channel_genesis": {
						"ack_sequences": [],
						"acknowledgements": [],
						"channels": [],
						"commitments": [],
						"next_channel_sequence": "0",
						"receipts": [],
						"recv_sequences": [],
						"send_sequences": []
					},
					"client_genesis": {
						"clients": [],
						"clients_consensus": [],
						"clients_metadata": [],
						"create_localhost": false,
						"next_client_sequence": "0",
						"params": {
							"allowed_clients": [
								"06-solomachine",
								"07-tendermint"
							]
						}
					},
					"connection_genesis": {
						"client_connection_paths": [],
						"connections": [],
						"next_connection_sequence": "0",
						"params": {
							"max_expected_time_per_block": "30000000000"
						}
					}
				},
				"mint": {
					"minter": {
						"annual_provisions": "0.000000000000000000",
						"inflation": "0.130000000000000000"
					},
					"params": {
						"blocks_per_year": "6311520",
						"goal_bonded": "0.670000000000000000",
						"inflation_max": "0.200000000000000000",
						"inflation_min": "0.070000000000000000",
						"inflation_rate_change": "0.130000000000000000",
						"mint_denom": "stake"
					}
				},
				"object": {
					"params": {},
					"whatIsList": []
				},
				"params": null,
				"registry": {
					"params": {},
					"whoIsList": []
				},
				"slashing": {
					"missed_blocks": [],
					"params": {
						"downtime_jail_duration": "600s",
						"min_signed_per_window": "0.500000000000000000",
						"signed_blocks_window": "100",
						"slash_fraction_double_sign": "0.050000000000000000",
						"slash_fraction_downtime": "0.010000000000000000"
					},
					"signing_infos": []
				},
				"staking": {
					"delegations": [],
					"exported": false,
					"last_total_power": "0",
					"last_validator_powers": [],
					"params": {
						"bond_denom": "stake",
						"historical_entries": 10000,
						"max_entries": 7,
						"max_validators": 100,
						"unbonding_time": "1814400s"
					},
					"redelegations": [],
					"unbonding_delegations": [],
					"validators": []
				},
				"transfer": {
					"denom_traces": [],
					"params": {
						"receive_enabled": true,
						"send_enabled": true
					},
					"port_id": "transfer"
				},
				"upgrade": {},
				"vault": {
					"params": {}
				},
				"vesting": {}
			}
		}
```

## Install sonrd Binary to Validator Node

Download the sonrd binary and run the init command to create the following files

You can also copy app.toml, client.toml, config.toml & genesis.json from the proposer node as a secondary choice.

```markdown
The files go as follows:
    1. addrbook.json
    2. app.toml
    3. client.toml
    4. config.toml
    5. genesis.json
    6. node_key.json
    7. priv_validator_key.json
```

## Tendermint Documentation

Install Tendermint
[Tendermint Installation](https://docs.tendermint.com/v0.35/introduction/install.html)

Follow creation of validator node via Documentation
[Tendermint Validator Documentation](https://docs.tendermint.com/v0.35/nodes/validators.html)

## Deployment from Validator Node

Run Initialization of Validator

```shell
  tendermint init validator
```

Use the files created from .tendermint/config/genesis.json and copy the validator array into the proposer node via the genesis.json file

## Deployment from Proposer Node

Grab the Node ID from the proposer

```shell
  ./sonrd tendermint show-node-id
```

## Proposer -> config.toml

Head to line 188 in config.toml

### Adding the peer

The peer structure is the node ID @ IP of validator node and port. Example below

```Text
  ff2063aec84c4f15eff6c9c673523d33fd7e2160@0.0.0.0:26656
```

**Make sure to add the proper IP address of the validator node**

## Adding Persistent Peer into the Validator Node

Follow the same steps as above, head to the genesis.json file on the validator node, copy the genesis.json file from the proposer node and paste into the file.

### Add the proposer node as a peer on the validator node

Follow the same steps, use the node ID from the proposer node and paste NODE\_ID\:IP\:PORT into line 188 on coonfig.toml

## Run the validator node

```shell
  ./sonrd start
```

You should see the validator start to catch up to the proposer node and start validating new blocks!
