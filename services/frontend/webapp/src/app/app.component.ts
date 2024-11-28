import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { DrawerMenuComponent } from "./shared/components/drawer-menu/drawer-menu.component";

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, DrawerMenuComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'webapp';
}
