import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateSubjectModalComponent } from './create-subject-modal.component';

describe('CreateSubjectModalComponent', () => {
  let component: CreateSubjectModalComponent;
  let fixture: ComponentFixture<CreateSubjectModalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreateSubjectModalComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateSubjectModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
