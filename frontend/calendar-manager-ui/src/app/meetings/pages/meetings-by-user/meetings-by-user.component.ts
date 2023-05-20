import { Component, OnInit } from '@angular/core';
import { MeetingService } from '../../services/meeting.service';
import { MeetingModel } from '../../entities/meeting.model';
import { tap } from 'rxjs';

@Component({
  templateUrl: './meetings-by-user.component.html',
  styleUrls: ['./meetings-by-user.component.scss'],
})
export class MeetingsByUserComponent implements OnInit {
  isLoaingMeetings: boolean = true;
  meetingsByUser: MeetingModel[] = [];
  constructor(private meetingService: MeetingService) {}
  ngOnInit(): void {
    const user = this.meetingService.userData;
    if (user) {
      this.meetingService
        .getMeetingByUserId(user?.id)
        .pipe(
          tap((response) => {
            if (response !== null) {
              this.meetingsByUser = response;
            }
          })
        )
        .subscribe({
          next: () => {
            this.isLoaingMeetings = false;
          },
          error: () => {
            this.isLoaingMeetings = false;
          },
        });
    }
  }
}
