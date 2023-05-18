import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UsersRoutingModule } from './users-routing.module';
import { UserInfoComponent } from './pages/user-info/user-info.component';
import { MaterialModule } from '../ux-library/material/material.module';
import { RegistryUserComponent } from './pages/registry-user/registry-user.component';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [UserInfoComponent, RegistryUserComponent],
  imports: [
    CommonModule,
    UsersRoutingModule,
    MaterialModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
})
export class UsersModule {}
