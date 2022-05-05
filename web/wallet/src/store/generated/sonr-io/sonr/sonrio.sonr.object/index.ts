import { txClient, queryClient, MissingWalletError , registry} from './module'
// @ts-ignore
import { SpVuexError } from '@starport/vuex'

import { ObjectDoc } from "./module/types/object/object"
import { ObjectField } from "./module/types/object/object"
import { ObjectFieldArray } from "./module/types/object/object"
import { ObjectFieldText } from "./module/types/object/object"
import { ObjectFieldNumber } from "./module/types/object/object"
import { ObjectFieldBool } from "./module/types/object/object"
import { ObjectFieldTime } from "./module/types/object/object"
import { ObjectFieldGeopoint } from "./module/types/object/object"
import { ObjectFieldBlob } from "./module/types/object/object"
import { ObjectFieldBlockchainAddress } from "./module/types/object/object"
import { Params } from "./module/types/object/params"


export { ObjectDoc, ObjectField, ObjectFieldArray, ObjectFieldText, ObjectFieldNumber, ObjectFieldBool, ObjectFieldTime, ObjectFieldGeopoint, ObjectFieldBlob, ObjectFieldBlockchainAddress, Params };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
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

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				Params: {},

				_Structure: {
						ObjectDoc: getStructure(ObjectDoc.fromPartial({})),
						ObjectField: getStructure(ObjectField.fromPartial({})),
						ObjectFieldArray: getStructure(ObjectFieldArray.fromPartial({})),
						ObjectFieldText: getStructure(ObjectFieldText.fromPartial({})),
						ObjectFieldNumber: getStructure(ObjectFieldNumber.fromPartial({})),
						ObjectFieldBool: getStructure(ObjectFieldBool.fromPartial({})),
						ObjectFieldTime: getStructure(ObjectFieldTime.fromPartial({})),
						ObjectFieldGeopoint: getStructure(ObjectFieldGeopoint.fromPartial({})),
						ObjectFieldBlob: getStructure(ObjectFieldBlob.fromPartial({})),
						ObjectFieldBlockchainAddress: getStructure(ObjectFieldBlockchainAddress.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),

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

		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: sonrio.sonr.object initialized!')
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
					throw new SpVuexError('Subscriptions: ' + e.message)
				}
			})
		},






		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryParams()).data


				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryParams', 'API Node Unavailable. Could not perform query: ' + e.message)

			}
		},


		async sendMsgDeactivateObject({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.MsgDeactivateObject(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee,
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgDeactivateObject:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDeactivateObject:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgReadObject({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgReadObject(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee,
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgReadObject:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgReadObject:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateObject({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateObject(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee,
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgCreateObject:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateObject:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateObject({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateObject(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee,
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgUpdateObject:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateObject:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},

		async MsgDeactivateObject({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.MsgDeactivateObject(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgDeactivateObject:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDeactivateObject:Create', 'Could not create message: ' + e.message)

				}
			}
		},
		async MsgReadObject({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgReadObject(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgReadObject:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgReadObject:Create', 'Could not create message: ' + e.message)

				}
			}
		},
		async MsgCreateObject({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateObject(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgCreateObject:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateObject:Create', 'Could not create message: ' + e.message)

				}
			}
		},
		async MsgUpdateObject({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateObject(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgUpdateObject:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateObject:Create', 'Could not create message: ' + e.message)

				}
			}
		},

	}
}
