import { MeetingModel } from '../entities/meeting.model';

export const EMPTY_MEETING: MeetingModel = {
  description: '',
  event_duration: 15,
  id: '',
  invited_guest: [],
  meeting_date: new Date(),
  title: '',
  user_id: '',
  video_conference_link: '',
};
