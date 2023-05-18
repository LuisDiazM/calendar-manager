import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { UserService } from '../../services/user.service';
import { SocialUser } from '@abacritt/angularx-social-login';
import { UserCalendarManager } from '../../entities/user.model';
import { tap } from 'rxjs';
import { Router } from '@angular/router';

@Component({
  selector: 'app-registry-user',
  templateUrl: './registry-user.component.html',
  styleUrls: ['./registry-user.component.scss'],
})
export class RegistryUserComponent implements OnInit {
  userLogged: SocialUser | undefined;
  isUserExists: boolean = true;
  constructor(
    private formBuilder: FormBuilder,
    private userService: UserService,
    private router: Router
  ) {}
  public userForm: FormGroup = this.formBuilder.group({
    name: [''],
    email: [''],
    telephone: [''],
    id: [''],
  });
  ngOnInit(): void {
    const user = this.userService.getUserStored();
    if (user !== null) {
      this.userService.getUserFromBackend(user.id).subscribe((userData) => {
        if (userData !== null) {
          this.isUserExists = true;
        } else {
          this.isUserExists = false;
        }
      });
      this.userLogged = user;
      this.userForm.setValue({
        name: user.name,
        email: user.email,
        telephone: '',
        id: user.id,
      });
    }
  }

  onSubmit() {
    const userData = this.userForm.value as UserCalendarManager;
    this.userService.createUser(userData).subscribe(() => {
      this.router.navigateByUrl('users');
    });
  }
}
