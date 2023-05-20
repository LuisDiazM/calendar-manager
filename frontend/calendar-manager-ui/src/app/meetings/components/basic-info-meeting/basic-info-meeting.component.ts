import { Component, Input } from '@angular/core';
import { MeetingModel } from '../../entities/meeting.model';
import { EMPTY_MEETING } from '../../utilities/constants';
import { Router } from '@angular/router';

@Component({
  selector: 'app-basic-info-meeting',
  templateUrl: './basic-info-meeting.component.html',
  styleUrls: ['./basic-info-meeting.component.scss'],
})
export class BasicInfoMeetingComponent {
  constructor(private route: Router) {}
  @Input() meeting: MeetingModel = EMPTY_MEETING;

  openMeeting() {
    this.route.navigate([`meetings/id/${this.meeting.id}`]);
  }
}
