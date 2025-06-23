package com.mian.jwtauthenticationbackend.repository;

import com.mian.jwtauthenticationbackend.model.User;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;

public interface UserRepository extends JpaRepository<User,Integer> {
    Optional<User> findByUsername(String username);
}