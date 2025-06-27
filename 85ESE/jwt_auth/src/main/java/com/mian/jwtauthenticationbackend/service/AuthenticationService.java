package com.mian.jwtauthenticationbackend.service;

import com.mian.jwtauthenticationbackend.model.AuthenticationResponse;
import com.mian.jwtauthenticationbackend.model.User;
import com.mian.jwtauthenticationbackend.repository.UserRepository;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.Optional;
@Service
public class AuthenticationService {
    private final UserRepository repository;
    // private final PasswordEncoder passwordEncoder;
    private final JwtService jwtService;
    private final AuthenticationManager authenticationManager;

    public AuthenticationService(UserRepository repository, JwtService jwtService, AuthenticationManager authenticationManager) {
        this.repository = repository;
        // this.passwordEncoder = passwordEncoder;
        this.jwtService = jwtService;
        this.authenticationManager = authenticationManager;
    }
    public AuthenticationResponse register(User request){
        User user = new User();
        user.setFirstName(request.getFirstName());
        user.setLastname((request.getLastname()));
        user.setUsername(request.getUsername());
        user.setPassword((request.getPassword()));
        user.setRole(request.getRole());
        user =repository.save(user);

        System.out.println("Senha na classe authen: " + user.getPassword());

        String token = jwtService.generateToken(user);
        return new AuthenticationResponse(token);

    }

    // public AuthenticationResponse authenticationResponse(User request){
    //     authenticationManager.authenticate(
    //             new UsernamePasswordAuthenticationToken(
    //                  request.getUsername() ,
    //                  request.getPassword()
    //             )
    //     );
    //       User user = repository.findByUsername(request.getUsername()).orElseThrow();
    //       String token =jwtService.generateToken(user);
    //       System.out.println("login com sucesso");
    //       return new AuthenticationResponse(token);
    // }

    public AuthenticationResponse authenticationResponse(User request) {
        System.out.println("CHEGOUUUU");
        try {
            authenticationManager.authenticate(
                    new UsernamePasswordAuthenticationToken(
                            request.getUsername(),
                            request.getPassword()
                    )
            );
        } catch (Exception e) {
            if (repository.findByUsername(request.getUsername()).isEmpty()) {
                System.out.println("Usuário não encontrado: " + request.getUsername());
                throw new RuntimeException("Usuário não encontrado");
            } else {
                System.out.println("Senha incorreta para o usuário: " + request.getUsername());
                System.out.println("Senha digitada: " + request.getPassword());
                throw new RuntimeException("Senha incorreta");
            }
        }

        User user = repository.findByUsername(request.getUsername()).orElseThrow();
        String token = jwtService.generateToken(user);
        System.out.println("Login com sucesso para: " + request.getUsername());
        return new AuthenticationResponse(token);
    }

}