import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-button',
  templateUrl: './button.component.html',
  styleUrls: ['./button.component.scss']
})
export class ButtonComponent {
  @Output() onClick: EventEmitter<void> = new EventEmitter();

  @Input() value: string = "Click";
  @Input() disabled: boolean = false;
}