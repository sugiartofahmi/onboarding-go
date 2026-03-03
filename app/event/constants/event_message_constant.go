package constants

const (
	EVENT_PAGINATION_SUCCESS = "Events retrieved successfully"
	EVENT_DETAIL_SUCCESS     = "Event retrieved successfully"
	EVENT_CREATE_SUCCESS     = "Event created successfully"
	EVENT_UPDATE_SUCCESS     = "Event updated successfully"
	EVENT_DELETE_SUCCESS     = "Event deleted successfully"
	EVENT_PUBLISH_SUCCESS    = "Event published successfully"

	EVENT_NOT_FOUND         = "Event not found"
	EVENT_NOT_OWNER         = "You are not the owner of this event"
	EVENT_ALREADY_PUBLISHED = "Event already published"
	EVENT_PUBLISH_NO_TICKET = "Cannot publish event without tickets"
	EVENT_MIN_ONE_TICKET    = "Event must have at least one ticket"

	EVENT_CATEGORY_NOT_FOUND = "Category not found"
	EVENT_TITLE_EXISTS       = "Event title already exists"
	EVENT_SLUG_EXISTS        = "Event slug already exists"
)
