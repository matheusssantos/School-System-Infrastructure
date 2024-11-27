import { Component } from '@angular/core';
import { InputComponent } from "../../shared/components/input/input.component";
import { SelectComponent } from "../../shared/components/select/select.component";
import { ButtonComponent } from "../../shared/components/button/button.component";

@Component({
  selector: 'app-create-student',
  imports: [InputComponent, SelectComponent, ButtonComponent],
  templateUrl: './create-student.component.html',
  styleUrl: './create-student.component.scss'
})
export class CreateStudentComponent {

}
