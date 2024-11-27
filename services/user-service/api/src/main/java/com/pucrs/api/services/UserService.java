package com.pucrs.api.services;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.pucrs.api.repositories.UserRepository;
import com.pucrs.api.dtos.CreateUserDto;
import com.pucrs.api.dtos.ResponseDto;
import com.pucrs.api.models.User;

@Service
public class UserService {

  @Autowired
  private UserRepository userRepository;

  public ResponseDto<User> register(CreateUserDto studentDto) {
    try {
      if (
        !studentDto.getType().equals("Teacher") && 
        !studentDto.getType().equals("Student")
      ) {
        return new ResponseDto<>("Tipo de usuário inválido, valores aceitos: 'Teacher' e 'Student");
      }

      final String uuid = this.generateUUID();

      User data = new User(
        studentDto.getName(), 
        studentDto.getRg(), 
        studentDto.getAddress(), 
        uuid, 
        studentDto.getType()
      );
      
      User user = this.userRepository.save(data);
      if (user.equals(null)) {
        return new ResponseDto<>("Erro ao criar usuário");
      }

      return new ResponseDto<>(user);
    } catch (Error error) {
      return new ResponseDto<>("Erro interno no servidor");
    }
  }

  private String generateUUID() {
    return "CODE" + Math.random();
  }

  public ResponseDto<User> getByid(String id) {
    try {
      
      User user = this.userRepository.findById(Integer.parseInt(id));
      if (user == null) {
        return new ResponseDto<>("Usuário não encontrado");
      }
      return new ResponseDto<>(user);
    } catch (Error error) {
      return new ResponseDto<>("Erro interno no servidor");
    }
  }

  public ResponseDto<List<User>> getAll() {
    try {
      return new ResponseDto<>(this.userRepository.findAll());
    } catch (Error error) {
      return new ResponseDto<>("Erro interno no servidor");
    }
  }

  public ResponseDto<List<User>> getByName(String queryName) {
    try {
      return new ResponseDto<>(
        this.userRepository.findByNameContaining(queryName)
      );
    } catch (Error error) {
      return new ResponseDto<>("Erro interno no servidor");
    }
  }
}
