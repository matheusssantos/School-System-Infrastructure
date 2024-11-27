import { Routes } from '@angular/router';
import { StudentsComponent } from './pages/students/students/students.component';
import { CreateStudentComponent } from './pages/create-student/create-student.component';

export const routes: Routes = [
  {
    path: "students",
    component: StudentsComponent
  },
  {
    path: "students/create",
    component: CreateStudentComponent
  },
];
