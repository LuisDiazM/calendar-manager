import { SocialUser } from '@abacritt/angularx-social-login';
import { Component, OnInit } from '@angular/core';
import { tap } from 'rxjs';
import { userManagerSubject } from 'src/app/shared/utilities/services/usermanager';

@Component({
  selector: 'app-meeting-info',
  templateUrl: './meeting-info.component.html',
  styleUrls: ['./meeting-info.component.scss'],
})
export class MeetingInfoComponent implements OnInit {
  user: any;
  ngOnInit(): void {
    userManagerSubject.getUser().subscribe(console.log);
  }
}
