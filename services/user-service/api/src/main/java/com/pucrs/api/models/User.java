package com.pucrs.api.models;

import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.OneToOne;
import jakarta.persistence.Table;

@Entity
@Table(name = "users")
public class User {

  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  private Integer id;

  @Column(nullable = false)
  private String name;

  @Column(nullable = false)
  private String RG;


  @Column(nullable = false)
  private String type;

  @OneToOne(cascade = CascadeType.ALL)
  @JoinColumn(name = "idAddress", nullable = false)
  private Address address;

  public User() {}

  public User(String name, String RG, Address address, String uuid, String type) {
    this.name = name;
    this.RG = RG;
    this.address = address;
    this.type = type;
  }

  public Address getAddress() {
    return this.address;
  }

  public String getName() {
    return this.name;
  }

  public String getRG() {
    return this.RG;
  }

  public Integer getId() {
    return this.id;
  }
}
