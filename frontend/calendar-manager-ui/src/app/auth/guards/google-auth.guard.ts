import { CanMatchFn } from '@angular/router';
import { of } from 'rxjs';
import { USER } from 'src/app/shared/utilities/constants';

export const googleAuthGuardMatch: CanMatchFn = (route, state) => {
  const user = sessionStorage.getItem(USER);
  if (user !== null) {
    return of(true);
  } else {
    return of(false);
  }
};
