import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserInfoComponent } from './pages/user-info/user-info.component';
import { RegistryUserComponent } from './pages/registry-user/registry-user.component';

const routes: Routes = [
  { path: '', component: UserInfoComponent },
  { path: 'registry', component: RegistryUserComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class UsersRoutingModule {}
