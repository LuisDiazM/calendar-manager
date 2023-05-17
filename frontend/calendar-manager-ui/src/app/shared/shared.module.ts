import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ErrorComponent } from './pages/error/error.component';
import { NavigationComponent } from './components/navigation/navigation.component';
import { MaterialModule } from '../ux-library/material/material.module';
import { AppRoutingModule } from '../app-routing.module';

@NgModule({
  declarations: [ErrorComponent, NavigationComponent],
  imports: [CommonModule, MaterialModule, AppRoutingModule],
  exports: [NavigationComponent],
})
export class SharedModule {}
