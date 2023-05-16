import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { MeetingsRoutingModule } from './meetings-routing.module';
import { MeetingInfoComponent } from './pages/meeting-info/meeting-info.component';


@NgModule({
  declarations: [
    MeetingInfoComponent
  ],
  imports: [
    CommonModule,
    MeetingsRoutingModule
  ]
})
export class MeetingsModule { }
