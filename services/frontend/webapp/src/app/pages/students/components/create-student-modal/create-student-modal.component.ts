import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { BaseModalComponent } from "../../../../shared/components/base-modal/base-modal.component";
import { ButtonComponent } from "../../../../shared/components/button/button.component";
import { InputComponent } from "../../../../shared/components/input/input.component";
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { Response } from '../../../../shared/types/response';
import { StudentDTO } from '../../../../shared/dtos/student.dto';

@Component({
  selector: 'app-create-student-modal',
  imports: [BaseModalComponent, ButtonComponent, InputComponent, ReactiveFormsModule, FormsModule],
  templateUrl: './create-student-modal.component.html',
  styleUrl: './create-student-modal.component.scss'
})
export class CreateStudentModalComponent implements OnInit {
  @Output() closeEmitter: EventEmitter<boolean> = new EventEmitter();

  public studentForm!: FormGroup;
  private readonly API_URL = 'http://localhost:8080';

  constructor(
    private http: HttpClient,
    private formBuilder: FormBuilder,
  ) {}

  ngOnInit(): void {
    this.studentForm = this.formBuilder.group({
      name: '',
      rg: '',
      address: '',
      street: '',
      number: '',
      complement: '',
      zipcode: '',
    });
  }

  public async onSave(): Promise<void> {
    const data: any = {
      ...this.studentForm.value,
      address: { ...this.studentForm.value },
      type: 'Student',
    }
    const user = await this.createStudent(data);
    if (user) {
      this.studentForm.reset();
      this.closeEmitter.emit(true);
    }
  }

  private async createStudent(data: any): Promise<StudentDTO | null> {
    const $obs = this.http.post<Response<StudentDTO>>(`${this.API_URL}/users/create`, data);
    const res = await firstValueFrom($obs);
    if (res.success) {
      return res.message;
    } else {
      alert(res.message);
      return null;
    }
  }
}
