import { SocialAuthService, SocialUser } from '@abacritt/angularx-social-login';
import { Injectable } from '@angular/core';
import { tap } from 'rxjs';
import { ACCESS_TOKEN } from 'src/app/shared/utilities/constants';
import { userManagerSubject } from 'src/app/shared/utilities/services/usermanager';

@Injectable({
  providedIn: 'any',
})
export class UserAuthService {
  user: SocialUser | undefined;

  constructor(private socialAuthService: SocialAuthService) {}

  getAuthenticatedUser() {
    this.socialAuthService.authState
      .pipe(
        tap((user) => (this.user = user)),
        tap((user) => {
          userManagerSubject.setUser(user);
        }),
        tap((user) => {
          console.log(user)
          if (user) {
            sessionStorage.setItem(ACCESS_TOKEN, user.idToken);
          }
        })
      )
      .subscribe();
  }
}
