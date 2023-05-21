import { Component } from '@angular/core';
import { MeetingModel } from '../../entities/meeting.model';
import { EMPTY_MEETING } from '../../utilities/constants';
import { MeetingService } from '../../services/meeting.service';

@Component({
  selector: 'app-write-meeting',
  templateUrl: './write-meeting.component.html',
  styleUrls: ['./write-meeting.component.scss'],
})
export class WriteMeetingComponent {
  meetingBackend: MeetingModel = EMPTY_MEETING;

  constructor(private meetingService: MeetingService) {}

  receiveMeetingData(event: MeetingModel) {
    this.meetingBackend = event;
    this.meetingService.createMeeting(event).subscribe();
  }
}
