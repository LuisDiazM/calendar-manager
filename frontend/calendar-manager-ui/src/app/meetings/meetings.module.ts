import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { MeetingsRoutingModule } from './meetings-routing.module';
import { MeetingInfoComponent } from './pages/meeting-info/meeting-info.component';
import { MeetingsByUserComponent } from './pages/meetings-by-user/meetings-by-user.component';
import { MaterialModule } from '../ux-library/material/material.module';
import { PopConfirmDeleteComponent } from './components/pop-confirm-delete/pop-confirm-delete.component';
import { BasicInfoMeetingComponent } from './components/basic-info-meeting/basic-info-meeting.component';


@NgModule({
  declarations: [
    MeetingInfoComponent,
    MeetingsByUserComponent,
    PopConfirmDeleteComponent,
    BasicInfoMeetingComponent
  ],
  imports: [
    CommonModule,
    MeetingsRoutingModule,
    MaterialModule
  ]
})
export class MeetingsModule { }
