import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ErrorComponent } from './shared/pages/error/error.component';
import {
  googleAuthGuardMatch,
} from './auth/guards/google-auth.guard';

const routes: Routes = [
  {
    path: 'auth',
    loadChildren: () => import('./auth/auth.module').then((m) => m.AuthModule),
  },
  {
    path: 'users',
    loadChildren: () =>
      import('./users/users.module').then((m) => m.UsersModule),
    canMatch: [googleAuthGuardMatch],
  },
  {
    path: 'meetings',
    loadChildren: () =>
      import('./meetings/meetings.module').then((m) => m.MeetingsModule),
    canMatch: [googleAuthGuardMatch],
  },
  { path: '', redirectTo: 'auth', pathMatch: 'full' },
  { path: '404', component: ErrorComponent },
  { path: '**', redirectTo: '404' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
