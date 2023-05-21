export interface MeetingModel {
  id?: string;
  invited_guest: string[];
  meeting_date: Date;
  title: string;
  description: string;
  event_duration: number;
  video_conference_link?: string;
  user_id: string;
}
