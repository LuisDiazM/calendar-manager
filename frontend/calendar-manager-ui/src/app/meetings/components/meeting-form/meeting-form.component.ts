import {
  Component,
  ElementRef,
  EventEmitter,
  Input,
  OnInit,
  Output,
} from '@angular/core';
import { MeetingModel } from '../../entities/meeting.model';
import { EMPTY_MEETING } from '../../utilities/constants';
import { FormArray, FormBuilder, FormGroup } from '@angular/forms';
import { COMMA, ENTER } from '@angular/cdk/keycodes';
import { MatChipEditedEvent, MatChipInputEvent } from '@angular/material/chips';
import {
  meetingCreateFormToBackend,
  meetingFromBackendToForm,
} from '../../adapters/meeting.adapter';
import { MeetingService } from '../../services/meeting.service';
import { SocialUser } from '@abacritt/angularx-social-login';
import { Router } from '@angular/router';
import { timer } from 'rxjs';

export interface Invited {
  email: string;
}

export interface MeetingForm {
  title: string;
  description: string;
  meeting_date: Date;
  event_duration: number;
  meeting_hour: string;
  invited_guest: string[];
  user_id: string;
  id: string;
}

@Component({
  selector: 'app-meeting-form',
  templateUrl: './meeting-form.component.html',
  styleUrls: ['./meeting-form.component.scss'],
})
export class MeetingFormComponent implements OnInit {
  user: SocialUser | undefined;
  @Input() meeting: MeetingModel = EMPTY_MEETING;
  @Output() newMeetingEvent = new EventEmitter<MeetingModel>();
  @Output() isCloseForm = new EventEmitter<boolean>();
  Inviteds: Invited[] = [];

  constructor(
    private formBuilder: FormBuilder,
    private meetingService: MeetingService,
    private route: Router,
     
  ) {}

  ngOnInit(): void {
    const meetingConverter = meetingFromBackendToForm(this.meeting);
    this.meetingForm.setValue({
      title: meetingConverter.title,
      description: meetingConverter.description,
      meeting_date: meetingConverter.meeting_date,
      event_duration: meetingConverter.event_duration,
      meeting_hour: meetingConverter.meeting_hour,
    });
    let userInfo = this.meetingService.userData;
    if (userInfo !== null) {
      this.user = userInfo;
    }
    this.Inviteds = meetingConverter.invited_guest.map((element) => {
      return {
        email: element,
      };
    });
  }

  public meetingForm: FormGroup = this.formBuilder.group({
    title: [''],
    description: [''],
    meeting_date: [''],
    event_duration: [''],
    meeting_hour: [''],
  });

  addOnBlur = true;
  readonly separatorKeysCodes = [ENTER, COMMA] as const;

  add(event: MatChipInputEvent): void {
    const value = (event.value || '').trim();
    // Add our Invited
    if (value) {
      this.Inviteds.push({ email: value });
    }
    // Clear the input value
    event.chipInput!.clear();
  }

  remove(Invited: Invited): void {
    const index = this.Inviteds.indexOf(Invited);
    if (index >= 0) {
      this.Inviteds.splice(index, 1);
    }
  }

  edit(Invited: Invited, event: MatChipEditedEvent) {
    const value = event.value.trim();

    // Remove Invited if it no longer has a name
    if (!value) {
      this.remove(Invited);
      return;
    }

    // Edit existing Invited
    const index = this.Inviteds.indexOf(Invited);
    if (index >= 0) {
      this.Inviteds[index].email = value;
    }
  }

  onSave() {
    let meetingValues: MeetingForm = {
      ...this.meetingForm.value,
      invited_guest: [...this.Inviteds.map((invited) => invited.email)],
      user_id: this.user?.id,
      id: this.meeting.id
    };
    const converter = meetingCreateFormToBackend(meetingValues);
    this.newMeetingEvent.emit(converter);
    this.isCloseForm.emit(true);
    timer(600).subscribe(() => {
      this.route.navigate(['meetings']);
    });
  }
}
