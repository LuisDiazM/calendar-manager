import { Component, ElementRef, Inject, OnInit } from '@angular/core';
import { MeetingModel } from '../../entities/meeting.model';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { EMPTY_MEETING } from '../../utilities/constants';
import { MeetingService } from '../../services/meeting.service';

@Component({
  selector: 'app-edit-meeting-dialog',
  templateUrl: './edit-meeting-dialog.component.html',
  styleUrls: ['./edit-meeting-dialog.component.scss'],
})
export class EditMeetingDialogComponent implements OnInit {
  meetingData: MeetingModel = EMPTY_MEETING;
  meetingDataEdit: MeetingModel = EMPTY_MEETING;
  isCloseModal: boolean = false;
  constructor(
    @Inject(MAT_DIALOG_DATA) public data: MeetingModel,
    private meetingService: MeetingService,
    private host: ElementRef<HTMLElement>
  ) {}
  ngOnInit(): void {
    this.meetingData = { ...this.data };
  }

  receiveMeetingData(event: MeetingModel) {
    this.meetingDataEdit = event;
    this.meetingService.editMeeting(event).subscribe();
  }

  closeModal(event: boolean) {
    this.isCloseModal = event;
    if (event) {
      this.host.nativeElement.remove();
    }
  }
}
