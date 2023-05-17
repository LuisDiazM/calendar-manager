import { Subject } from 'rxjs';

export class UserManager {
  subjectUser$ = new Subject();

  getUser() {
    return this.subjectUser$.asObservable();
  }

  setUser(user: any) {
    this.subjectUser$.next(user);
  }
}
