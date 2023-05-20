import { MeetingModel } from '../entities/meeting.model';

export const meetingBackendAdapter = (meeting: MeetingModel): MeetingModel => {
  return { ...meeting };
};
