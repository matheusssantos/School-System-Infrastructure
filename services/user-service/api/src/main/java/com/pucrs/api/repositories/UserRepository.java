package com.pucrs.api.repositories;

import java.util.List;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.pucrs.api.models.User;

@Repository
public interface UserRepository extends JpaRepository<User, Integer>{

  User findByUuid(String uuid);

  List<User> findByNameContaining(String name);
}
