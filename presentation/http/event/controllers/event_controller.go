package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	eventConstants "event-backend/app/event/constants"
	eventInterfaces "event-backend/app/event/interfaces"
	roleConstants "event-backend/app/role/constants"
	"event-backend/infrastructure/constants"
	guards "event-backend/infrastructure/guards"
	"event-backend/infrastructure/middlewares"
	"event-backend/infrastructure/utils"
	uuidValidator "event-backend/infrastructure/validators"
	eventDtos "event-backend/presentation/http/event/dtos"
)

type EventController struct {
	eventService eventInterfaces.EventServiceInterface
}

func NewEventController(router *gin.Engine, eventService eventInterfaces.EventServiceInterface) {
	controller := &EventController{
		eventService: eventService,
	}

	eventRoute := router.Group("/api/v1/events", middlewares.AuthorizationMiddleware())
	eventRoute.GET("", controller.Pagination())
	eventRoute.GET("/:id", controller.Detail())

	organizerRoute := eventRoute.Group("", guards.RoleGuard([]string{roleConstants.ORGANIZER}))
	organizerRoute.POST("", controller.Create())
	organizerRoute.PUT("/:id", controller.Update())
	organizerRoute.DELETE("/:id", controller.Delete())
}

func (controller *EventController) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := eventDtos.AssignEventQueryRequestDto(httpContext)
		result := controller.eventService.Pagination(ctx, dto)
		meta := utils.PaginationMetaBuilder(dto.Page, dto.PerPage, int(result.Count))
		items := eventDtos.EventResponseDtoFromEntities(result.Data)
		response := utils.SuccessResponsePagination(http.StatusOK, eventConstants.EVENT_PAGINATION_SUCCESS, items, *meta)

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *EventController) Detail() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		result := controller.eventService.Detail(ctx, id)
		response := utils.SuccessResponse(http.StatusOK, eventConstants.EVENT_DETAIL_SUCCESS, eventDtos.EventDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *EventController) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := &eventDtos.EventCreateRequestDto{}
		httpContext.ShouldBindJSON(dto)

		currentUserId := utils.GetAuthUserId(httpContext)
		dto.CreatedBy = &currentUserId

		result := controller.eventService.Create(ctx, dto)
		response := utils.SuccessResponse(http.StatusCreated, eventConstants.EVENT_CREATE_SUCCESS, eventDtos.EventDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusCreated, response)
	}
}

func (controller *EventController) Update() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))
		dto := &eventDtos.EventUpdateRequestDto{}
		httpContext.ShouldBindJSON(dto)

		currentUserId := utils.GetAuthUserId(httpContext)
		dto.UpdatedBy = &currentUserId
		dto.Id = id

		result := controller.eventService.Update(ctx, dto)
		response := utils.SuccessResponse(http.StatusOK, eventConstants.EVENT_UPDATE_SUCCESS, eventDtos.EventDetailResponseDtoFromEntity(*result))

		httpContext.JSON(http.StatusOK, response)
	}
}

func (controller *EventController) Delete() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := uuidValidator.ValidateUUID(httpContext.Param("id"))

		userClaims := httpContext.MustGet(constants.AuthUserKey).(utils.JWTUser)
		controller.eventService.SoftDelete(ctx, id, userClaims.Id)
		response := utils.SuccessResponse(http.StatusOK, eventConstants.EVENT_DELETE_SUCCESS, nil)

		httpContext.JSON(http.StatusOK, response)
	}
}
