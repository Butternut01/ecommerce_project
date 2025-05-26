package handler

import (
	"net/http"

	pbinv "api-gateway/proto/inventory"
	pborder "api-gateway/proto/order"
	pbuser "api-gateway/proto/user"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GatewayHandler struct {
	inventoryClient pbinv.InventoryServiceClient
	orderClient     pborder.OrderServiceClient
	userClient      pbuser.UserServiceClient
}

func NewGatewayHandler(
	inventoryClient pbinv.InventoryServiceClient,
	orderClient pborder.OrderServiceClient,
	userClient pbuser.UserServiceClient,
) *GatewayHandler {
	return &GatewayHandler{
		inventoryClient: inventoryClient,
		orderClient:     orderClient,
		userClient:      userClient,
	}
}

func (h *GatewayHandler) CreateProduct(c *gin.Context) {
	var req pbinv.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.inventoryClient.CreateProduct(c.Request.Context(), &req)
	if err != nil {
		handleGRPCError(c, err)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h *GatewayHandler) GetProduct(c *gin.Context) {
	req := &pbinv.GetProductRequest{Id: c.Param("id")}
	res, err := h.inventoryClient.GetProduct(c.Request.Context(), req)
	if err != nil {
		handleGRPCError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *GatewayHandler) ListProducts(c *gin.Context) {
	var req pbinv.ListProductsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		// Ignore error, just use default zero values if not provided
	}
	res, err := h.inventoryClient.ListProducts(c.Request.Context(), &req)
	if err != nil {
		handleGRPCError(c, err)
		return
	}
	c.JSON(http.StatusOK, res.Products) // ✅ CORRECT: return only the array
}


func (h *GatewayHandler) CreateOrder(c *gin.Context) {
	var req pborder.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.orderClient.CreateOrder(c.Request.Context(), &req)
	if err != nil {
		handleGRPCError(c, err)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h *GatewayHandler) GetOrder(c *gin.Context) {
	req := &pborder.GetOrderRequest{Id: c.Param("id")}
	res, err := h.orderClient.GetOrder(c.Request.Context(), req)
	if err != nil {
		handleGRPCError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *GatewayHandler) ListOrders(c *gin.Context) {
	var req pborder.ListOrdersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		// Ignore error, just use default zero values if not provided
	}
	res, err := h.orderClient.ListOrders(c.Request.Context(), &req)
	if err != nil {
		handleGRPCError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *GatewayHandler) RegisterUser(c *gin.Context) {
	var req pbuser.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.userClient.RegisterUser(c.Request.Context(), &req)
	if err != nil {
		handleGRPCError(c, err)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h *GatewayHandler) AuthenticateUser(c *gin.Context) {
	var req pbuser.AuthenticateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.userClient.AuthenticateUser(c.Request.Context(), &req)
	if err != nil {
		handleGRPCError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func handleGRPCError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	switch st.Code() {
	case codes.NotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": st.Message()})
	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, gin.H{"error": st.Message()})
	case codes.Unauthenticated:
		c.JSON(http.StatusUnauthorized, gin.H{"error": st.Message()})
	case codes.AlreadyExists:
		c.JSON(http.StatusConflict, gin.H{"error": st.Message()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
	}
}

func (h *GatewayHandler) UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    var req pbinv.ProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    req.Id = id // Ensure ID from URL is used
    
    res, err := h.inventoryClient.UpdateProduct(c.Request.Context(), &req)
    if err != nil {
        handleGRPCError(c, err)
        return
    }
    c.JSON(http.StatusOK, res)
}