import { SocialAuthService, SocialUser } from '@abacritt/angularx-social-login';
import { Injectable } from '@angular/core';
import { switchMap, tap } from 'rxjs';
import { ACCESS_TOKEN } from 'src/app/shared/utilities/constants';
import { userManagerSubject } from 'src/app/shared/utilities/services/usermanager';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'any',
})
export class UserAuthService {
  user: SocialUser | undefined;

  constructor(
    private socialAuthService: SocialAuthService,
  ) {}

  getAuthenticatedUser() {
    this.socialAuthService.authState
      .pipe(
        tap((user) => (this.user = user)),
        tap((user) => {
          userManagerSubject.setUser(user);
        }),
        tap((user) => {
          if (user) {
            sessionStorage.setItem(ACCESS_TOKEN, user.idToken);
          }
        }),
      )
      .subscribe();
  }

  logout() {
    try {
      this.socialAuthService.signOut();
      sessionStorage.clear();
    } catch (error) {
      console.error(error);
    }
  }
}
