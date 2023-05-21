import { MeetingForm } from '../components/meeting-form/meeting-form.component';
import { MeetingModel } from '../entities/meeting.model';

const transform12hTo24 = (hour: number, abbreviation: string): number => {
  if (hour === 12 && abbreviation === 'AM') {
    return 0;
  }
  if (hour < 12 && abbreviation === 'PM') {
    return hour + 12;
  }
  return hour;
};

const transformDateTimePicker = (hour: string, date: Date): Date => {
  const [hours, minutes] = hour.split(/:\s*/);
  const [_, abbreviation] = hour.split(/\s/);
  const copyDate = new Date(date);
  copyDate.setHours(transform12hTo24(parseInt(hours), abbreviation));
  copyDate.setMinutes(parseInt(minutes));
  return copyDate;
};

export const meetingCreateFormToBackend = (
  meeting: MeetingForm
): MeetingModel => {
  return {
    description: meeting.description,
    event_duration: meeting.event_duration,
    invited_guest: meeting.invited_guest,
    title: meeting.title,
    user_id: meeting.user_id,
    meeting_date: transformDateTimePicker(
      meeting.meeting_hour,
      meeting.meeting_date
    ),
    id: meeting.id,
  };
};

const transformToLocal = (timestamp: Date): string => {
  const dateInfo = new Date(timestamp);
  return dateInfo.toLocaleTimeString('en-US', {
    timeZone: 'America/Bogota',
    hour12: true,
    hour: 'numeric',
    minute: 'numeric',
  });
};

export const meetingFromBackendToForm = (
  meeting: MeetingModel
): MeetingForm => {
  return {
    description: meeting.description,
    event_duration: meeting.event_duration,
    invited_guest: meeting.invited_guest,
    title: meeting.title,
    user_id: meeting.user_id,
    meeting_date: meeting.meeting_date,
    meeting_hour: transformToLocal(meeting.meeting_date),
    id: meeting?.id ?? '',
  };
};
