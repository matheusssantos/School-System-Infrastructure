import { Routes } from '@angular/router';
import { StudentsComponent } from './pages/students/students.component';
import { SubjectsComponent } from './pages/subjects/subjects.component';

export const routes: Routes = [
  {
    path: "students",
    component: StudentsComponent
  },
  {
    path: "subjects",
    component: SubjectsComponent
  },
];
