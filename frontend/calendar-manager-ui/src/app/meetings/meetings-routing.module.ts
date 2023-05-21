import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MeetingInfoComponent } from './pages/meeting-info/meeting-info.component';
import { MeetingsByUserComponent } from './pages/meetings-by-user/meetings-by-user.component';
import { WriteMeetingComponent } from './pages/write-meeting/write-meeting.component';

const routes: Routes = [
  { path: '', component: MeetingsByUserComponent },
  { path: 'id/:id', component: MeetingInfoComponent },
  { path: 'creator', component: WriteMeetingComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class MeetingsRoutingModule {}
