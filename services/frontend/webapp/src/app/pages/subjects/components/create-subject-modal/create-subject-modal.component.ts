import { Component, EventEmitter, Output } from '@angular/core';
import { BaseModalComponent } from "../../../../shared/components/base-modal/base-modal.component";
import { InputComponent } from "../../../../shared/components/input/input.component";
import { ButtonComponent } from "../../../../shared/components/button/button.component";
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Response } from '../../../../shared/types/response';
import { firstValueFrom } from 'rxjs';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-create-subject-modal',
  imports: [BaseModalComponent, InputComponent, ButtonComponent, CommonModule, ReactiveFormsModule, FormsModule],
  templateUrl: './create-subject-modal.component.html',
  styleUrl: './create-subject-modal.component.scss'
})
export class CreateSubjectModalComponent {
  @Output() closeEmitter: EventEmitter<boolean> = new EventEmitter();

  public subjectForm!: FormGroup;
  private readonly API_URL = 'http://127.0.0.1:8000';

  constructor(
    private http: HttpClient,
    private formBuilder: FormBuilder,
  ) {}

  ngOnInit(): void {
    this.subjectForm = this.formBuilder.group({
      name: '',
      code: '',
      turn: '',
    });
  }

  public async onSave(): Promise<void> {
    const user = await this.createStudent(this.subjectForm.value);
    if (user) {
      this.subjectForm.reset();
      this.closeEmitter.emit(true);
    }
  }

  private async createStudent(data: any): Promise<any | null> {
    const $obs = this.http.post<Response>(`${this.API_URL}/subjects/create`, data);
    const res = await firstValueFrom($obs);
    if (res.success) {
      return res.message;
    } else {
      alert(res.message);
      return null;
    }
  }

}
