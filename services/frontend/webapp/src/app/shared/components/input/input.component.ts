import { CommonModule } from '@angular/common';
import {
  Component,
  EventEmitter,
  Input,
  OnInit,
  Optional,
  Output,
  Self,
} from '@angular/core';
import { ControlValueAccessor, NgControl, Validators } from '@angular/forms';

type InputTypes =
  | 'text'
  | 'number'
  | 'email'
  | 'date'
  | 'password'

@Component({
  selector: 'app-input',
  imports: [CommonModule],
  templateUrl: './input.component.html',
  styleUrls: ['./input.component.scss'],
})
export class InputComponent implements ControlValueAccessor {
  @Output() onChangeEmitter: EventEmitter<string> = new EventEmitter();

  @Input() type: InputTypes = 'text';
  @Input() disabled: boolean = false;
  @Input() placeholder: string = '';
  @Input() label: string = '';
  @Input() required: boolean = false;
  @Input() value: string = '';

  public onChange: any = () => {};
  public onTouched: any = () => {};

  constructor(
    @Self() @Optional() public ngControl: NgControl,
  ) {
    if (this.ngControl != null) {
      this.ngControl.valueAccessor = this;
    }
  }

  public registerOnChange(fn: any): void {
    this.onChange = fn;
  }

  public registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }

  public writeValue(value: string): void {
    console.log(value);
  }

  public setDisabledState(isDisabled: boolean): void {
    this.disabled = isDisabled;
  }

  public getInputValue(event: Event): void {
    this.value = (event.target as HTMLInputElement).value;
    this.onChangeEmitter.emit(this.value);
    this.onChange(this.value);
    this.onTouched();
  }
}