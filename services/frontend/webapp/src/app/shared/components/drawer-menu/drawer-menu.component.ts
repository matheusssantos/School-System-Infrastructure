import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { Router } from '@angular/router';

type DrawerPage = {
  name: string;
  path: string;
  selected: boolean;
}

@Component({
  selector: 'app-drawer-menu',
  imports: [CommonModule],
  templateUrl: './drawer-menu.component.html',
  styleUrl: './drawer-menu.component.scss'
})
export class DrawerMenuComponent {

  constructor(private router: Router) {}

  public pages: DrawerPage[] = [
    {
      name: "Estudantes",
      path: "students",
      selected: false,
    },
    {
      name: "Disciplinas",
      path: "subjects",
      selected: false,
    },
    {
      name: "Matriculas",
      path: "",
      selected: false,
    },
  ];

  public goTo(page: DrawerPage): void {
    this.pages.map(p => p.selected = false);
    page.selected = true;
    this.router.navigate([page.path]);
  }
}
