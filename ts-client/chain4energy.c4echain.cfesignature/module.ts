// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgStoreSignature } from "./types/cfesignature/tx";
import { MsgPublishReferencePayloadLink } from "./types/cfesignature/tx";
import { MsgCreateAccount } from "./types/cfesignature/tx";


export { MsgStoreSignature, MsgPublishReferencePayloadLink, MsgCreateAccount };

type sendMsgStoreSignatureParams = {
  value: MsgStoreSignature,
  fee?: StdFee,
  memo?: string
};

type sendMsgPublishReferencePayloadLinkParams = {
  value: MsgPublishReferencePayloadLink,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateAccountParams = {
  value: MsgCreateAccount,
  fee?: StdFee,
  memo?: string
};


type msgStoreSignatureParams = {
  value: MsgStoreSignature,
};

type msgPublishReferencePayloadLinkParams = {
  value: MsgPublishReferencePayloadLink,
};

type msgCreateAccountParams = {
  value: MsgCreateAccount,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgStoreSignature({ value, fee, memo }: sendMsgStoreSignatureParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgStoreSignature: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgStoreSignature({ value: MsgStoreSignature.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgStoreSignature: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgPublishReferencePayloadLink({ value, fee, memo }: sendMsgPublishReferencePayloadLinkParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgPublishReferencePayloadLink: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgPublishReferencePayloadLink({ value: MsgPublishReferencePayloadLink.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgPublishReferencePayloadLink: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateAccount({ value, fee, memo }: sendMsgCreateAccountParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateAccount: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateAccount({ value: MsgCreateAccount.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateAccount: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgStoreSignature({ value }: msgStoreSignatureParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfesignature.MsgStoreSignature", value: MsgStoreSignature.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgStoreSignature: Could not create message: ' + e.message)
			}
		},
		
		msgPublishReferencePayloadLink({ value }: msgPublishReferencePayloadLinkParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfesignature.MsgPublishReferencePayloadLink", value: MsgPublishReferencePayloadLink.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgPublishReferencePayloadLink: Could not create message: ' + e.message)
			}
		},
		
		msgCreateAccount({ value }: msgCreateAccountParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfesignature.MsgCreateAccount", value: MsgCreateAccount.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateAccount: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]>;

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });
		this.tx = txClient({ signer: client.signer, addr: client.env.rpcURL, prefix: client.env.prefix ?? "cosmos" });
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			Chain4EnergyC4EchainCfesignature: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;