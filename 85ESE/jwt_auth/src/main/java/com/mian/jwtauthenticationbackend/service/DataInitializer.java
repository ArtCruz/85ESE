package com.mian.jwtauthenticationbackend.service;

import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Component;

import com.mian.jwtauthenticationbackend.model.Role;
import com.mian.jwtauthenticationbackend.model.User;
import com.mian.jwtauthenticationbackend.repository.UserRepository;

import jakarta.annotation.PostConstruct;

@Component
public class DataInitializer {

    private final UserRepository userRepository;

    public DataInitializer(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    @PostConstruct
    public void init() {
        if (userRepository.findByUsername("admin").isEmpty()) {
            User user = new User();
            user.setId(1);
            user.setUsername("admin");
            user.setFirstName("Admin");
            user.setLastname("User");
            user.setPassword("admin");
            user.setRole(Role.ADMIN);
            userRepository.save(user);

            System.out.println("Usuário criado:");
            System.out.println("Username: " + user.getUsername());
            System.out.println("Senha: " + user.getPassword());
        }
    }
}
