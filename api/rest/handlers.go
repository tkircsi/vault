package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tkircsi/vault/models"
)

type RESTHandler struct {
	vault models.VaultHandler
}

type PutReqest struct {
	Value string `json:"value"`
	Exp   int    `json:"expire"`
}

func NewRESTHandler(v models.VaultHandler) *RESTHandler {
	return &RESTHandler{vault: v}
}

func (h *RESTHandler) Get(c *gin.Context) {
	token := c.Param("token")
	t, err := h.vault.Get(token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": t.Token,
		"value": t.Value,
	})
}

func (h *RESTHandler) Put(c *gin.Context) {
	var req PutReqest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.vault.Put(req.Value, time.Duration(req.Exp)*time.Second)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token.Token,
	})
}

func (h *RESTHandler) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
