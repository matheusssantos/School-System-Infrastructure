import { Component, OnInit } from '@angular/core';
import { ButtonComponent } from "../../../shared/components/button/button.component";
import { StudentDTO } from '../../../shared/dtos/student.dto';
import { CommonModule } from '@angular/common';
import { firstValueFrom } from 'rxjs';
import Response from '../../../shared/types/response';
import { HttpClient, HttpClientModule } from '@angular/common/http';

@Component({
  selector: 'app-students',
  imports: [ButtonComponent, CommonModule, HttpClientModule],
  templateUrl: './students.component.html',
  styleUrl: './students.component.scss',
})
export class StudentsComponent implements OnInit {

  private readonly API_URL = 'http://localhost:8080';

  public students: StudentDTO[] = [];

  constructor(
    private http: HttpClient,
  ) {}

  async ngOnInit(): Promise<void> {
    this.students = await this.getStudents();
  }

  async getStudents(): Promise<StudentDTO[]> {
    const $obs = this.http.get<Response>(`${this.API_URL}/users`);
    const res = await firstValueFrom($obs);
    if (res.success) {
      return res.message;
    } else {
      alert(res.message);
      return [];
    }
  }
}
