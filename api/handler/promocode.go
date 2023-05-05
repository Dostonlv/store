package handler

import (
	"app/api/models"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) CreatePromo(ctx *gin.Context) {
	var createPromo models.CreatePromoCode

	err := ctx.ShouldBindJSON(&createPromo)
	if err != nil {
		h.handlerResponse(ctx, "create promo", 400, err.Error())
		return
	}

	id, err := h.storages.PromoCode().Create(context.Background(), &createPromo)
	if err != nil {
		h.handlerResponse(ctx, "storage.promo.create", 500, err.Error())
		return
	}

	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCodePrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(ctx, "storage.promo.getByID", 500, err.Error())
		return
	}

	h.handlerResponse(ctx, "create promo", 201, resp)

}

func (h *Handler) GetPromoByID(ctx *gin.Context) {
	s := ctx.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		h.handlerResponse(ctx, "create promo", 400, err.Error())
		return
	}
	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCodePrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(ctx, "storage.promo.getByID", 500, err.Error())
		return
	}

	h.handlerResponse(ctx, "create promo", 200, resp)
}

func (h *Handler) GetListPromo(ctx *gin.Context) {
	resp, err := h.storages.PromoCode().GetList(context.Background(), &models.GetListPromoCodeRequest{})
	if err != nil {
		h.handlerResponse(ctx, "storage.promo.getByID", 500, err.Error())
		return
	}

	h.handlerResponse(ctx, "create promo", 200, resp)
}

//func (h *Handler) UpdatePromo(c *gin.Context) {
//	var updatePromo models.UpdatePromoCode
//
//	id := c.Param("id")
//
//	err := c.ShouldBindJSON(&updatePromo)
//	if err != nil {
//		h.handlerResponse(c, "update promo", http.StatusBadRequest, err.Error())
//		return
//	}
//
//	idInt, err := strconv.Atoi(id)
//	if err != nil {
//		h.handlerResponse(c, "storage.promo.getByID", http.StatusBadRequest, "id incorrect")
//		return
//	}
//
//	updatePromo.ID = idInt
//
//	rowsAffected, err := h.storages.PromoCode().Update(context.Background(), &updatePromo)
//	if err != nil {
//		h.handlerResponse(c, "storage.promo.update", http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	if rowsAffected <= 0 {
//		h.handlerResponse(c, "storage.promo.update", http.StatusBadRequest, "now rows affected")
//		return
//	}
//
//	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCodePrimaryKey{ID: idInt})
//	if err != nil {
//		h.handlerResponse(c, "storage.promo.getByID", http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	h.handlerResponse(c, "update promo", http.StatusAccepted, resp)
//
//}

func (h *Handler) DeletePromo(ctx *gin.Context) {
	s := ctx.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		h.handlerResponse(ctx, "create promo", 400, err.Error())
		return
	}

	rowsAffected, err := h.storages.PromoCode().Delete(context.Background(), &models.PromoCodePrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(ctx, "storage.promo.delete", 500, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(ctx, "storage.promo.delete", 400, "now rows affected")
		return
	}

	h.handlerResponse(ctx, "create promo", 200, "success")
}
