import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { firstValueFrom } from 'rxjs';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { ButtonComponent } from '../../shared/components/button/button.component';
import { StudentDTO } from '../../shared/dtos/student.dto';
import { Response } from '../../shared/types/response';
import { CreateStudentModalComponent } from "./components/create-student-modal/create-student-modal.component";

@Component({
  selector: 'app-students',
  imports: [ButtonComponent, CommonModule, HttpClientModule, CreateStudentModalComponent],
  templateUrl: './students.component.html',
  styleUrl: './students.component.scss',
})
export class StudentsComponent implements OnInit {

  private readonly API_URL = 'http://localhost:8080';

  public students: StudentDTO[] = [];
  public showCreateStudentModal: boolean = false;

  constructor(
    private http: HttpClient,
  ) {}

  async ngOnInit(): Promise<void> {
    this.students = await this.getStudents();
  }

  public async getStudents(): Promise<StudentDTO[]> {
    const $obs = this.http.get<Response>(`${this.API_URL}/users`);
    const res = await firstValueFrom($obs);
    if (res.success) {
      return res.message;
    } else {
      alert(res.message);
      return [];
    }
  }

  public async onCloseModal(created: boolean): Promise<void> {
    this.showCreateStudentModal = false;
    if (created) {
      await this.getStudents();
    }
  }
}
