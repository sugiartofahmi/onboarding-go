package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	eventregistrationConstants "event-backend/app/event_registration/constants"
	eventregistrationInterfaces "event-backend/app/event_registration/interfaces"
	roleConstants "event-backend/app/role/constants"
	guards "event-backend/infrastructure/guards"
	"event-backend/infrastructure/middlewares"
	"event-backend/infrastructure/utils"
	uuidValidator "event-backend/infrastructure/validators"
	eventregistrationDtos "event-backend/presentation/http/event_registration/dtos"
)

type EventRegistrationController struct {
	eventRegistrationService eventregistrationInterfaces.EventRegistrationServiceInterface
}

func NewEventRegistrationController(router *gin.Engine, eventRegistrationService eventregistrationInterfaces.EventRegistrationServiceInterface) {
	controller := &EventRegistrationController{
		eventRegistrationService: eventRegistrationService,
	}

	registrationRoute := router.Group("/api/v1/registrations", middlewares.AuthorizationMiddleware())
	registrationRoute.GET("", controller.Pagination())
	registrationRoute.GET("/:id", controller.Detail())
	registrationRoute.POST("", guards.RoleGuard([]string{roleConstants.ATTENDEE, roleConstants.ORGANIZER}), middlewares.ValidateRequestJson[eventregistrationDtos.EventRegistrationCreateRequestDto](), controller.Create())
	registrationRoute.POST("/:id/cancel", guards.RoleGuard([]string{roleConstants.ATTENDEE, roleConstants.ORGANIZER}), middlewares.ValidateRequestJson[eventregistrationDtos.EventRegistrationCancelRequestDto](), controller.Cancel())
}

func (controller *EventRegistrationController) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := &eventregistrationDtos.EventRegistrationQueryRequestDto{}
		httpContext.ShouldBindQuery(dto)

		dto.CurrentUserId = utils.GetAuthUserId(httpContext)
		dto.CurrentUserRoleName = utils.GetAuthUserRoleName(httpContext)

		result := controller.eventRegistrationService.Pagination(ctx, dto)
		meta := utils.PaginationMetaBuilder(dto.Page, dto.PerPage, int(result.Count))
		items := eventregistrationDtos.EventRegistrationResponseDtoFromEntities(result.Data)
		response := utils.SuccessResponsePagination(http.StatusOK, eventregistrationConstants.REGISTRATION_SUCCESS, items, *meta)

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *EventRegistrationController) Detail() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		result := controller.eventRegistrationService.Detail(ctx, id)
		response := utils.SuccessResponse(http.StatusOK, eventregistrationConstants.REGISTRATION_SUCCESS, eventregistrationDtos.EventRegistrationDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *EventRegistrationController) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := httpContext.MustGet(middlewares.RequestBodyJsonKey).(*eventregistrationDtos.EventRegistrationCreateRequestDto)

		currentUserId := utils.GetAuthUserId(httpContext)
		dto.CurrentUserId = &currentUserId

		result := controller.eventRegistrationService.Create(ctx, dto)
		response := utils.SuccessResponse(http.StatusCreated, eventregistrationConstants.REGISTRATION_SUCCESS, eventregistrationDtos.EventRegistrationDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusCreated, response)
	}
}

func (controller *EventRegistrationController) Cancel() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := httpContext.MustGet(middlewares.RequestBodyJsonKey).(*eventregistrationDtos.EventRegistrationCancelRequestDto)

		currentUserId := utils.GetAuthUserId(httpContext)
		dto.CurrentUserId = &currentUserId

		result := controller.eventRegistrationService.Cancel(ctx, dto)
		response := utils.SuccessResponse(http.StatusOK, eventregistrationConstants.REGISTRATION_CANCEL_SUCCESS, eventregistrationDtos.EventRegistrationDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusOK, response)
	}
}
