import { SocialAuthService, SocialUser } from '@abacritt/angularx-social-login';
import { Injectable } from '@angular/core';
import { Observable, tap } from 'rxjs';
import { ACCESS_TOKEN, USER } from 'src/app/shared/utilities/constants';
import { userManagerSubject } from 'src/app/shared/utilities/services/usermanager';

@Injectable({
  providedIn: 'any',
})
export class UserAuthService {
  user: SocialUser | undefined;

  constructor(private socialAuthService: SocialAuthService) {}

  getAuthenticatedUser(): Observable<SocialUser> {
    return this.socialAuthService.authState.pipe(
      tap((user) => (this.user = user)),
      tap((user) => {
        userManagerSubject.setUser(user);
      }),
      tap((user) => {
        if (user) {
          sessionStorage.setItem(ACCESS_TOKEN, user.idToken);
          sessionStorage.setItem(USER, JSON.stringify(user));
        }
      })
    );
  }

  logout() {
    try {
      this.socialAuthService.signOut();
      sessionStorage.clear();
    } catch (error) {
      console.error(error);
    }
  }

  get userData() {
    return this.user;
  }
}
