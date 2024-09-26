import Dexie from "dexie";
import {
	PublicKeyCredentialRequestOptions,
	PublicKeyCredentialCreationOptions,
} from "@simplewebauthn/types";

export class Motr {
	constructor(config) {
		this.config = config;
		this.vault = null;
		this.initializeVault();
	}

	initializeVault() {
		const { schema } = this.config;
		this.vault = new Dexie("Vault");
		this.vault.version(schema.version).stores(schema);
	}

	// Account methods
	async insertAccount(accountData) {
		return this.vault.account.add(accountData);
	}

	async getAccount(id) {
		return this.vault.account.get(id);
	}

	async updateAccount(id, accountData) {
		return this.vault.account.update(id, accountData);
	}

	async deleteAccount(id) {
		return this.vault.account.delete(id);
	}

	// Asset methods
	async insertAsset(assetData) {
		return this.vault.asset.add(assetData);
	}

	async getAsset(id) {
		return this.vault.asset.get(id);
	}

	async updateAsset(id, assetData) {
		return this.vault.asset.update(id, assetData);
	}

	async deleteAsset(id) {
		return this.vault.asset.delete(id);
	}

	// Chain methods
	async insertChain(chainData) {
		return this.vault.chain.add(chainData);
	}

	async getChain(id) {
		return this.vault.chain.get(id);
	}

	async updateChain(id, chainData) {
		return this.vault.chain.update(id, chainData);
	}

	async deleteChain(id) {
		return this.vault.chain.delete(id);
	}

	// Credential methods
	async insertCredential(credentialData) {
		const publicKey = await this.createPublicKeyCredential(credentialData);
		credentialData.credentialId = publicKey.id;
		credentialData.publicKey = publicKey.publicKey;
		return this.vault.credential.add(credentialData);
	}

	async getCredential(id) {
		return this.vault.credential.get(id);
	}

	async updateCredential(id, credentialData) {
		return this.vault.credential.update(id, credentialData);
	}

	async deleteCredential(id) {
		return this.vault.credential.delete(id);
	}

	// JWK methods
	async insertJwk(jwkData) {
		return this.vault.jwk.add(jwkData);
	}

	async getJwk(id) {
		return this.vault.jwk.get(id);
	}

	async updateJwk(id, jwkData) {
		return this.vault.jwk.update(id, jwkData);
	}

	async deleteJwk(id) {
		return this.vault.jwk.delete(id);
	}

	// Grant methods
	async insertGrant(grantData) {
		return this.vault.grant.add(grantData);
	}

	async getGrant(id) {
		return this.vault.grant.get(id);
	}

	async updateGrant(id, grantData) {
		return this.vault.grant.update(id, grantData);
	}

	async deleteGrant(id) {
		return this.vault.grant.delete(id);
	}

	// Keyshare methods
	async insertKeyshare(keyshareData) {
		return this.vault.keyshare.add(keyshareData);
	}

	async getKeyshare(id) {
		return this.vault.keyshare.get(id);
	}

	async updateKeyshare(id, keyshareData) {
		return this.vault.keyshare.update(id, keyshareData);
	}

	async deleteKeyshare(id) {
		return this.vault.keyshare.delete(id);
	}

	// PublicKey methods
	async insertPublicKey(publicKeyData) {
		return this.vault.publicKey.add(publicKeyData);
	}

	async getPublicKey(id) {
		return this.vault.publicKey.get(id);
	}

	async updatePublicKey(id, publicKeyData) {
		return this.vault.publicKey.update(id, publicKeyData);
	}

	async deletePublicKey(id) {
		return this.vault.publicKey.delete(id);
	}

	// Profile methods
	async insertProfile(profileData) {
		return this.vault.profile.add(profileData);
	}

	async getProfile(id) {
		return this.vault.profile.get(id);
	}

	async updateProfile(id, profileData) {
		return this.vault.profile.update(id, profileData);
	}

	async deleteProfile(id) {
		return this.vault.profile.delete(id);
	}

	// WebAuthn methods
	async createPublicKeyCredential(): Promise<PublicKeyCredentialRequestOptions> {
		const publicKeyCredentialCreationOptions = {
			challenge: new Uint8Array(32),
			rp: {
				name: this.config.motr.origin,
				id: new URL(this.config.motr.origin).hostname,
			},
			user: {
				id: new TextEncoder().encode(options.subject),
				name: options.subject,
				displayName: options.label,
			},
			pubKeyCredParams: [
				{ alg: -7, type: "public-key" },
				{ alg: -257, type: "public-key" },
			],
			authenticatorSelection: {
				authenticatorAttachment: "platform",
				userVerification: "required",
			},
			timeout: 60000,
			attestation: "direct",
		};

		try {
			const credential = await navigator.credentials.create({
				publicKey: publicKeyCredentialCreationOptions,
			});

			const publicKeyJwk = await crypto.subtle.exportKey(
				"jwk",
				credential.response.getPublicKey(),
			);

			return {
				id: credential.id,
				publicKey: publicKeyJwk,
				type: credential.type,
				transports: credential.response.getTransports(),
			};
		} catch (error) {
			console.error("Error creating credential:", error);
			throw error;
		}
	}

	async getPublicKeyCredential(options) {
		const publicKeyCredentialRequestOptions = {
			challenge: new Uint8Array(32),
			rpId: new URL(this.config.motr.origin).hostname,
			allowCredentials: options.allowCredentials || [],
			userVerification: "required",
			timeout: 60000,
		};

		try {
			const assertion = await navigator.credentials.get({
				publicKey: publicKeyCredentialRequestOptions,
			});

			return {
				id: assertion.id,
				type: assertion.type,
				rawId: new Uint8Array(assertion.rawId),
				response: {
					authenticatorData: new Uint8Array(
						assertion.response.authenticatorData,
					),
					clientDataJSON: new Uint8Array(assertion.response.clientDataJSON),
					signature: new Uint8Array(assertion.response.signature),
					userHandle: new Uint8Array(assertion.response.userHandle),
				},
			};
		} catch (error) {
			console.error("Error getting credential:", error);
			throw error;
		}
	}
}
