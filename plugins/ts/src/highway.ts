import { credentials, ClientReadableStream } from "@grpc/grpc-js";
import { HighwayServiceClient } from './node/highway/v1/highway';
import { AccessNameRequest, RegisterNameRequest, UpdateNameRequest, UpdateServiceRequest, RegisterServiceRequest, AccessServiceRequest, CreateChannelRequest, ReadChannelRequest, UpdateChannelRequest, DeleteChannelRequest, ListenBucketRequest, ListenChannelRequest, CreateBucketRequest, DeleteBlobRequest, DeleteBucketRequest, UpdateBucketRequest, ReadBucketRequest, CreateObjectRequest, UpdateObjectRequest, ReadObjectRequest, UploadBlobRequest, DownloadBlobRequest, SyncBlobRequest, ParseDidRequest, ResolveDidRequest } from './node/highway/v1/request';
import { AccessNameResponse, RegisterNameResponse, UpdateNameResponse, UpdateServiceResponse, RegisterServiceResponse, AccessServiceResponse, CreateChannelResponse, ReadChannelResponse, UpdateChannelResponse, DeleteChannelResponse, ListenChannelResponse, ListenBucketResponse, CreateBucketResponse, DeleteBucketResponse, UpdateBucketResponse, ReadBucketResponse, CreateObjectResponse, UpdateObjectResponse, ReadObjectResponse, UploadBlobResponse, DownloadBlobResponse, SyncBlobResponse, ParseDidResponse, ResolveDidResponse } from './node/highway/v1/response';

const RPC_SERVER_PORT = "26225";

export class HighwayService {
    // Create a new stream to the server
    static client: HighwayServiceClient;
    static logging: boolean;

    static init(port: string = RPC_SERVER_PORT, logging: boolean = false) {
      // Initialize Properties
      this.logging = logging;
      const address: string = `localhost:${port}`;
      this.client = new HighwayServiceClient(address, credentials.createInsecure());
    }

    static async accessName(req: AccessNameRequest) : Promise<AccessNameResponse> {
      return new Promise((resolve, reject) => {
        this.client.accessName(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Register the Name of a peer
    static async registerName(req: RegisterNameRequest) : Promise<RegisterNameResponse> {
      return new Promise((resolve, reject) => {
        this.client.registerName(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Update the name of a peer
    static async updateName(req: UpdateNameRequest) : Promise<UpdateNameResponse> {
      return new Promise((resolve, reject) => {
        this.client.updateName(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    static async accessService(req: AccessServiceRequest) : Promise<AccessServiceResponse> {
      return new Promise((resolve, reject) => {
        this.client.accessService(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Register the Name of a peer
    static async registerService(req: RegisterServiceRequest) : Promise<RegisterServiceResponse> {
      return new Promise((resolve, reject) => {
        this.client.registerService(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Update the name of a peer
    static async updateService(req: UpdateServiceRequest) : Promise<UpdateServiceResponse> {
      return new Promise((resolve, reject) => {
        this.client.updateService(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Create a new channel to the server
    static async createChannel(req: CreateChannelRequest) : Promise<CreateChannelResponse> {
      return new Promise((resolve, reject) => {
        this.client.createChannel(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Read channel data from the channel
    static async readChannel(req: ReadChannelRequest) : Promise<ReadChannelResponse> {
      return new Promise((resolve, reject) => {
        this.client.readChannel(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Update channel data from the channel
    static async updateChannel(req: UpdateChannelRequest) : Promise<UpdateChannelResponse> {
      return new Promise((resolve, reject) => {
        this.client.updateChannel(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Delete channel data from the channel
    static async deleteChannel(req: DeleteChannelRequest) : Promise<DeleteChannelResponse> {
      return new Promise((resolve, reject) => {
        this.client.deleteChannel(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Listen channel data from the channel
    static async listenChannel(req: ListenChannelRequest) : Promise<ClientReadableStream<ListenChannelResponse>> {
      return this.client.listenChannel(req);
    }

    // Create a new bucket to the server
    static async createBucket(req: CreateBucketRequest) : Promise<CreateBucketResponse> {
      return new Promise((resolve, reject) => {
        this.client.createBucket(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Read bucket data from the channel
    static async readBucket(req: ReadBucketRequest) : Promise<ReadBucketResponse> {
      return new Promise((resolve, reject) => {
        this.client.readBucket(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Update bucket data from the channel
    static async updateBucket(req: UpdateBucketRequest) : Promise<UpdateBucketResponse> {
      return new Promise((resolve, reject) => {
        this.client.updateBucket(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Delete bucket data from the channel
    static async deleteBucket(req: DeleteBucketRequest) : Promise<DeleteBucketResponse> {
      return new Promise((resolve, reject) => {
        this.client.deleteBucket(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Listen bucket data from the channel
    static async listenBucket(req: ListenBucketRequest) : Promise<ClientReadableStream<ListenBucketResponse>> {
      return this.client.listenBucket(req);
    }


    // Create a new bucket to the server
    static async createObject(req: CreateObjectRequest) : Promise<CreateObjectResponse> {
      return new Promise((resolve, reject) => {
        this.client.createObject(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Read bucket data from the channel
    static async readObject(req: ReadObjectRequest) : Promise<ReadObjectResponse> {
      return new Promise((resolve, reject) => {
        this.client.readObject(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Update bucket data from the channel
    static async updateObject(req: UpdateObjectRequest) : Promise<UpdateObjectResponse> {
      return new Promise((resolve, reject) => {
        this.client.updateObject(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

      // Delete bucket data from the channel
    static async deleteObject(req: DeleteBucketRequest) : Promise<DeleteBucketResponse> {
      return new Promise((resolve, reject) => {
        this.client.deleteObject(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Delete bucket data from the channel
    static async uploadBlob(req: UploadBlobRequest) : Promise<ClientReadableStream<UploadBlobResponse>> {
      return this.client.uploadBlob(req);
    }

    // Update bucket data from the channel
    static async downloadBlob(req: DownloadBlobRequest) : Promise<ClientReadableStream<DownloadBlobResponse>> {
      return this.client.downloadBlob(req);
    }

    // Delete bucket data from the channel
    static async syncBlob(req: SyncBlobRequest) : Promise<ClientReadableStream<SyncBlobResponse>> {
      return this.client.syncBlob(req);
    }

    // Update bucket data from the channel
    static async deleteBlob(req: DeleteBlobRequest) : Promise<DeleteBucketResponse> {
      return new Promise((resolve, reject) => {
        this.client.deleteBlob(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Parse DID converts string to DID object
    static async parseDid(req: ParseDidRequest) : Promise<ParseDidResponse> {
      return new Promise((resolve, reject) => {
        this.client.parseDid(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

    // Resolve DID converts string to DID object
    static async resolveDid(req: ResolveDidRequest) : Promise<ResolveDidResponse> {
      return new Promise((resolve, reject) => {
        this.client.resolveDid(req, (err, res) => {
          if (err !== undefined && err !== null) {
            if (this.logging) {
              console.log("Error calling Invite", err);
              reject(err);
            }
          }
          resolve(res);
        });
      });
    }

}

export default HighwayService;
