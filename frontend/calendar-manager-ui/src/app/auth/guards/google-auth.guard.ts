import { CanActivateFn, CanMatchFn } from '@angular/router';

export const googleAuthGuardActivate: CanActivateFn = (route, state) => {
  return true;
};

export const googleAuthGuardMatch: CanMatchFn = (route, state) => {
  return true;
};
