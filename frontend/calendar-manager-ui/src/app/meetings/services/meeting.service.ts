import { SocialUser } from '@abacritt/angularx-social-login';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environments } from 'src/app/environments/environment';
import { ACCESS_TOKEN, USER } from 'src/app/shared/utilities/constants';
import { MeetingModel } from '../entities/meeting.model';
import { Observable, of, switchMap } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class MeetingService {
  url: string = environments.CALENDAR_MANAGER_API;
  userData: SocialUser | undefined;
  constructor(private http: HttpClient) {
    this.userData = this.getUserStored() ?? undefined;
  }

  get accessToken() {
    return sessionStorage.getItem(ACCESS_TOKEN);
  }

  private getUserStored(): SocialUser | null {
    const user = sessionStorage.getItem(USER);
    if (user !== null) {
      const userData = JSON.parse(user) as SocialUser;
      return userData;
    }
    return null;
  }

  getMeetingById(id: string): Observable<MeetingModel> | Observable<null> {
    const token = this.accessToken;
    if (token !== null) {
      const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
      return this.http.get<MeetingModel>(`${this.url}/api/v1/meeting/${id}`, {
        headers,
      });
    } else {
      return of(null);
    }
  }

  deleteMeetingById(id: string) {
    const token = this.accessToken;
    if (token !== null) {
      const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
      return this.http.delete(`${this.url}/api/v1/meeting/${id}`, { headers });
    } else {
      return of(true);
    }
  }

  getMeetingByUserId(userId: string): Observable<MeetingModel[]> {
    const token = this.accessToken;
    if (token !== null) {
      const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
      const timestamp = new Date();
      timestamp.setMonth(timestamp.getMonth() - 1);
      const params = new HttpParams().set('after', timestamp.toISOString());
      return this.http.get<MeetingModel[]>(
        `${this.url}/api/v1/meetings/user/${userId}`,
        {
          headers,
          params,
        }
      );
    } else {
      return of([]);
    }
  }
}
