package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kataras/go-events"
	metrics "github.com/sonr-io/sonr/internal/highway/x/prometheus"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// UploadBlob uploads a buffer or file to IPFS and returns its CID.
//
// @Summary Upload Blob
// @Schemes
// @Description UploadBlob uploads a buffer or file to IPFS and returns its CID.
// @Tags Blob
// @Produce json
// @Success      200  {string}  cid
// @Failure      500  {string}  message
// @Router /v1/ipfs/upload [post]
func (s *HighwayServer) UploadBlob(c *gin.Context) {
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
		return
	}

	// Read the file at the given path
	buf, err := ioutil.ReadFile(dst)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to load the file from temporary directory",
		})
		return
	}

	// Upload the file to ipfsProtocol
	cid, err := s.ipfsProtocol.PutData(buf)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to upload the file to IPFS",
		})
		return
	}

	defer events.Emit(metrics.ON_BLOB_ADD, "")

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": fmt.Sprintf("Succesfully uploaded blob of size %d to IPFS!", len(buf)),
		"cid":     cid.String(),
	})
}

// DownloadBlob downloads a file from storage given its CID.
// @Summary Download File
// @Schemes
// @Description DownloadBlob downloads a file from storage given its CID.
// @Tags Blob
// @Produce json
// @Success      200  {array}  byte
// @Failure      500  {string}  message
// @Router /v1/ipfs/download/:cid [get]
func (s *HighwayServer) DownloadBlob(c *gin.Context) {
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
		return
	}

	// Save the file to temporary directory
	c.Data(http.StatusOK, "application/octet-stream", buf)
}

// UnpinBlob deletes a file from storage given its CID.
// @Summary Unpin Blob
// @Schemes
// @Description UnpinBlob deletes a file from storage given its CID.
// @Tags Blob
// @Produce json
// @Success      200  {boolean}  success
// @Failure      500  {string}  message
// @Router /v1/ipfs/remove/:cid [post]
func (s *HighwayServer) UnpinBlob(c *gin.Context) {
	cid := c.Param("cid")
	if cid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing CID",
		})
		return
	}

	// Remove the file from ipfsProtocol
	err := s.ipfsProtocol.RemoveFile(cid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to download the file from IPFS",
		})
		return
	}

	defer events.Emit(metrics.ON_BLOB_REMOVE, "")
	// Save the file to temporary directory
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": fmt.Sprintf("Succesfully deleted blob with CID %s from IPFS!", cid),
		"cid":     cid,
	})
}
