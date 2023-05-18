import { SocialUser } from '@abacritt/angularx-social-login';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ACCESS_TOKEN, USER } from 'src/app/shared/utilities/constants';
import { UserCalendarManager } from '../entities/user.model';
import { environments } from 'src/app/environments/environment';
import { of } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  url: string = environments.CALENDAR_MANAGER_API;
  constructor(private http: HttpClient) {}

  get accessToken() {
    return sessionStorage.getItem(ACCESS_TOKEN);
  }

  getUserStored(): SocialUser | null {
    const user = sessionStorage.getItem(USER);
    if (user !== null) {
      const userData = JSON.parse(user) as SocialUser;
      return userData;
    }
    return null;
  }

  getUserFromBackend(id: string) {
    const token = this.accessToken;
    const userCalendarApi: UserCalendarManager = {
      telephone: '',
      created: new Date(),
      email: '',
      name: '',
      id: '',
    };
    if (token !== null) {
      const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
      return this.http.get<UserCalendarManager>(
        `${this.url}/api/v1/user/${id}`,
        { headers }
      );
    } else {
      return of(userCalendarApi);
    }
  }

  createUser(user: UserCalendarManager) {
    const token = this.accessToken;
    if (token !== null) {
      const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
      return this.http.post(`${this.url}/api/v1/users`, user, { headers });
    } else {
      return of(false);
    }
  }
}
