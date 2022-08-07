package api

import (
// "net/http"

// "github.com/gin-gonic/gin"
// t "github.com/sonr-io/sonr/types"
// rt "github.com/sonr-io/sonr/x/registry/types"
)

// // @Summary Create WhoIs Entry
// // @Schemes
// // @Description This method takes a DIDDocument as an input along with the did of the account calling the TX, and verifies the signature. If succesful, and there is no existing WhoIs created for the user or application. Paramaters include: signature, diddocument, address, and whoIsType.
// // @Tags Registry
// // @Produce json
// // @Param 		 data body rt.MsgCreateWhoIs true "Parameters"
// // @Success 	 200  {object}  rt.MsgCreateWhoIsResponse
// // @Failure      500  {string}  message
// // @Router /v1/registry/create/whois [post]
// func (s *HighwayServer) CreateWhoIs(c *gin.Context) {
// 	var req rt.MsgCreateWhoIs
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": t.ErrRequestBody.Error(),
// 		})

// 		return
// 	}
// 	resp, err := s.Cosmos.BroadcastCreateWhoIs(&req)
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{
// 			"error": err.Error(),
// 		})

// 		return
// 	}
// 	// Return the response
// 	c.JSON(http.StatusOK, resp)
// }

// // @Summary Update WhoIs Entry
// // @Schemes
// // @Description This method takes an updated DIDDocument as a JSON buffer along with the signature of the current tx creator, and then verifies the account calling the TX is a controller of the On-chain DIDDocument. If so, the DIDDocument is updated on the blockchain and the transaction is broadcast.
// // @Tags Registry
// // @Produce json
// // @Param 		 data body rt.MsgUpdateWhoIs true "Parameters"
// // @Success 	 200  {object}  rt.MsgUpdateWhoIsResponse
// // @Failure      500  {string}  message
// // @Router /v1/registry/update/whois [post]
// func (s *HighwayServer) UpdateWhoIs(c *gin.Context) {
// 	var req rt.MsgUpdateWhoIs
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		return
// 	}
// 	resp, err := s.Cosmos.BroadcastUpdateWhoIs(&req)
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{
// 			"error": err.Error(),
// 		})

// 		return
// 	}
// 	// Return the response
// 	c.JSON(http.StatusOK, resp)
// }

// // @Summary Deactivate WhoIs Entry
// // @Schemes
// // @Description This method sets the state of a particular WhoIs to be deactivated. In order to Succesfully perform this request, the TX creator and signature must be the same as the WhoIs owner.
// // @Tags Registry
// // @Produce json
// // @Param 		 data body rt.MsgDeactivateWhoIs true "Parameters"
// // @Success 	 200  {object}  rt.MsgDeactivateWhoIsResponse
// // @Failure      500  {string}  message
// // @Router /v1/registry/deactivate/whois [post]
// func (s *HighwayServer) DeactivateWhoIs(c *gin.Context) {
// 	var req rt.MsgDeactivateWhoIs
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		return
// 	}
// 	resp, err := s.Cosmos.BroadcastDeactivateWhoIs(&req)
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{
// 			"error": err.Error(),
// 		})

// 		return
// 	}
// 	// Return the response
// 	c.JSON(http.StatusOK, resp)
// }
