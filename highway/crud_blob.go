package highway



// // UploadBlob uploads a blob.
// func (s *HighwayServer) UploadBlob(req *v1.UploadBlobRequest, stream v1.HighwayService_UploadBlobServer) error {
// 	// hash, err := s.ipfs.Upload(req.Path)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	logger.Debug("Uploaded blob to IPFS", "hash")
// 	return nil
// }

// // DownloadBlob downloads a blob.
// func (s *HighwayServer) DownloadBlob(req *v1.DownloadBlobRequest, stream v1.HighwayService_DownloadBlobServer) error {
// 	// path, err := s.ipfs.Download(req.GetDid())
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	logger.Debug("Downloaded blob from IPFS", "path")
// 	return nil
// }

// // SyncBlob synchronizes a blob with remote version.
// func (s *HighwayServer) SyncBlob(req *v1.SyncBlobRequest, stream v1.HighwayService_SyncBlobServer) error {
// 	return nil
// }

// // DeleteBlob deletes a blob.
// func (s *HighwayServer) DeleteBlob(ctx context.Context, req *v1.DeleteBlobRequest) (*v1.DeleteBlobResponse, error) {
// 	return nil, ErrMethodUnimplemented
// }

// // ParseDid parses a DID.
// func (s *HighwayServer) ParseDid(ctx context.Context, req *v1.ParseDidRequest) (*v1.ParseDidResponse, error) {
// 	//d, err := s.node.ParseDid(req.GetDid())
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	return &v1.ParseDidResponse{
// 		Did: "",
// 	}, nil
// }

// // ResolveDid resolves a DID.
// func (s *HighwayServer) ResolveDid(ctx context.Context, req *v1.ResolveDidRequest) (*v1.ResolveDidResponse, error) {
// 	return &v1.ResolveDidResponse{
// 		DidDocument: "",
// 	}, nil
// }
