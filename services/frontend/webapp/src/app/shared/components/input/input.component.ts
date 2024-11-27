import {
  Component,
  EventEmitter,
  Input,
  Output,
} from '@angular/core';

type InputTypes =
  | 'text'
  | 'number'
  | 'email'
  | 'date'
  | 'password'

@Component({
  selector: 'app-input',
  templateUrl: './input.component.html',
  styleUrls: ['./input.component.scss'],
})
export class InputComponent {
  @Output() onChangeEmitter: EventEmitter<string> = new EventEmitter();

  @Input() type: InputTypes = 'text';
  @Input() disabled: boolean = false;
  @Input() placeholder: string = '';
  @Input() label: string = '';
  @Input() required: boolean = false;
  @Input() value: string = '';

  public getInputValue(event: Event): void {
    const value = (event.target as HTMLInputElement).value;
    this.setValue(value);
  }

  private setValue(value: string): void {
    this.value = value;
    this.onChangeEmitter.emit(value);
  }
}