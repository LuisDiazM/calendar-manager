package meetings

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	meetingRepo "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/apis/zoom"
)

type WriteZoomMeeting struct {
	api *zoom.ZoomAPI
}

func NewZoomApiRepository(api *zoom.ZoomAPI) meetingRepo.ZoomApiRepository {
	return &WriteZoomMeeting{api: api}
}

func (repo *WriteZoomMeeting) GenerateAccessToken() *entities.AccessTokenResponse {
	return repo.api.GenerateAccessToken()
}

func (repo *WriteZoomMeeting) CreateZoomMeeting(token string, userId string, meeting entities.MeetingResponse) *entities.MeetingResponse {
	return repo.api.CreateZoomMeeting(token, userId, meeting)
}
