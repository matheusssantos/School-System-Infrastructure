class CreateUserDTO {
  declare name: string;
  declare rg: string;
  declare address: Address;
} 

class Address {
  declare street: string;
  declare number: string;
  declare complement: string;
  declare zipcode: string;
}