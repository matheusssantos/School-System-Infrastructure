import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { firstValueFrom } from 'rxjs';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { ButtonComponent } from '../../shared/components/button/button.component';
import { StudentDTO } from '../../shared/dtos/student.dto';
import { Response } from '../../shared/types/response';
import { CreateStudentModalComponent } from "./components/create-student-modal/create-student-modal.component";
import { InputComponent } from "../../shared/components/input/input.component";
import { CreateSubjectModalComponent } from "../subjects/components/create-subject-modal/create-subject-modal.component";

@Component({
  selector: 'app-students',
  imports: [ButtonComponent, CommonModule, HttpClientModule, CreateStudentModalComponent, InputComponent],
  templateUrl: './students.component.html',
  styleUrl: './students.component.scss',
})
export class StudentsComponent implements OnInit {

  private readonly API_URL = 'http://localhost:8081';

  public students: StudentDTO[] = [];
  public showCreateStudentModal: boolean = false;

  constructor(
    private http: HttpClient,
  ) {}

  async ngOnInit(): Promise<void> {
    await this.setStudents();
  }

  private async setStudents(): Promise<void> {
    const $obs = this.http.get<Response<StudentDTO[]>>(`${this.API_URL}/users`);
    const res = await firstValueFrom($obs);
    if (res.success) {
      this.students = res.message;
    } else {
      alert(res.message);
      this.students = [];
    }
  }

  public async onCloseModal(created: boolean): Promise<void> {
    this.showCreateStudentModal = false;
    if (created) {
      await this.setStudents();
    }
  }

  public async searchByName(name: string): Promise<void> {
    if (name == '') {
      await this.setStudents();
    } else {
      const $obs = this.http.get<Response<StudentDTO[]>>(`${this.API_URL}/users/name/` + name);
      const res = await firstValueFrom($obs);
      if (res.success) {
        this.students = res.message;
      } else {
        alert(res.message);
      }
    }
  }

  public async searchById(id: string): Promise<void> {
    if (id == '') {
      await this.setStudents();
    } else {
      const $obs = this.http.get<Response<StudentDTO>>(`${this.API_URL}/users/id/` + +id);
      const res = await firstValueFrom($obs);
      if (res.success) {
        this.students = [res.message];
      } else {
        alert(res.message);
      }
    }
  }
}
