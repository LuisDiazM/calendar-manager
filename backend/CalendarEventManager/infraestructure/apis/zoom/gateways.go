package zoom

import "github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/apis/zoom/models"

type ZoomApiGateway interface {
	GenerateAccessToken() *models.AccessTokenResponse
	CreateZoomMeeting(token string, userId string, meeting models.MeetingResponse) *models.MeetingResponse
}
