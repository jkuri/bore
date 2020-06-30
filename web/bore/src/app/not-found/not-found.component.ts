import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-not-found',
  templateUrl: './not-found.component.html'
})
export class NotFoundComponent implements OnInit {
  tunnelID!: string | null;

  constructor(private route: ActivatedRoute) {}

  ngOnInit(): void {
    this.tunnelID = this.route.snapshot.queryParamMap.get('tunnelID');
  }
}
