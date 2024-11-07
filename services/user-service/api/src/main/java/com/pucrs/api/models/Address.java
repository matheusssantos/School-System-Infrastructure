package com.pucrs.api.models;

import com.fasterxml.jackson.annotation.JsonIgnore;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.OneToOne;
import jakarta.persistence.Table;

@Entity
@Table(name = "addresses")
public class Address {

	@Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  private Integer id;

	@Column(nullable = false)
	private String street;

	@Column(nullable = false)
	private String number;

	@Column()
	private String complement;

	@Column(nullable = false)
	private String zipcode;

	@JsonIgnore
	@OneToOne(mappedBy = "address")
	private User student;

	public Address() {}

	public Address(String street, String number, String complement, String zipcode) {
		this.street = street;
		this.number = number;
		this.complement = complement;
		this.zipcode = zipcode;
	}

	public String getStreet() {
		return this.street;
	}

	public void setStreet(String street) {
		this.street = street;
	}

	public String getNumber() {
		return this.number;
	}

	public void setNumber(String number) {
		this.number = number;
	}

	public String getComplement() {
		return this.complement;
	}

	public void setComplement(String complement) {
		this.complement = complement;
	}

	public String getZipcode() {
		return this.zipcode;
	}

	public void setZipcode(String zipcode) {
		this.zipcode = zipcode;
	}

	public User getStudent() {
		return this.student;
	}

	public Integer getId() {
		return this.id;
	}
}