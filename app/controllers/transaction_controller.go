package controllers

import (
	"github.com/dwikiramadhan/money-transaction-api/app/dao"
	"github.com/dwikiramadhan/money-transaction-api/app/entity"
	"github.com/dwikiramadhan/money-transaction-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// Topup method to create a new topup.
// @Description Create a new topup.
// @Summary create a new topup
// @Tags Transaction
// @Accept json
// @Produce json
// @Param Payload body entity.TopupReq true "form Register"
// @Success 201 {object} entity.BaseResponse
// @Failure 400 {array} entity.BaseResponse
// @Failure 404 {array} entity.BaseResponse
// @Security ApiKeyAuth
// @Router /topup [post]
func Topup(c *fiber.Ctx) error {
	// Create a new user auth struct.
	var topupReq entity.TopupReq

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Checking received data from JSON body.
	if err := c.BodyParser(&topupReq); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Validate sign up fields.
	if err := validate.Struct(topupReq); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	result, _, err := dao.Topup(&topupReq)
	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	topupResp := &entity.TopupResp{
		TopUpID:       result.TopUpID,
		AmountTopUp:   result.AmountTopUp,
		BalanceBefore: result.BalanceBefore,
		BalanceAfter:  result.BalanceAfter,
		CreatedDate:   result.CreatedDate,
	}

	// Return status 201 OK.
	baseResponse := &entity.BaseResponse{
		Status: "SUCCESS",
		Result: topupResp,
	}
	return c.Status(fiber.StatusCreated).JSON(baseResponse)
}

// Pay method to create a new payment.
// @Description Create a new payment.
// @Summary create a new payment
// @Tags Transaction
// @Accept json
// @Produce json
// @Param Payload body entity.TopupReq true "form Register"
// @Success 201 {object} entity.BaseResponse
// @Failure 400 {array} entity.BaseResponse
// @Failure 404 {array} entity.BaseResponse
// @Security ApiKeyAuth
// @Router /pay [post]
func Payment(c *fiber.Ctx) error {
	// Create a new user auth struct.
	var payReq entity.PayReq

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Checking received data from JSON body.
	if err := c.BodyParser(&payReq); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Validate sign up fields.
	if err := validate.Struct(payReq); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	result, _, err := dao.Payment(&payReq)
	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	paymentResp := &entity.PaymentResp{
		PaymentID:     result.PaymentID,
		Amount:        result.Amount,
		Remarks:       result.Remarks,
		BalanceBefore: result.BalanceBefore,
		BalanceAfter:  result.BalanceAfter,
		CreatedDate:   result.CreatedDate,
	}

	// Return status 201 OK.
	baseResponse := &entity.BaseResponse{
		Status: "SUCCESS",
		Result: paymentResp,
	}
	return c.Status(fiber.StatusCreated).JSON(baseResponse)
}

// Transfer method to create a new transfer.
// @Description Create a new transfer.
// @Summary create a new transfer
// @Tags Transaction
// @Accept json
// @Produce json
// @Param Payload body entity.TransferReq true "form Register"
// @Success 201 {object} entity.BaseResponse
// @Failure 400 {array} entity.BaseResponse
// @Failure 404 {array} entity.BaseResponse
// @Security ApiKeyAuth
// @Router /transfer [post]
func Transfer(c *fiber.Ctx) error {
	// Create a new user auth struct.
	var transferReq entity.TransferReq

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Checking received data from JSON body.
	if err := c.BodyParser(&transferReq); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Validate sign up fields.
	if err := validate.Struct(transferReq); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	result, _, err := dao.Transfer(&transferReq)
	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	transferResp := &entity.TransferResp{
		TransferID:    result.TransferID,
		Amount:        result.Amount,
		BalanceBefore: result.BalanceBefore,
		BalanceAfter:  result.BalanceAfter,
		CreatedDate:   result.CreatedDate,
	}

	// Return status 201 OK.
	baseResponse := &entity.BaseResponse{
		Status: "SUCCESS",
		Result: transferResp,
	}
	return c.Status(fiber.StatusCreated).JSON(baseResponse)
}
