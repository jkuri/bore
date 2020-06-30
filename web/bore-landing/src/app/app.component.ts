import { Component, OnInit, OnDestroy } from '@angular/core';
import { interval, of, Subscription } from 'rxjs';
import { tap, takeWhile, delay } from 'rxjs/operators';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html'
})
export class AppComponent implements OnInit, OnDestroy {
  command: string = 'bore -s bore.network -lp 8080';
  result: string = 'Generated URL: https://532bbf43.bore.network';
  displayCommand: string = '';
  displayResult: string = '';
  sub = new Subscription();

  ngOnInit(): void {
    let index = 0;
    interval(100)
      .pipe(
        tap(() => index++),
        takeWhile(() => index <= this.command.length)
      )
      .subscribe(
        () => {
          const cmd = this.command.slice(0);
          this.displayCommand = cmd.slice(0, index);
        },
        err => {
          console.error(err);
        },
        () => {
          of(this.result)
            .pipe(delay(500))
            .subscribe(result => (this.displayResult = result));
        }
      );
  }

  ngOnDestroy(): void {
    this.sub.unsubscribe();
  }

  type(): void {
    let index = 0;

    this.sub.add(
      interval(80)
        .pipe(
          tap(() => index++),
          takeWhile(() => index <= this.command.length)
        )
        .subscribe(
          () => {
            const cmd = this.command.slice(0);
            this.displayCommand = cmd.slice(0, index);
          },
          err => {
            console.error(err);
          },
          () => {
            this.sub.add(
              of(this.result)
                .pipe(delay(500))
                .subscribe(result => (this.displayResult = result))
            );
          }
        )
    );
  }
}
