import {
  ChangeDetectionStrategy,
  Component,
  OnDestroy,
  OnInit,
} from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { MeetingService } from '../../services/meeting.service';
import { MeetingModel } from '../../entities/meeting.model';
import { Subject, switchMap, takeUntil, tap } from 'rxjs';
import { MatDialog, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { PopConfirmDeleteComponent } from '../../components/pop-confirm-delete/pop-confirm-delete.component';
import { EditMeetingDialogComponent } from '../../components/edit-meeting-dialog/edit-meeting-dialog.component';

@Component({
  selector: 'app-meeting-info',
  templateUrl: './meeting-info.component.html',
  styleUrls: ['./meeting-info.component.scss'],
})
export class MeetingInfoComponent implements OnInit, OnDestroy {
  meetingInfo: MeetingModel | null = null;
  isLoaderMeeting: boolean = true;
  private destroySubject: Subject<void> = new Subject();

  constructor(
    private route: ActivatedRoute,
    private meetingService: MeetingService,
    public dialog: MatDialog
  ) {}

  ngOnInit(): void {
    this.route.params
      .pipe(
        switchMap(({ id }) => this.meetingService.getMeetingById(id)),
        tap((meetingData: MeetingModel | null) => {
          this.meetingInfo = meetingData;
          this.isLoaderMeeting = false;
        }),
        takeUntil(this.destroySubject)
      )
      .subscribe({
        next: () => {},
        error: (e) => {
          console.error(e);
          this.isLoaderMeeting = false;
        },
        complete: () => {},
      });
  }

  ngOnDestroy(): void {
    this.destroySubject.unsubscribe();
  }

  openDialog(
    enterAnimationDuration: string,
    exitAnimationDuration: string
  ): void {
    this.dialog.open(PopConfirmDeleteComponent, {
      width: '250px',
      enterAnimationDuration,
      exitAnimationDuration,
      data: {
        id: this.meetingInfo?.id,
      },
    });
  }

  openDialogEditMeeting(
    enterAnimationDuration: string,
    exitAnimationDuration: string
  ): void {
    this.dialog.open(EditMeetingDialogComponent, {
      width: '60%',
      enterAnimationDuration,
      exitAnimationDuration,
      data: {
        ...this.meetingInfo,
      },
      closeOnNavigation: true,
    });
  }
}
