import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-base-modal',
  imports: [CommonModule],
  templateUrl: './base-modal.component.html',
  styleUrls: ['./base-modal.component.scss']
})
export class BaseModalComponent {
  @Output() onCloseModal: EventEmitter<void> = new EventEmitter();

  @Input() title: string = '';

  public close() {
    this.onCloseModal.emit();
  }

}