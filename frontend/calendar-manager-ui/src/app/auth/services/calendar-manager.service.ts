import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environments } from 'src/app/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class CalendarManagerService {
  private baseUrl: string = environments.CALENDAR_MANAGER_API;
  constructor(private http: HttpClient) {}

  getUserById(id: string) {
   return this.http.get(`${this.baseUrl}/api/v1/user/${id}`)
  }
}
