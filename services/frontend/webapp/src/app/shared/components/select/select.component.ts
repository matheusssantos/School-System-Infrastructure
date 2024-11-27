import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-select',
  imports: [CommonModule],
  templateUrl: './select.component.html',
  styleUrl: './select.component.scss'
})
export class SelectComponent<DataType> {
  @Output() onSelect: EventEmitter<string> = new EventEmitter(); 
  
  @Input() data: DataType[] = [];
  @Input() keywordProp!: keyof DataType;
  @Input() idProp!: keyof DataType;
  @Input() disabled: boolean = false;
  @Input() label: string = "";
  
  public selectedValue: string = "";

  public updateValue(event: Event): void {
    this.selectedValue = (event.target as HTMLSelectElement).value;
    this.onSelect.emit(this.selectedValue);
  }
}
