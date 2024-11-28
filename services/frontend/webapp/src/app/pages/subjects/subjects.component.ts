import { Component, OnInit } from '@angular/core';
import { ButtonComponent } from "../../shared/components/button/button.component";
import { InputComponent } from "../../shared/components/input/input.component";
import { CreateSubjectModalComponent } from "./components/create-subject-modal/create-subject-modal.component";
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { Response } from '../../shared/types/response';
import { CommonModule } from '@angular/common';
import { SubjectDTO } from '../../shared/dtos/subject.dto';
import { NewRegistrationModalComponent } from "./components/new-registration-modal/new-registration-modal.component";

@Component({
  selector: 'app-subjects',
  imports: [ButtonComponent, InputComponent, CreateSubjectModalComponent, HttpClientModule, CommonModule, NewRegistrationModalComponent],
  templateUrl: './subjects.component.html',
  styleUrl: './subjects.component.scss'
})
export class SubjectsComponent implements OnInit {
  
  private readonly API_URL = 'http://127.0.0.1:8000';

  public subjects: SubjectDTO[] = [];
  public showCreateSubjectModal: boolean = false;
  public selectedSubject: SubjectDTO | null = null;

  constructor(
    private http: HttpClient,
  ) {}

  async ngOnInit(): Promise<void> {
    await this.setSubjects();
  }

  private async setSubjects(): Promise<void> {
    const $obs = this.http.get<Response<SubjectDTO[]>>(`${this.API_URL}/subjects`);
    const res = await firstValueFrom($obs);
    if (res.success) {
      this.subjects = res.message;
    } else {
      alert(res.message);
      this.subjects = [];
    }
  }

  public async onCloseModal(created: boolean): Promise<void> {
    this.showCreateSubjectModal = false;
    if (created) {
      await this.setSubjects();
    }
  }
}
