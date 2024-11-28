import { Component, EventEmitter, Input, Output } from '@angular/core';
import { BaseModalComponent } from "../../../../shared/components/base-modal/base-modal.component";
import { ButtonComponent } from "../../../../shared/components/button/button.component";
import { InputComponent } from "../../../../shared/components/input/input.component";
import { CommonModule } from '@angular/common';
import { SubjectDTO } from '../../../../shared/dtos/subject.dto';

@Component({
  selector: 'app-new-registration-modal',
  imports: [BaseModalComponent, ButtonComponent, InputComponent, CommonModule],
  templateUrl: './new-registration-modal.component.html',
  styleUrl: './new-registration-modal.component.scss'
})
export class NewRegistrationModalComponent {
  @Output() closeEmitter: EventEmitter<boolean> = new EventEmitter();

  @Input() subject!: SubjectDTO;

  public async onSave(): Promise<void> {
  }

}
