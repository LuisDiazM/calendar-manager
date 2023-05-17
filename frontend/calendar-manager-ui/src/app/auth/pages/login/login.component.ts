import { Component, OnInit } from '@angular/core';
import { UserAuthService } from '../../services/user-auth.service';
import { ACCESS_TOKEN } from 'src/app/shared/utilities/constants';

@Component({
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit {
  isLogged: boolean = false;
  constructor(private userAuth: UserAuthService) {}

  ngOnInit(): void {
    this.userAuth.getAuthenticatedUser();
    const token = sessionStorage.getItem(ACCESS_TOKEN);
    if (token !== null) {
      this.isLogged = true;
    }
  }

  logout(){
    this.userAuth.logout()
  }

}
