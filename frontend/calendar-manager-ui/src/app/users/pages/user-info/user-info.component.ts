import { SocialUser } from '@abacritt/angularx-social-login';
import { Component, OnInit } from '@angular/core';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-user-info',
  templateUrl: './user-info.component.html',
  styleUrls: ['./user-info.component.scss'],
})
export class UserInfoComponent implements OnInit {

  constructor(private userService: UserService) {}
  userLogged: SocialUser | undefined;
  ngOnInit(): void {
    const user = this.userService.getUserStored();
    if (user !== null) {
      this.userLogged = user;
    }
  }


}
