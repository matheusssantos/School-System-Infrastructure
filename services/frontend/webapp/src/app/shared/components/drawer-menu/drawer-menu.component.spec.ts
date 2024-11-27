import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DrawerMenuComponent } from './drawer-menu.component';

describe('DrawerMenuComponent', () => {
  let component: DrawerMenuComponent;
  let fixture: ComponentFixture<DrawerMenuComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DrawerMenuComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DrawerMenuComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
