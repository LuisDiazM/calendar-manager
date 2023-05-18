import { SocialUser } from '@abacritt/angularx-social-login';
import { UserCalendarManager } from '../entities/user.model';

export const adapterUserBackend = (user: SocialUser): UserCalendarManager => {
  return {
    created: new Date(),
    email: user.email,
    id: user.id,
    name: user.name,
    telephone: '',
  };
};
