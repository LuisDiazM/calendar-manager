package entities

type MeetingResponse struct {
	UUID              string   `json:"uuid,omitempty"`
	ID                int64    `json:"id,omitempty"`
	HostID            string   `json:"host_id,omitempty"`
	HostEmail         string   `json:"host_email,omitempty"`
	Topic             string   `json:"topic,omitempty"`
	Type              int64    `json:"type,omitempty"`
	Status            string   `json:"status,omitempty"`
	StartTime         string   `json:"start_time,omitempty"`
	Duration          int64    `json:"duration,omitempty"`
	Timezone          string   `json:"timezone,omitempty"`
	Agenda            string   `json:"agenda,omitempty"`
	CreatedAt         string   `json:"created_at,omitempty"`
	StartURL          string   `json:"start_url,omitempty"`
	JoinURL           string   `json:"join_url,omitempty"`
	Password          string   `json:"password,omitempty"`
	H323Password      string   `json:"h323_password,omitempty"`
	PstnPassword      string   `json:"pstn_password,omitempty"`
	EncryptedPassword string   `json:"encrypted_password,omitempty"`
	Settings          Settings `json:"settings,omitempty"`
	PreSchedule       bool     `json:"pre_schedule,omitempty"`
}

type Settings struct {
	HostVideo                             bool                               `json:"host_video,omitempty"`
	ParticipantVideo                      bool                               `json:"participant_video,omitempty"`
	CNMeeting                             bool                               `json:"cn_meeting,omitempty"`
	InMeeting                             bool                               `json:"in_meeting,omitempty"`
	JoinBeforeHost                        bool                               `json:"join_before_host,omitempty"`
	JbhTime                               int64                              `json:"jbh_time,omitempty"`
	MuteUponEntry                         bool                               `json:"mute_upon_entry,omitempty"`
	Watermark                             bool                               `json:"watermark,omitempty"`
	UsePmi                                bool                               `json:"use_pmi,omitempty"`
	ApprovalType                          int64                              `json:"approval_type,omitempty"`
	Audio                                 string                             `json:"audio,omitempty"`
	AutoRecording                         string                             `json:"auto_recording,omitempty"`
	EnforceLogin                          bool                               `json:"enforce_login,omitempty"`
	EnforceLoginDomains                   string                             `json:"enforce_login_domains,omitempty"`
	AlternativeHosts                      string                             `json:"alternative_hosts,omitempty"`
	AlternativeHostUpdatePolls            bool                               `json:"alternative_host_update_polls,omitempty"`
	CloseRegistration                     bool                               `json:"close_registration,omitempty"`
	ShowShareButton                       bool                               `json:"show_share_button,omitempty"`
	AllowMultipleDevices                  bool                               `json:"allow_multiple_devices,omitempty"`
	RegistrantsConfirmationEmail          bool                               `json:"registrants_confirmation_email,omitempty"`
	WaitingRoom                           bool                               `json:"waiting_room,omitempty"`
	RequestPermissionToUnmuteParticipants bool                               `json:"request_permission_to_unmute_participants,omitempty"`
	RegistrantsEmailNotification          bool                               `json:"registrants_email_notification,omitempty"`
	MeetingAuthentication                 bool                               `json:"meeting_authentication,omitempty"`
	EncryptionType                        string                             `json:"encryption_type,omitempty"`
	ApprovedOrDeniedCountriesOrRegions    ApprovedOrDeniedCountriesOrRegions `json:"approved_or_denied_countries_or_regions,omitempty"`
	BreakoutRoom                          ApprovedOrDeniedCountriesOrRegions `json:"breakout_room,omitempty"`
	AlternativeHostsEmailNotification     bool                               `json:"alternative_hosts_email_notification,omitempty"`
	ShowJoinInfo                          bool                               `json:"show_join_info,omitempty"`
	DeviceTesting                         bool                               `json:"device_testing,omitempty"`
	FocusMode                             bool                               `json:"focus_mode,omitempty"`
	EnableDedicatedGroupChat              bool                               `json:"enable_dedicated_group_chat,omitempty"`
	PrivateMeeting                        bool                               `json:"private_meeting,omitempty"`
	EmailNotification                     bool                               `json:"email_notification,omitempty"`
	HostSaveVideoOrder                    bool                               `json:"host_save_video_order,omitempty"`
	SignLanguageInterpretation            ApprovedOrDeniedCountriesOrRegions `json:"sign_language_interpretation,omitempty"`
	EmailInAttendeeReport                 bool                               `json:"email_in_attendee_report,omitempty"`
}

type ApprovedOrDeniedCountriesOrRegions struct {
	Enable bool `json:"enable,omitempty"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
}
