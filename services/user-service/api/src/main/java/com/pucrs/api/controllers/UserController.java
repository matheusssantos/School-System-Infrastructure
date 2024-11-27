package com.pucrs.api.controllers;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.pucrs.api.dtos.CreateUserDto;
import com.pucrs.api.dtos.ResponseDto;
import com.pucrs.api.models.User;
import com.pucrs.api.services.UserService;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;


@RestController
@RequestMapping("/users")
public class UserController {

  @Autowired
  private UserService userService;

  @PostMapping("/create")
  public ResponseDto<User> createUser(@RequestBody CreateUserDto data) {
    return this.userService.register(data);
  }

  @GetMapping("/id/{id}")
  public ResponseDto<User> findUserByCode(@PathVariable String id) {
    return this.userService.getByid(id);
  }

  @GetMapping("/name/{name}")
  public ResponseDto<List<User>> findUsersByName(@PathVariable String name) {
    return this.userService.getByName(name);
  }
  
  @GetMapping("")
  public ResponseDto<List<User>> findUsers() {
    return this.userService.getAll();
  }
}
