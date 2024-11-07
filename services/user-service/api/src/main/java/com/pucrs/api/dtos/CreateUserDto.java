package com.pucrs.api.dtos;

import com.pucrs.api.models.Address;

public class CreateUserDto {
  private String name;
  private String rg;
  private Address address;
  private String type;

  public String getName() {
    return this.name;
  }

  public String getRg() {
    return this.rg;
  }

  public Address getAddress() {
    return this.address;
  }

  public String getType() {
    return this.type;
  }
}
