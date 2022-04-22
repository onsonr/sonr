package highway

import (
	context "context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	v1 "go.buf.build/grpc/go/sonr-io/core/highway/v1"
)

// UploadBlob uploads a file to IPFS and returns its CID.
func (s *HighwayServer) UploadBlob(ctx context.Context, req *v1.MsgUploadBlob) (*v1.MsgUploadBlobResponse, error) {
	// Read the file at the given path
	buf, err := ioutil.ReadFile(req.Path)
	if err != nil {
		return nil, err
	}

	// Upload the file to ipfsProtocol
	resp, err := s.ipfsProtocol.PutData(buf)
	if err != nil {
		return nil, err
	}

	// Return the response
	return &v1.MsgUploadBlobResponse{
		Code:    200,
		Message: fmt.Sprintf("Succesfully uploaded blob of size %d to IPFS!", len(buf)),
		Cid:     resp.String(),
	}, nil
}

// UploadBlobHTTP uploads a file to IPFS and returns its CID via HTTP.
func (s *HighwayServer) UploadBlobHTTP(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Create Destination path and save the file
	dst := filepath.Join(os.TempDir(), fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename)))
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file to temporary directory",
		})
	}

	// Read the file at the given path
	buf, err := ioutil.ReadFile(dst)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to load the file from temporary directory",
		})
	}

	// Upload the file to ipfsProtocol
	resp, err := s.ipfsProtocol.PutData(buf)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to upload the file to IPFS",
		})
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": fmt.Sprintf("Succesfully uploaded blob of size %d to IPFS!", len(buf)),
		"cid":     resp.String(),
	})
}

// DownloadBlob downloads a file from IPFS given its CID.
func (s *HighwayServer) DownloadBlob(ctx context.Context, req *v1.MsgDownloadBlob) (*v1.MsgDownloadBlobResponse, error) {
	// Upload the file to ipfsProtocol
	resp, err := s.ipfsProtocol.GetData(req.GetCid())
	if err != nil {
		return nil, err
	}

	// Save the file to temporary directory
	dst := filepath.Join(req.GetOutPath(), uuid.New().String())
	if err := ioutil.WriteFile(dst, resp, 0644); err != nil {
		return nil, err
	}

	// Return the response
	return &v1.MsgDownloadBlobResponse{
		Code:    200,
		Message: fmt.Sprintf("Succesfully uploaded blob of size %d to IPFS!", len(resp)),
		Cid:     req.GetCid(),
		Size:    int32(len(resp)),
		Path:    dst,
	}, nil
}

// DownloadBlobHTTP downloads a file from IPFS given its CID via HTTP.
func (s *HighwayServer) DownloadBlobHTTP(c *gin.Context) {
	cid := c.Param("cid")
	if cid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing CID",
		})
		return
	}

	// Download the file from ipfsProtocol
	buf, err := s.ipfsProtocol.GetData(cid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to download the file from IPFS",
		})
	}

	// Save the file to temporary directory
	c.Data(http.StatusOK, "application/octet-stream", buf)
}

// RemoveBlob deletes a file from IPFS given its CID.
func (s *HighwayServer) RemoveBlob(ctx context.Context, req *v1.MsgRemoveBlob) (*v1.MsgRemoveBlobResponse, error) {
	// Upload the file to ipfsProtocol
	err := s.ipfsProtocol.RemoveFile(req.GetCid())
	if err != nil {
		return nil, err
	}

	// Return the response
	return &v1.MsgRemoveBlobResponse{
		Code:    200,
		Message: fmt.Sprintf("Succesfully deleted blob with CID %s from IPFS!", req.GetCid()),
		Cid:     req.GetCid(),
	}, nil
}

// RemoveBlobHTTP downloads a file from IPFS given its CID via HTTP.
func (s *HighwayServer) RemoveBlobHTTP(c *gin.Context) {
	cid := c.Param("cid")
	if cid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing CID",
		})
		return
	}

	// Download the file from ipfsProtocol
	err := s.ipfsProtocol.RemoveFile(cid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to download the file from IPFS",
		})
	}

	// Save the file to temporary directory
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": fmt.Sprintf("Succesfully deleted blob with CID %s from IPFS!", cid),
		"cid":     cid,
	})
}
