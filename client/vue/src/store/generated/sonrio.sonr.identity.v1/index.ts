import { Client, registry, MissingWalletError } from 'sonr-io-sonr-client-ts'

import { DidDocument } from "sonr-io-sonr-client-ts/sonrio.sonr.identity.v1/types"
import { VerificationMethod } from "sonr-io-sonr-client-ts/sonrio.sonr.identity.v1/types"
import { VerificationRelationship } from "sonr-io-sonr-client-ts/sonrio.sonr.identity.v1/types"
import { Service } from "sonr-io-sonr-client-ts/sonrio.sonr.identity.v1/types"
import { Services } from "sonr-io-sonr-client-ts/sonrio.sonr.identity.v1/types"
import { VerificationMethods } from "sonr-io-sonr-client-ts/sonrio.sonr.identity.v1/types"
import { VerificationRelationships } from "sonr-io-sonr-client-ts/sonrio.sonr.identity.v1/types"
import { Params } from "sonr-io-sonr-client-ts/sonrio.sonr.identity.v1/types"
import { Proof } from "sonr-io-sonr-client-ts/sonrio.sonr.identity.v1/types"
import { JSONWebSignature2020Proof } from "sonr-io-sonr-client-ts/sonrio.sonr.identity.v1/types"
import { VerifiableCredential } from "sonr-io-sonr-client-ts/sonrio.sonr.identity.v1/types"


export { DidDocument, VerificationMethod, VerificationRelationship, Service, Services, VerificationMethods, VerificationRelationships, Params, Proof, JSONWebSignature2020Proof, VerifiableCredential };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				Did: {},
				DidAll: {},
				
				_Structure: {
						DidDocument: getStructure(DidDocument.fromPartial({})),
						VerificationMethod: getStructure(VerificationMethod.fromPartial({})),
						VerificationRelationship: getStructure(VerificationRelationship.fromPartial({})),
						Service: getStructure(Service.fromPartial({})),
						Services: getStructure(Services.fromPartial({})),
						VerificationMethods: getStructure(VerificationMethods.fromPartial({})),
						VerificationRelationships: getStructure(VerificationRelationships.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						Proof: getStructure(Proof.fromPartial({})),
						JSONWebSignature2020Proof: getStructure(JSONWebSignature2020Proof.fromPartial({})),
						VerifiableCredential: getStructure(VerifiableCredential.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getDid: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Did[JSON.stringify(params)] ?? {}
		},
				getDidAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DidAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: sonrio.sonr.identity.v1 initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SonrioSonrIdentityV1.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDid({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SonrioSonrIdentityV1.query.queryDid( key.did)).data
				
					
				commit('QUERY', { query: 'Did', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDid', payload: { options: { all }, params: {...key},query }})
				return getters['getDid']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDid API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDidAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SonrioSonrIdentityV1.query.queryDidAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SonrioSonrIdentityV1.query.queryDidAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'DidAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDidAll', payload: { options: { all }, params: {...key},query }})
				return getters['getDidAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDidAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgDeleteDidDocument({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SonrioSonrIdentityV1.tx.sendMsgDeleteDidDocument({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteDidDocument:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteDidDocument:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateDidDocument({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SonrioSonrIdentityV1.tx.sendMsgUpdateDidDocument({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateDidDocument:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateDidDocument:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateDidDocument({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SonrioSonrIdentityV1.tx.sendMsgCreateDidDocument({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateDidDocument:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateDidDocument:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgDeleteDidDocument({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SonrioSonrIdentityV1.tx.msgDeleteDidDocument({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteDidDocument:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteDidDocument:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateDidDocument({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SonrioSonrIdentityV1.tx.msgUpdateDidDocument({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateDidDocument:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateDidDocument:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateDidDocument({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SonrioSonrIdentityV1.tx.msgCreateDidDocument({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateDidDocument:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateDidDocument:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
