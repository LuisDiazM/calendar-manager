import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MeetingService } from '../../services/meeting.service';
import { Router } from '@angular/router';

export interface DialogData {
  id: string;
}

@Component({
  selector: 'app-pop-confirm-delete',
  templateUrl: './pop-confirm-delete.component.html',
  styleUrls: ['./pop-confirm-delete.component.scss'],
})
export class PopConfirmDeleteComponent {
  constructor(
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    private meetingService: MeetingService,
    private route: Router
  ) {}

  deleteMeeting() {
    this.meetingService.deleteMeetingById(this.data.id).subscribe();
    this.route.navigate(["/meetings"])
  }
}
