package controllers

import (
	"github.com/dwikiramadhan/money-transaction-api/app/dal"
	"github.com/dwikiramadhan/money-transaction-api/app/dao"
	"github.com/dwikiramadhan/money-transaction-api/app/entity"
	"github.com/dwikiramadhan/money-transaction-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

// Authentications method to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags Authentications
// @Accept json
// @Produce json
// @Param Payload body entity.SignUpReq true "form Register"
// @Success 201 {object} entity.BaseResponse
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	// Create a new user auth struct.
	var signUp entity.SignUpReq

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Checking received data from JSON body.
	if err := c.BodyParser(&signUp); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Validate sign up fields.
	if err := validate.Struct(signUp); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Check email is exist
	_, rows, _ := dao.GetUserByPhoneNumber(*signUp.PhoneNumber)
	if rows > 0 {
		baseResponse := &entity.BaseResponse{
			Status:  "FAILED",
			Message: "Phone number is already registered",
		}
		return c.Status(fiber.StatusBadRequest).JSON(baseResponse)
	}

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(*signUp.Pin), bcrypt.DefaultCost)
	if err != nil {
		baseResponse := &entity.BaseResponse{
			Status: "FAILED",
			Result: "Error on hashing pin",
		}
		return c.Status(fiber.StatusBadRequest).JSON(baseResponse)
	}

	hashedPinString := string(hashedPin)
	signUp.Pin = &hashedPinString

	result, _, err := dao.CreateUser(&signUp)
	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	userResp := &entity.SignUpReResp{
		UserID:      result.UserID,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		PhoneNumber: result.PhoneNumber,
		Address:     result.Address,
		CreatedAt:   result.CreatedAt,
	}

	// Return status 201 OK.
	baseResponse := &entity.BaseResponse{
		Status: "SUCCESS",
		Result: userResp,
	}
	return c.Status(fiber.StatusCreated).JSON(baseResponse)
}

// UserSignIn method to auth user and return access and refresh tokens.
// @Description Auth user and return access and refresh token.
// @Summary auth user and return access and refresh token
// @Tags Authentications
// @Accept json
// @Produce json
// @Param Payload body entity.LoginReq true "form User Login"
// @Success 200 {string} status "ok"
// @Router /login [post]
func Login(c *fiber.Ctx) error {

	formLogin := &entity.LoginReq{}

	// Checking received data from JSON body.
	if err := c.BodyParser(formLogin); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get user by email.
	foundedUser, err := dal.User.Where(dal.User.PhoneNumber.Eq(formLogin.PhoneNumber)).First()
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "User with the given phone number is not found",
		})
	}

	// Compare given user password with stored in found user.
	compareUserPIN := utils.ComparePIN(*foundedUser.Pin, formLogin.PIN)
	if !compareUserPIN {
		// Return, if password is not compare to stored in database.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Phone Number and PIN doesn't match.",
		})
	}

	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(foundedUser.UserID, *foundedUser.PhoneNumber)
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	baseResponse := &entity.BaseResponse{
		Status: "SUCCESS",
		Result: tokens,
	}
	return c.JSON(baseResponse)
}
