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
    private final PasswordEncoder passwordEncoder;

    public DataInitializer(UserRepository userRepository, PasswordEncoder passwordEncoder) {
        this.userRepository = userRepository;
        this.passwordEncoder = passwordEncoder;
        System.out.println("DataInitializer carregado.");
        init();
    }

    @PostConstruct
    public void init() {
            User user = new User();
            user.setId(1);
            user.setUsername("admin");
            user.setFirstName("Admin");
            user.setLastname("User");
            // user.setPassword("admin");
            user.setPassword(passwordEncoder.encode("admin"));
            user.setRole(Role.ADMIN);
            userRepository.save(user);

            System.out.println("Usuário criado:");
            System.out.println("Username: " + user.getUsername());
            System.out.println("Senha: " + user.getPassword());
        }
    
}
